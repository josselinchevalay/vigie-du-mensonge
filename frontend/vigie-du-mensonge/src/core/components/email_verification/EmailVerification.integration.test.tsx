import {render, screen} from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import {beforeAll, beforeEach, describe, expect, it, vi} from 'vitest';
import {createMemoryHistory, createRootRoute, createRoute, createRouter, RouterProvider} from '@tanstack/react-router';
import {type ReactNode, Suspense} from 'react';
import {http, HttpResponse} from 'msw';
import {server} from '@/test/testServer';
import {Toaster} from '@/core/shadcn/components/ui/sonner';
import {EmailVerification} from './EmailVerification';
import {EmailVerificationController} from '@/core/dependencies/email_verification/emailVerificationController';
import {authManager} from '@/core/dependencies/auth/authManager';
import {Auth} from '@/core/models/auth';
import {toast} from '@/core/utils/toast';
import {navigate} from '@/core/utils/router';

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
    // Reset auth state and local storage between tests
    authManager.authStore.setState(() => null);
    Auth.clearStorage();
    vi.clearAllMocks();
});

describe('EmailVerification integration (MSW)', () => {
    it('inquire: sends request and shows success toast', async () => {
        const inquireResolver = vi.fn(async ({request}) => {
            const url = new URL(request.url);
            expect(url.pathname).toBe('/api/v1/email-verification/inquire');
            expect(request.method).toBe('POST');
            return new HttpResponse(null, {status: 204});
        });

        server.resetHandlers();
        server.use(
            http.post('/api/v1/email-verification/inquire', inquireResolver),
            http.post('http://localhost:8080/api/v1/email-verification/inquire', inquireResolver),
        );

        // Build controller without token (user clicks button to inquire)
        const controller = new EmailVerificationController(null);
        const router = buildTestRouter(<EmailVerification controller={controller}/>);

        render(
            <Suspense fallback={<div>Loading…</div>}>
                <RouterProvider router={router}/>
            </Suspense>
        );

        await screen.findByRole('button', {name: /recevoir l'email de validation/i});
        // Click the inquire button
        await userEvent.click(screen.getByRole('button', {name: /recevoir l'email de validation/i}));

        // Assert request happened
        await vi.waitFor(() => expect(inquireResolver).toHaveBeenCalledTimes(1));

        // Success toast was triggered via adapter
        expect(toast).toHaveBeenCalledWith("L'email de validation a été envoyé");
    });

    it('process: with token processes successfully, marks email verified and navigates home', async () => {
        // Prepare an authenticated state with emailVerified=false
        const now = Date.now();
        const auth = Auth.fromJson({
            accessTokenExpiry: new Date(now + 60 * 60 * 1000).toISOString(),
            refreshTokenExpiry: new Date(now + 2 * 60 * 60 * 1000).toISOString(),
            emailVerified: false,
            roles: [],
        });
        authManager.authStore.setState(() => auth);

        const processResolver = vi.fn(async ({request}) => {
            const url = new URL(request.url);
            expect(url.pathname).toBe('/api/v1/email-verification/process');
            expect(request.method).toBe('POST');
            const body = await request.json();
            expect(body).toEqual({token: 'abc123'});
            return new HttpResponse(null, {status: 204});
        });

        server.resetHandlers();
        server.use(
            http.post('/api/v1/email-verification/process', processResolver),
            http.post('http://localhost:8080/api/v1/email-verification/process', processResolver),
        );


        // Creating the controller with a token will auto-trigger processing
        const controller = new EmailVerificationController('abc123');

        const router = buildTestRouter(<EmailVerification controller={controller}/>);
        render(
            <Suspense fallback={<div>Loading…</div>}>
                <RouterProvider router={router}/>
            </Suspense>
        );

        // Initially, shows the processing state
        expect(await screen.findByText('Validation en cours...')).toBeInTheDocument();

        // Wait for the network call and subsequent navigation
        await vi.waitFor(() => expect(processResolver).toHaveBeenCalledTimes(1));
        await vi.waitFor(() => expect(navigate).toHaveBeenCalledWith({to: '/', replace: true}));

        // Auth should now be marked as verified
        expect(authManager.authStore.state?.emailVerified).toBe(true);
    });

    //TODO: fix error Transitioner not wrapped in act
    it('process: shows error toast on failure and returns to inquire UI', async () => {
        // Prepare an authenticated state with emailVerified=false
        const now = Date.now();
        const auth = Auth.fromJson({
            accessTokenExpiry: new Date(now + 60 * 60 * 1000).toISOString(),
            refreshTokenExpiry: new Date(now + 2 * 60 * 60 * 1000).toISOString(),
            emailVerified: false,
            roles: [],
        });
        authManager.authStore.setState(() => auth);

        const processErrorResolver = vi.fn(async () => {
            return new HttpResponse(null, {status: 400});
        });

        server.resetHandlers();
        server.use(
            http.post('/api/v1/email-verification/process', processErrorResolver),
            http.post('http://localhost:8080/api/v1/email-verification/process', processErrorResolver),
        );

        const controller = new EmailVerificationController('bad-token');

        const router = buildTestRouter(<EmailVerification controller={controller}/>);
        render(
            <Suspense fallback={<div>Loading…</div>}>
                <RouterProvider router={router}/>
            </Suspense>
        );

        // Wait for the request to be performed
        await vi.waitFor(() => expect(processErrorResolver).toHaveBeenCalledTimes(1));

        // Error toast was triggered via adapter
        expect(toast.error).toHaveBeenCalledWith('Une erreur est survenue. Veuillez réessayer.');

        // Token should be cleared, switching to inquire UI -> button should be present
        expect(
            await screen.findByRole('button', {name: /recevoir l'email de validation/i})
        ).toBeInTheDocument();
    });
});
