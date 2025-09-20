import {beforeEach, describe, expect, it, vi} from 'vitest';
import {http, HttpResponse} from 'msw';
import {server} from '@/test/testServer';
import {EmailVerificationController} from './emailVerificationController';
import {authManager} from '@/core/dependencies/auth/authManager';
import {Auth} from '@/core/models/auth';
import {router as appRouter} from '@/main';
import {toast} from '@/core/utils/toast';

// Keep tests small and simple: focus on controller side-effects with MSW

vi.mock('@/core/utils/toast', () => {
    const toast = Object.assign(vi.fn(), {
        success: vi.fn(),
        error: vi.fn(),
        message: vi.fn(),
        dismiss: vi.fn(),
    });
    return { toast };
});

beforeEach(() => {
    // Reset auth and storage between tests
    authManager.authStore.setState(() => null);
    Auth.clearStorage();
    vi.restoreAllMocks();
});

describe('EmailVerificationController (integration with MSW)', () => {
    it('inquire(): performs POST and resolves', async () => {
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

        const controller = new EmailVerificationController(null);
        await controller.inquire();

        expect(inquireResolver).toHaveBeenCalledTimes(1);
    });

    it('process(): on success marks email verified, shows toast and navigates home', async () => {
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
            expect(body).toEqual({token: 'ok-token'});
            return new HttpResponse(null, {status: 204});
        });

        server.resetHandlers();
        server.use(
            http.post('/api/v1/email-verification/process', processResolver),
            http.post('http://localhost:8080/api/v1/email-verification/process', processResolver),
        );

        const navigateSpy = vi.spyOn(appRouter, 'navigate').mockResolvedValue();

        const controller = new EmailVerificationController('ok-token');
        expect(controller.tokenStore.state, 'ok-token');

        await new Promise(resolve => setTimeout(resolve, 100)); //dirty fix for controller.process() being called in constructor
        
        expect(toast).toHaveBeenCalledTimes(1);
        expect(processResolver).toHaveBeenCalledTimes(1);
        expect(authManager.authStore.state?.emailVerified).toBe(true);
        expect(navigateSpy).toHaveBeenCalledWith({to: '/', replace: true});
    });

    it('process(): on error clears token and shows error toast', async () => {
        // Prepare an authenticated state
        const now = Date.now();
        const auth = Auth.fromJson({
            accessTokenExpiry: new Date(now + 60 * 60 * 1000).toISOString(),
            refreshTokenExpiry: new Date(now + 2 * 60 * 60 * 1000).toISOString(),
            emailVerified: false,
            roles: [],
        });
        authManager.authStore.setState(() => auth);

        const processErrorResolver = vi.fn(async () => new HttpResponse(null, {status: 400}));
        server.resetHandlers();
        server.use(
            http.post('/api/v1/email-verification/process', processErrorResolver),
            http.post('http://localhost:8080/api/v1/email-verification/process', processErrorResolver),
        );

        const errorSpy = vi.spyOn(toast, 'error');

        const controller = new EmailVerificationController('bad-token');

        await new Promise(resolve => setTimeout(resolve, 100)); //dirty fix for controller.process() being called in constructor

        expect(processErrorResolver).toHaveBeenCalledTimes(1);
        expect(errorSpy).toHaveBeenCalled();
        expect(controller.tokenStore.state).toBeNull();
    });
});
