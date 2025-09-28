import {render, screen} from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import {describe, expect, it, vi} from 'vitest';
import {http, HttpResponse} from 'msw';
import {server} from '@/test/testServer.ts';
import {SignUp} from '@/core/components/auth/SignUp.tsx';
import {SignUpController} from '@/core/dependencies/auth/signUpController.ts';
import {toast} from '@/core/utils/toast.ts';
import {authClient} from "@/core/dependencies/auth/authClient.ts";

// Router, Toaster, and matchMedia are globally mocked in src/test/mocks.ts via setupFiles

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

        const controller = new SignUpController(authClient, null);

        render(<SignUp controller={controller}/>);

        await userEvent.type(await screen.findByPlaceholderText('vous@exemple.com'), 'newuser@example.com');
        await userEvent.click(screen.getByRole('button', {name: /recevoir l'email d'inscription/i}));

        await vi.waitFor(() => expect(inquireResolver).toHaveBeenCalledTimes(1));

        expect(toast.success).toHaveBeenCalledWith("L'e-mail contenant le lien d'inscription a été envoyé.");
    });

    it('inquire: shows error toast on failure', async () => {
        const inquireErrorResolver = vi.fn(async () => new HttpResponse(null, {status: 500}));

        server.resetHandlers();
        server.use(http.post('http://localhost:8080/api/v1/auth/sign-up/inquire', inquireErrorResolver));

        const controller = new SignUpController(authClient, null);

        render(<SignUp controller={controller}/>);

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
            expect(body).toEqual({token: 'token-xyz', password: 'Abcdef1!', username: 'clemovitch'});

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

        const controller = new SignUpController(authClient, 'token-xyz');

        render(<SignUp controller={controller}/>);

        const [passwordInput, confirmInput] = await screen.findAllByPlaceholderText('••••••••');
        await userEvent.type(passwordInput, 'Abcdef1!');
        await userEvent.type(confirmInput, 'Abcdef1!');
        await userEvent.type(await screen.findByPlaceholderText('username'), 'clemovitch');

        await userEvent.click(screen.getByRole('button', {name: /créer le compte/i}));

        await vi.waitFor(() => expect(processResolver).toHaveBeenCalledTimes(1));

        expect(toast.success).toHaveBeenCalledWith('Votre inscription est terminée!');
        const {navigate} = await import('@/core/utils/router.ts');
        expect(navigate).toHaveBeenCalledWith({to: '/', replace: true});
    });

    it('process: shows error toast on failure', async () => {
        const processErrorResolver = vi.fn(async () => new HttpResponse(null, {status: 400}));

        server.resetHandlers();
        server.use(http.post('http://localhost:8080/api/v1/auth/sign-up/process', processErrorResolver));

        const controller = new SignUpController(authClient, 'token-err');

        render(<SignUp controller={controller}/>);


        const [passwordInput, confirmInput] = await screen.findAllByPlaceholderText('••••••••');
        await userEvent.type(passwordInput, 'Abcdef1!');
        await userEvent.type(confirmInput, 'Abcdef1!');
        await userEvent.type(await screen.findByPlaceholderText('username'), 'clemovitch');

        await userEvent.click(screen.getByRole('button', {name: /créer le compte/i}));

        await vi.waitFor(() => expect(processErrorResolver).toHaveBeenCalledTimes(1));

        expect(toast.error).toHaveBeenCalledWith('Une erreur est survenue. Veuillez réessayer.');
    });
});
