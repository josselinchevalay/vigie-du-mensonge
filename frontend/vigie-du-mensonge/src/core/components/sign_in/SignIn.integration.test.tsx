import {render, screen} from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import {describe, expect, it, vi} from 'vitest';
import {http, HttpResponse} from 'msw';
import {server} from '@/test/testServer';
import {SignIn} from '@/core/components/sign_in/SignIn';
import {SignInController} from '@/core/dependencies/sign_in/signInController';
import {toast} from '@/core/utils/toast';

// Router, Toaster, and matchMedia are globally mocked in src/test/mocks.ts via setupFiles

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

        render(<SignIn controller={controller}/>);

        await userEvent.type(await screen.findByPlaceholderText('vous@exemple.com'), 'user@example.com');
        // password input uses placeholder bullets
        await userEvent.type(screen.getByPlaceholderText('••••••••'), 'P@ssword1');

        await userEvent.click(screen.getByRole('button', {name: /se connecter/i}));

        await vi.waitFor(() => expect(successResolver).toHaveBeenCalledTimes(1));

        // Should not show error toast
        expect(toast.error).not.toHaveBeenCalled();

        // Navigation to home (adapter is mocked globally)
        const {navigate} = await import('@/core/utils/router');
        expect(navigate).toHaveBeenCalledWith({to: '/', replace: true});
    });

    it('sign-in: shows specific error toast on 401 invalid credentials', async () => {
        const errorResolver = vi.fn(async () => new HttpResponse(null, {status: 401}));

        server.resetHandlers();
        server.use(http.post('http://localhost:8080/api/v1/auth/sign-in', errorResolver));

        const controller = new SignInController();

        render(<SignIn controller={controller}/>);

        await userEvent.type(await screen.findByPlaceholderText('vous@exemple.com'), 'wrong@example.com');
        await userEvent.type(screen.getByPlaceholderText('••••••••'), 'wrong');

        await userEvent.click(screen.getByRole('button', {name: /se connecter/i}));

        await vi.waitFor(() => expect(errorResolver).toHaveBeenCalledTimes(1));

        expect(toast.error).toHaveBeenCalledWith('Identifiants invalides.');
    });
});
