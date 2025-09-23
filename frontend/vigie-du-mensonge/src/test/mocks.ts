import React from 'react';
import { vi } from 'vitest';

// 1) Mock the adapter toast to prevent timers and allow call assertions
vi.mock('@/core/utils/toast', () => {
  const toast = Object.assign(vi.fn(), {
    success: vi.fn(),
    error: vi.fn(),
    message: vi.fn(),
    dismiss: vi.fn(),
  });
  return { toast };
});

// 2) Mock the Toaster component to a no-op so no timers or matchMedia are used
vi.mock('@/core/shadcn/components/ui/sonner', () => ({
  Toaster: () => null,
}));

// 3) Mock router adapter used by the UI
vi.mock('@/core/utils/router', () => ({
  // simple anchor fallback for Link
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  Link: ({ to, children, ...rest }: any) =>
    React.createElement('a', { href: typeof to === 'string' ? to : '', ...rest }, children),
  useLocation: () => ({ pathname: '/' }),
  useNavigate: () => async () => {},
  redirect: (opts: unknown) => opts,
  navigate: vi.fn(async () => {}),
}));

// 4) Provide a matchMedia stub for jsdom (some UI libs use it)
if (!('matchMedia' in window)) {
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
