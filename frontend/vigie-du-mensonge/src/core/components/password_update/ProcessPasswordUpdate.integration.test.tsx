import {render, screen} from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import {beforeAll, beforeEach, describe, expect, it, vi} from 'vitest';
import {createMemoryHistory, createRootRoute, createRoute, createRouter, RouterProvider} from '@tanstack/react-router';
import {type ReactNode, Suspense} from 'react';
import {http, HttpResponse} from 'msw';
import {server} from '@/test/testServer';
import {Toaster} from '@/core/shadcn/components/ui/sonner';
import {PasswordUpdate} from '@/core/components/password_update/PasswordUpdate';
import {PasswordUpdateController} from '@/core/dependencies/password_update/passwordUpdateController';
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

describe('ProcessPasswordUpdate integration (MSW)', () => {
    it('process: sends request with token and shows success toast', async () => {
        const processResolver = vi.fn(async ({request}) => {
            const url = new URL(request.url);
            expect(url.pathname).toBe('/api/v1/password-update/process');
            expect(request.method).toBe('POST');
            const body = await request.json();
            expect(body).toEqual({token: 'token-123', newPassword: 'Abcdef1!'});
            return new HttpResponse(null, {status: 204});
        });

        server.resetHandlers();
        server.use(
            http.post('/api/v1/password-update/process', processResolver),
            http.post('http://localhost:8080/api/v1/password-update/process', processResolver),
        );

        // Controller with token -> process form will be shown
        const controller = new PasswordUpdateController('token-123');
        const router = buildTestRouter(<PasswordUpdate controller={controller}/>);

        render(
            <Suspense fallback={<div>Loading…</div>}>
                <RouterProvider router={router}/>
            </Suspense>
        );

        const [passwordInput, confirmInput] = await screen.findAllByPlaceholderText('••••••••');

        await userEvent.type(passwordInput, 'Abcdef1!');
        await userEvent.type(confirmInput, 'Abcdef1!');

        await userEvent.click(screen.getByRole('button', {name: /modifier le mot de passe/i}));

        await vi.waitFor(() => expect(processResolver).toHaveBeenCalledTimes(1));

        // Success toast
        expect(toast.success).toHaveBeenCalledWith('Votre mot de passe a été mis à jour.');

        // UI shows success text from the form
        expect(
            await screen.findByText('Votre mot de passe a été mis à jour.')
        ).toBeInTheDocument();
    });

    it('process: shows error toast on failure and keeps form', async () => {
        const processErrorResolver = vi.fn(async () => new HttpResponse(null, {status: 400}));

        server.resetHandlers();
        server.use(
            http.post('/api/v1/password-update/process', processErrorResolver),
            http.post('http://localhost:8080/api/v1/password-update/process', processErrorResolver),
        );

        const controller = new PasswordUpdateController('token-err');
        const router = buildTestRouter(<PasswordUpdate controller={controller}/>);

        render(
            <Suspense fallback={<div>Loading…</div>}>
                <RouterProvider router={router}/>
            </Suspense>
        );

        const [passwordInput, confirmInput] = await screen.findAllByPlaceholderText('••••••••');

        await userEvent.type(passwordInput, 'Abcdef1!');
        await userEvent.type(confirmInput, 'Abcdef1!');

        await userEvent.click(screen.getByRole('button', {name: /modifier le mot de passe/i}));

        await vi.waitFor(() => expect(processErrorResolver).toHaveBeenCalledTimes(1));

        expect(toast.error).toHaveBeenCalledWith('Une erreur est survenue. Veuillez réessayer.');
    });
});
