import {render, screen} from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import {beforeAll, beforeEach, describe, expect, it, vi} from 'vitest';
import {createMemoryHistory, createRootRoute, createRoute, createRouter, RouterProvider} from '@tanstack/react-router';
import {type ReactNode, Suspense} from 'react';
import {http, HttpResponse} from 'msw';
import {server} from '@/test/testServer';
import {Toaster} from '@/core/shadcn/components/ui/sonner';
import {SignUp} from '@/core/components/sign_up/SignUp';
import {SignUpController} from '@/core/dependencies/sign_up/signUpController';
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

describe('SignUp integration (MSW)', () => {
    it('inquire: sends email and shows success toast', async () => {
        const inquireResolver = vi.fn(async ({request}) => {
            const url = new URL(request.url);
            expect(url.pathname).toBe('/api/v1/auth/sign-up/inquire');
            expect(request.method).toBe('POST');
            const body = await request.json();
            expect(body).toEqual({email: 'newuser@example.com'});
            return new HttpResponse(null, {status: 204});
        });

        server.resetHandlers();
        server.use(http.post('http://localhost:8080/api/v1/auth/sign-up/inquire', inquireResolver));

        const controller = new SignUpController(null);
        const router = buildTestRouter(<SignUp controller={controller}/>);

        render(
            <Suspense fallback={<div>Loading…</div>}>
                <RouterProvider router={router}/>
            </Suspense>
        );

        await userEvent.type(await screen.findByPlaceholderText('vous@exemple.com'), 'newuser@example.com');
        await userEvent.click(screen.getByRole('button', {name: /recevoir l'email d'inscription/i}));

        await vi.waitFor(() => expect(inquireResolver).toHaveBeenCalledTimes(1));

        expect(toast.success).toHaveBeenCalledWith("L'e-mail contenant le lien d'inscription a été envoyé.");
    });

    it('inquire: shows error toast on failure', async () => {
        const inquireErrorResolver = vi.fn(async () => new HttpResponse(null, {status: 500}));

        server.resetHandlers();
        server.use(http.post('http://localhost:8080/api/v1/auth/sign-up/inquire', inquireErrorResolver));

        const controller = new SignUpController(null);
        const router = buildTestRouter(<SignUp controller={controller}/>);

        render(
            <Suspense fallback={<div>Loading…</div>}>
                <RouterProvider router={router}/>
            </Suspense>
        );

        await userEvent.type(await screen.findByPlaceholderText('vous@exemple.com'), 'newuser@example.com');
        await userEvent.click(screen.getByRole('button', {name: /recevoir l'email d'inscription/i}));

        await vi.waitFor(() => expect(inquireErrorResolver).toHaveBeenCalledTimes(1));

        expect(toast.error).toHaveBeenCalledWith('Une erreur est survenue. Veuillez réessayer.');
    });

    it('process: sends token and password, shows success toast and navigates', async () => {
        const processResolver = vi.fn(async ({request}) => {
            const url = new URL(request.url);
            expect(url.pathname).toBe('/api/v1/auth/sign-up/process');
            expect(request.method).toBe('POST');
            const body = await request.json();
            expect(body).toEqual({token: 'token-xyz', password: 'Abcdef1!'});

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
        server.use(http.post('http://localhost:8080/api/v1/auth/sign-up/process', processResolver));

        const controller = new SignUpController('token-xyz');
        const router = buildTestRouter(<SignUp controller={controller}/>);

        render(
            <Suspense fallback={<div>Loading…</div>}>
                <RouterProvider router={router}/>
            </Suspense>
        );

        const [passwordInput, confirmInput] = await screen.findAllByPlaceholderText('••••••••');
        await userEvent.type(passwordInput, 'Abcdef1!');
        await userEvent.type(confirmInput, 'Abcdef1!');

        await userEvent.click(screen.getByRole('button', {name: /créer le compte/i}));

        await vi.waitFor(() => expect(processResolver).toHaveBeenCalledTimes(1));

        expect(toast.success).toHaveBeenCalledWith('Votre inscription est terminée!');
        const {navigate} = await import('@/core/utils/router');
        expect(navigate).toHaveBeenCalledWith({to: '/', replace: true});
    });

    it('process: shows error toast on failure', async () => {
        const processErrorResolver = vi.fn(async () => new HttpResponse(null, {status: 400}));

        server.resetHandlers();
        server.use(http.post('http://localhost:8080/api/v1/auth/sign-up/process', processErrorResolver));

        const controller = new SignUpController('token-err');
        const router = buildTestRouter(<SignUp controller={controller}/>);

        render(
            <Suspense fallback={<div>Loading…</div>}>
                <RouterProvider router={router}/>
            </Suspense>
        );

        const [passwordInput, confirmInput] = await screen.findAllByPlaceholderText('••••••••');
        await userEvent.type(passwordInput, 'Abcdef1!');
        await userEvent.type(confirmInput, 'Abcdef1!');

        await userEvent.click(screen.getByRole('button', {name: /créer le compte/i}));

        await vi.waitFor(() => expect(processErrorResolver).toHaveBeenCalledTimes(1));

        expect(toast.error).toHaveBeenCalledWith('Une erreur est survenue. Veuillez réessayer.');
    });
});
