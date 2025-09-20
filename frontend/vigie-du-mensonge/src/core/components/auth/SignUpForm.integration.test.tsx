import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { beforeAll, describe, it, expect, vi } from 'vitest';
import { RouterProvider, createMemoryHistory, createRouter, createRootRoute, createRoute } from '@tanstack/react-router';
import { Suspense } from 'react';
import { SignUpForm } from './SignUpForm';
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

// Integration test using MSW: do not mock authManager or Router

function buildTestRouter(initialPath: string) {
  const rootRoute = createRootRoute({
    component: () => (
      <div>
        <Suspense fallback={<div>Loading…</div>}>
          <SignUpForm />
        </Suspense>
        <Toaster position="top-center" duration={3000} />
      </div>
    ),
  });

  // Minimal home route to navigate to
  const homeRoute = createRoute({
    getParentRoute: () => rootRoute,
    path: '/',
    component: () => <h3>Welcome Home!</h3>,
  });

  const routeTree = rootRoute.addChildren([homeRoute]);
  const history = createMemoryHistory({ initialEntries: [initialPath] });
  const router = createRouter({ routeTree, history });
  return router;
}

beforeAll(() => {
  ensureMatchMediaStub();
});

describe('SignUpForm integration (MSW)', () => {
  it('sends correct payload and navigates home (Approach 1)', async () => {
    // Spy resolver to assert request shape
    const resolver = vi.fn(async ({ request }) => {
      const url = new URL(request.url);
      expect(url.pathname).toBe('/api/v1/auth/sign-up');
      expect(request.method).toBe('POST');
      const body = await request.json();
      expect(body).toEqual({ email: 'jane@doe.com', password: 'StrongP@ssw0rd' });

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
      http.post('/api/v1/auth/sign-up', resolver),
      http.post('http://localhost:8080/api/v1/auth/sign-up', resolver)
    );

    const router = buildTestRouter('/');

    render(
      <Suspense fallback={<div>Loading…</div>}>
        <RouterProvider router={router} />
      </Suspense>
    );

    // Wait for form heading
    await screen.findByText('Créer un compte');

    // Fill email
    await userEvent.type(screen.getByPlaceholderText('vous@exemple.com'), 'jane@doe.com');

    // Fill both password fields (same placeholder)
    const passwordInputs = screen.getAllByPlaceholderText('••••••••');
    await userEvent.type(passwordInputs[0], 'StrongP@ssw0rd');
    await userEvent.type(passwordInputs[1], 'StrongP@ssw0rd');

    // Submit
    await userEvent.click(screen.getByRole('button', { name: /créer le compte/i }));

    // Assert MSW handler was called exactly once (wait for async network)
    await vi.waitFor(() => expect(resolver).toHaveBeenCalledTimes(1));

    // After MSW sign-up, component should navigate to '/'
    expect(router.state.location.pathname).toBe('/');
  });

  it('shows error toast on 409 and stays on the page', async () => {
    const resolver409 = vi.fn(async () => {
      return HttpResponse.text('', { status: 409 });
    });
    server.resetHandlers();
    server.use(
      http.post('/api/v1/auth/sign-up', resolver409),
      http.post('http://localhost:8080/api/v1/auth/sign-up', resolver409)
    );

    const router = buildTestRouter('/');

    render(
      <Suspense fallback={<div>Loading…</div>}>
        <RouterProvider router={router} />
      </Suspense>
    );

    await screen.findByText('Créer un compte');
    await userEvent.type(screen.getByPlaceholderText('vous@exemple.com'), 'existing@doe.com');
    const pwInputs = screen.getAllByPlaceholderText('••••••••');
    await userEvent.type(pwInputs[0], 'StrongP@ssw0rd');
    await userEvent.type(pwInputs[1], 'StrongP@ssw0rd');
    await userEvent.click(screen.getByRole('button', { name: /créer le compte/i }));

    await vi.waitFor(() => expect(resolver409).toHaveBeenCalledTimes(1));

    // Assert the specific toast message for 409
    expect(
      await screen.findByText('Cette adresse email est déjà associée à un compte.')
    ).toBeInTheDocument();
  });
});
