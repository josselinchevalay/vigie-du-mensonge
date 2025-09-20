import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import {beforeAll, describe, expect, it, vi } from 'vitest';
import { createMemoryHistory, createRootRoute, createRoute, createRouter, RouterProvider } from '@tanstack/react-router';
import { Suspense } from 'react';
import { SignInForm } from './SignInForm';
import { http, HttpResponse } from 'msw';
import { server } from '@/test/testServer';
import { Toaster } from '@/core/shadcn/components/ui/sonner';

// Provide a matchMedia stub for jsdom (Toaster uses it internally)
function ensureMatchMediaStub() {
     
    Object.defineProperty(window, 'matchMedia', {
        writable: true,
        value: vi.fn().mockImplementation((query: string) => ({
            matches: false,
            media: query,
            onchange: null,
            addEventListener: vi.fn(),
            removeEventListener: vi.fn(),
            addListener: vi.fn(),
            removeListener: vi.fn(),
            dispatchEvent: vi.fn(),
        })),
    });
}

// Important: We do not mock anything here. MSW is started in setupTests.ts

function buildTestRouter(initialPath: string) {
    const rootRoute = createRootRoute({
        component: () => (
            <div>
                <Suspense fallback={<div>Loading…</div>}>
                    <SignInForm />
                </Suspense>
                <Toaster position="top-center" duration={3000} />
            </div>
        ),
    });

    // Provide a minimal child route for '/'
    const homeRoute = createRoute({
        getParentRoute: () => rootRoute,
        path: '/',
        component: () => <h3>Welcome Home!</h3>,
    });

    const routeTree = rootRoute.addChildren([homeRoute]);
    const history = createMemoryHistory({ initialEntries: [initialPath] });
    return createRouter({ routeTree, history });
}

beforeAll(() => {
    ensureMatchMediaStub();
});

describe('SignInForm integration (MSW)', () => {
    it('sends correct payload and navigates home', async () => {
        // Spy resolver to assert request shape
        const resolver = vi.fn(async ({ request }) => {
            const url = new URL(request.url);
            expect(url.pathname).toBe('/api/v1/auth/sign-in');
            expect(request.method).toBe('POST');
            const body = await request.json();
            expect(body).toEqual({ email: 'john@doe.com', password: 'SuperSecret!' });

            const now = Date.now();
            return HttpResponse.json({
                accessTokenExpiry: new Date(now + 60 * 60 * 1000).toISOString(),
                refreshTokenExpiry: new Date(now + 2 * 60 * 60 * 1000).toISOString(),
                emailVerified: true,
                roles: [],
            });
        });
        server.resetHandlers();
        server.use(
            http.post('/api/v1/auth/sign-in', resolver),
            http.post('http://localhost:8080/api/v1/auth/sign-in', resolver)
        );

        const router = buildTestRouter('/');

        render(
            <Suspense fallback={<div>Loading…</div>}>
                <RouterProvider router={router} />
            </Suspense>
        );

        // Wait for the form to appear
        await screen.findByText('Connexion');

        // Fill the form by placeholders as in the UI
        await userEvent.type(screen.getByPlaceholderText('vous@exemple.com'), 'john@doe.com');
        await userEvent.type(screen.getByPlaceholderText('••••••••'), 'SuperSecret!');

        // Submit
        await userEvent.click(screen.getByRole('button', { name: /se connecter/i }));

        // Assert MSW handler was called exactly once (wait for async network)
        await vi.waitFor(() => expect(resolver).toHaveBeenCalledTimes(1));

        // After successful MSW sign-in, the component navigates to '/'
        expect(router.state.location.pathname).toBe('/');
    });

    it('shows error toast on 401 and stays on the page', async () => {
        const resolver401 = vi.fn(async () => {
            return HttpResponse.text('', { status: 401 });
        });
        server.resetHandlers();
        server.use(
            http.post('/api/v1/auth/sign-in', resolver401),
            http.post('http://localhost:8080/api/v1/auth/sign-in', resolver401)
        );

        const router = buildTestRouter('/');

        render(
            <Suspense fallback={<div>Loading…</div>}>
                <RouterProvider router={router} />
            </Suspense>
        );

        await screen.findByText('Connexion');
        await userEvent.type(screen.getByPlaceholderText('vous@exemple.com'), 'john@doe.com');
        await userEvent.type(screen.getByPlaceholderText('••••••••'), 'WrongPassword');
        await userEvent.click(screen.getByRole('button', { name: /se connecter/i }));

        await vi.waitFor(() => expect(resolver401).toHaveBeenCalledTimes(1));

        // Assert the specific toast message for 401
        expect(await screen.findByText('Identifiants invalides.')).toBeInTheDocument();
    });

    it('shows error toast on 404 and stays on the page', async () => {
        const resolver404 = vi.fn(async () => {
            return HttpResponse.text('', { status: 404 });
        });
        server.resetHandlers();
        server.use(
            http.post('/api/v1/auth/sign-in', resolver404),
            http.post('http://localhost:8080/api/v1/auth/sign-in', resolver404)
        );

        const router = buildTestRouter('/');

        render(
            <Suspense fallback={<div>Loading…</div>}>
                <RouterProvider router={router} />
            </Suspense>
        );

        await screen.findByText('Connexion');
        await userEvent.type(screen.getByPlaceholderText('vous@exemple.com'), 'unknown@doe.com');
        await userEvent.type(screen.getByPlaceholderText('••••••••'), 'AnyPassword');
        await userEvent.click(screen.getByRole('button', { name: /se connecter/i }));

        await vi.waitFor(() => expect(resolver404).toHaveBeenCalledTimes(1));

        // Assert the specific toast message for 404
        expect(await screen.findByText('Aucun compte ne correspond à cette adresse email.')).toBeInTheDocument();
    });
});
