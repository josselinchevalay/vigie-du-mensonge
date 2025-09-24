import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { describe, expect, it, vi } from 'vitest';
import { http, HttpResponse } from 'msw';
import { server } from '@/test/testServer';
import { PasswordUpdate } from '@/core/components/password_update/PasswordUpdate';
import { PasswordUpdateController } from '@/core/dependencies/password_update/passwordUpdateController';
import { toast } from '@/core/utils/toast';
import {PasswordUpdateClient} from "@/core/dependencies/password_update/passwordUpdateClient.ts";
import {api} from "@/core/dependencies/api.ts";

// Router, Toaster, and matchMedia are globally mocked in src/test/mocks.ts via setupFiles

describe('PasswordUpdate integration (MSW)', () => {
  it("inquire: sends request and shows success toast", async () => {
    const inquireResolver = vi.fn(async ({ request }) => {
      const url = new URL(request.url);
      expect(url.pathname).toBe('/api/v1/password-update/inquire');
      expect(request.method).toBe('POST');
      const body = await request.json();
      expect(body).toEqual({ email: 'user@example.com' });
      return new HttpResponse(null, { status: 204 });
    });

    server.resetHandlers();
    server.use(http.post('http://localhost:8080/api/v1/password-update/inquire', inquireResolver));

    const controller = new PasswordUpdateController(new PasswordUpdateClient(api), null);

    render(<PasswordUpdate controller={controller} />);

    // Fill email and submit
    const emailInput = await screen.findByPlaceholderText('vous@exemple.com');
    await userEvent.clear(emailInput);
    await userEvent.type(emailInput, 'user@example.com');

    await userEvent.click(screen.getByRole('button', { name: /recevoir l'email de modification/i }));

    // Assert request happened
    await vi.waitFor(() => expect(inquireResolver).toHaveBeenCalledTimes(1));

    // Success toast was triggered via adapter
    expect(toast.success).toHaveBeenCalledWith("L'email de modification a été envoyé");

    // UI shows success text from the form
    expect(
      await screen.findByText("L'email contenant le lien de modification a été envoyé.")
    ).toBeInTheDocument();
  });

  it('inquire: shows error toast on failure', async () => {
    const inquireErrorResolver = vi.fn(async () => new HttpResponse(null, { status: 400 }));

    server.resetHandlers();
    server.use(http.post('http://localhost:8080/api/v1/password-update/inquire', inquireErrorResolver));

      const controller = new PasswordUpdateController(new PasswordUpdateClient(api), null);

    render(<PasswordUpdate controller={controller} />);

    const emailInput = await screen.findByPlaceholderText('vous@exemple.com');
    await userEvent.type(emailInput, 'user@example.com');

    await userEvent.click(screen.getByRole('button', { name: /recevoir l'email de modification/i }));

    await vi.waitFor(() => expect(inquireErrorResolver).toHaveBeenCalledTimes(1));

    expect(toast.error).toHaveBeenCalledWith('Une erreur est survenue. Veuillez réessayer.');

    // form should still be visible (no success replacement)
    expect(
      await screen.findByRole('button', { name: /recevoir l'email de modification/i })
    ).toBeInTheDocument();
  });

  it('process: sends request with token and shows success toast', async () => {
    const processResolver = vi.fn(async ({ request }) => {
      const url = new URL(request.url);
      expect(url.pathname).toBe('/api/v1/password-update/process');
      expect(request.method).toBe('POST');
      const body = await request.json();
      expect(body).toEqual({ token: 'token-123', newPassword: 'Abcdef1!' });
      return new HttpResponse(null, { status: 204 });
    });

    server.resetHandlers();
    server.use(http.post('http://localhost:8080/api/v1/password-update/process', processResolver));

      const controller = new PasswordUpdateController(new PasswordUpdateClient(api), 'token-123');

    render(<PasswordUpdate controller={controller} />);

    const [passwordInput, confirmInput] = await screen.findAllByPlaceholderText('••••••••');
    await userEvent.type(passwordInput, 'Abcdef1!');
    await userEvent.type(confirmInput, 'Abcdef1!');

    await userEvent.click(screen.getByRole('button', { name: /modifier le mot de passe/i }));

    await vi.waitFor(() => expect(processResolver).toHaveBeenCalledTimes(1));

    // Success toast
    expect(toast.success).toHaveBeenCalledWith('Votre mot de passe a été mis à jour.');

    // UI shows success text from the form
    expect(
      await screen.findByText('Votre mot de passe a été mis à jour.')
    ).toBeInTheDocument();
  });

  it('process: shows error toast on failure and keeps form', async () => {
    const processErrorResolver = vi.fn(async () => new HttpResponse(null, { status: 400 }));

    server.resetHandlers();
    server.use(http.post('http://localhost:8080/api/v1/password-update/process', processErrorResolver));

      const controller = new PasswordUpdateController(new PasswordUpdateClient(api), 'token-err');

    render(<PasswordUpdate controller={controller} />);

    const [passwordInput, confirmInput] = await screen.findAllByPlaceholderText('••••••••');

    await userEvent.type(passwordInput, 'Abcdef1!');
    await userEvent.type(confirmInput, 'Abcdef1!');

    await userEvent.click(screen.getByRole('button', { name: /modifier le mot de passe/i }));

    await vi.waitFor(() => expect(processErrorResolver).toHaveBeenCalledTimes(1));

    expect(toast.error).toHaveBeenCalledWith('Une erreur est survenue. Veuillez réessayer.');
  });
});
