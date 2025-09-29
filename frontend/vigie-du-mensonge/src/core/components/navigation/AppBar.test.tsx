import {render, screen} from '@testing-library/react';
import {beforeAll, beforeEach, describe, expect, it, vi} from 'vitest';
import {authManager} from '@/core/dependencies/auth/authManager.ts';
import {Auth} from '@/core/models/auth.ts';

// Mock TanStack Router's Link to avoid needing a RouterProvider

vi.mock('@tanstack/react-router', () => ({
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    Link: ({to, children, ...rest}: any) => (
        <a href={typeof to === 'string' ? to : ''} {...rest}>
            {children}
        </a>
    ),
    // Minimal mock to satisfy AppBar's use of useLocation().pathname
    useLocation: () => ({ pathname: '/' }),
}));

// Provide a matchMedia stub for jsdom (AppBar uses it in an effect)
function ensureMatchMediaStub() {
    Object.defineProperty(window, 'matchMedia', {
        writable: true,
        value: vi.fn().mockImplementation((query: string) => ({
            matches: false,
            media: query,
            onchange: null,
            addEventListener: vi.fn(),
            removeEventListener: vi.fn(),
            addListener: vi.fn(), // deprecated but sometimes accessed
            removeListener: vi.fn(), // deprecated
            dispatchEvent: vi.fn(),
        })),
    });
}

beforeAll(() => {
    ensureMatchMediaStub();
});

beforeEach(() => {
    vi.restoreAllMocks();
});

describe('AppBar', () => {
    it('shows Connexion and Inscription when not authenticated', async () => {
        ensureMatchMediaStub();
        // Ensure store is empty
        authManager.authStore.setState(() => null);

        const {default: AppBar} = await import('./AppBar.tsx');
        render(<AppBar/>);

        expect(screen.getByText('Connexion')).toBeInTheDocument();
        expect(screen.getByText('Inscription')).toBeInTheDocument();
        expect(screen.queryByText('Déconnexion')).not.toBeInTheDocument();
    });

    it('shows Déconnexion when authenticated', async () => {
        ensureMatchMediaStub();
        const now = Date.now();
        const auth = Auth.fromJson({
            accessTokenExpiry: new Date(now + 60 * 60 * 1000).toISOString(),
            refreshTokenExpiry: new Date(now + 2 * 60 * 60 * 1000).toISOString(),
            roles: [],
            tag: 'clemovitch0123',
        });

        authManager.authStore.setState(() => auth);

        const {default: AppBar} = await import('./AppBar.tsx');
        render(<AppBar/>);

        expect(screen.getByText('Déconnexion')).toBeInTheDocument();
        expect(screen.queryByText('Connexion')).not.toBeInTheDocument();
        expect(screen.queryByText('Inscription')).not.toBeInTheDocument();
    });
});
