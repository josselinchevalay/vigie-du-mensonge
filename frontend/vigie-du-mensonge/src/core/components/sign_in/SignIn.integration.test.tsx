import {render, screen} from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import {beforeAll, beforeEach, describe, expect, it, vi} from 'vitest';
import {createMemoryHistory, createRootRoute, createRoute, createRouter, RouterProvider} from '@tanstack/react-router';
import {type ReactNode, Suspense} from 'react';
import {http, HttpResponse} from 'msw';
import {server} from '@/test/testServer';
import {Toaster} from '@/core/shadcn/components/ui/sonner';
import {SignIn} from '@/core/components/sign_in/SignIn';
import {SignInController} from '@/core/dependencies/sign_in/signInController';
import {toast} from '@/core/utils/toast';

// Mock the adapter toast to prevent timers and allow call assertions
vi.mock('@/core/utils/toast', () => {
    const toast = Object.assign(vi.fn(), {
        success: vi.fn(),
        error: vi.fn(),
        message: vi.fn(),
        dismiss: vi.fn(),
    });
    return {toast};
});

// Mock the Toaster component to a no-op so no timers or matchMedia are used
vi.mock('@/core/shadcn/components/ui/sonner', () => ({
    Toaster: () => null,
}));

vi.mock('@/core/utils/router', () => ({
    // simple anchor fallback for Link
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    Link: ({to, children, ...rest}: any) => (
        <a href={typeof to === 'string' ? to : ''} {...rest}>
            {children}
        </a>
    ),
    useLocation: () => ({pathname: '/'}),
    useNavigate: () => async () => {
    },
    redirect: (opts: unknown) => opts,
    navigate: vi.fn(async () => {
    }),
}));

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

function buildTestRouter(ui: ReactNode, initialPath: string = '/') {
    const rootRoute = createRootRoute({
        component: () => (
            <div>
                <Suspense fallback={<div>Loading…</div>}>
                    {ui}
                </Suspense>
                <Toaster position="top-center" duration={3000}/>
            </div>
        )
    });

    const homeRoute = createRoute({
        getParentRoute: () => rootRoute,
        path: '/',
        component: () => <h3>Welcome Home!</h3>,
    });

    const routeTree = rootRoute.addChildren([homeRoute]);
    const history = createMemoryHistory({initialEntries: [initialPath]});
    return createRouter({routeTree, history});
}

beforeAll(() => {
    ensureMatchMediaStub();
});

beforeEach(() => {
    vi.clearAllMocks();
});

describe('SignIn integration (MSW)', () => {
    it('sign-in: sends credentials and navigates to home on success', async () => {
        const successResolver = vi.fn(async ({request}) => {
            const url = new URL(request.url);
            expect(url.pathname).toBe('/api/v1/auth/sign-in');
            expect(request.method).toBe('POST');
            const body = await request.json();
            expect(body).toEqual({email: 'user@example.com', password: 'P@ssword1'});

            const access = new Date();
            access.setMinutes(access.getMinutes() + 10);
            const refresh = new Date();
            refresh.setMinutes(refresh.getMinutes() + 30);

            return HttpResponse.json({
                accessTokenExpiry: access.toISOString(),
                refreshTokenExpiry: refresh.toISOString(),
                emailVerified: true,
                roles: ['user'],
            });
        });

        server.resetHandlers();
        server.use(http.post('http://localhost:8080/api/v1/auth/sign-in', successResolver));

        const controller = new SignInController();
        const router = buildTestRouter(<SignIn controller={controller}/>);

        render(
            <Suspense fallback={<div>Loading…</div>}>
                <RouterProvider router={router}/>
            </Suspense>
        );

        await userEvent.type(await screen.findByPlaceholderText('vous@exemple.com'), 'user@example.com');
        // password input uses placeholder bullets
        await userEvent.type(screen.getByPlaceholderText('••••••••'), 'P@ssword1');

        await userEvent.click(screen.getByRole('button', {name: /se connecter/i}));

        await vi.waitFor(() => expect(successResolver).toHaveBeenCalledTimes(1));

        // Should not show error toast
        expect(toast.error).not.toHaveBeenCalled();

        // Navigation to home
        const {navigate} = await import('@/core/utils/router');
        expect(navigate).toHaveBeenCalledWith({to: '/', replace: true});
    });

    it('sign-in: shows specific error toast on 401 invalid credentials', async () => {
        const errorResolver = vi.fn(async () => new HttpResponse(null, {status: 401}));

        server.resetHandlers();
        server.use(http.post('http://localhost:8080/api/v1/auth/sign-in', errorResolver));

        const controller = new SignInController();
        const router = buildTestRouter(<SignIn controller={controller}/>);

        render(
            <Suspense fallback={<div>Loading…</div>}>
                <RouterProvider router={router}/>
            </Suspense>
        );

        await userEvent.type(await screen.findByPlaceholderText('vous@exemple.com'), 'wrong@example.com');
        await userEvent.type(screen.getByPlaceholderText('••••••••'), 'wrong');

        await userEvent.click(screen.getByRole('button', {name: /se connecter/i}));

        await vi.waitFor(() => expect(errorResolver).toHaveBeenCalledTimes(1));

        expect(toast.error).toHaveBeenCalledWith('Identifiants invalides.');
    });
});
