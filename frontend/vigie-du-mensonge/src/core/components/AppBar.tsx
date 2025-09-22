import {Link} from '@tanstack/react-router';
import {useEffect, useMemo, useState} from 'react';
import {Moon, Sun} from 'lucide-react';
import {authManager} from "@/core/dependencies/auth/authManager.ts";
import {useStore} from "@tanstack/react-store";
import {Button} from "../shadcn/components/ui/button";

function getSystemPrefersDark(): boolean {
    if (typeof window === 'undefined') return false;
    return window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches;
}

function getStoredTheme(): 'dark' | 'light' | null {
    try {
        const v = localStorage.getItem('theme');
        return v === 'dark' || v === 'light' ? v : null;
    } catch {
        return null;
    }
}

function storeTheme(theme: 'dark' | 'light') {
    try {
        localStorage.setItem('theme', theme);
    } catch {
        // Ignore write errors (e.g., privacy mode)
        return;
    }
}

export default function AppBar() {
    const initialIsDark = useMemo(() => {
        const stored = getStoredTheme();
        if (stored) return stored === 'dark';
        return getSystemPrefersDark();
    }, []);

    const auth = useStore(authManager.authStore);

    const [isDark, setIsDark] = useState<boolean>(initialIsDark);

    // Apply theme class to <html> with a short transition
    useEffect(() => {
        const root = document.documentElement;

        // Add a temporary class that enables CSS color transitions
        root.classList.add('theme-transition');
        const timeout = window.setTimeout(() => {
            root.classList.remove('theme-transition');
        }, 300);

        if (isDark) {
            root.classList.add('dark');
            storeTheme('dark');
        } else {
            root.classList.remove('dark');
            storeTheme('light');
        }

        return () => window.clearTimeout(timeout);
    }, [isDark]);

    // Optional: update with system changes if user hasn't explicitly toggled in this session
    useEffect(() => {
        const mql = window.matchMedia('(prefers-color-scheme: dark)');

        const onChange = (e: MediaQueryListEvent) => {
            const stored = getStoredTheme();
            if (!stored) {
                setIsDark(e.matches);
            }
        };

        mql.addEventListener('change', onChange);
        return () => {
            mql.removeEventListener('change', onChange);
        };
    }, []);

    return (
        <nav
            className="sticky top-0 z-40 w-full border-b bg-background/80 backdrop-blur supports-[backdrop-filter]:bg-background/60">
            <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
                <div className="flex h-14 items-center justify-between">
                    <Link to="/" className="text-base font-semibold text-foreground">
                        Vigie du mensonge
                    </Link>

                    <div className="flex items-center gap-2 min-w-0">
                        <button
                            type="button"
                            aria-label={isDark ? 'Activer le thème clair' : 'Activer le thème sombre'}
                            title={isDark ? 'Mode clair' : 'Mode sombre'}
                            aria-pressed={isDark}
                            onClick={() => setIsDark(v => !v)}
                            className="p-2 inline-flex items-center justify-center rounded-md bg-background text-foreground hover:bg-accent hover:text-accent-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 focus-visible:ring-offset-background"
                        >
                            {isDark ? (
                                <Sun className="h-4 w-4" aria-hidden="true"/>
                            ) : (
                                <Moon className="h-4 w-4" aria-hidden="true"/>
                            )}
                        </button>

                        {!auth ? (
                            <>
                                <Link
                                    to="/sign-in"
                                    className="inline-flex items-center rounded-md border px-3 py-2 text-sm font-medium hover:bg-accent hover:text-accent-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 focus-visible:ring-offset-background"
                                >
                                    Connexion
                                </Link>
                                <Link
                                    to="/sign-up"
                                    className="inline-flex items-center rounded-md bg-primary px-3 py-2 text-sm font-medium text-primary-foreground hover:bg-primary/90 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 focus-visible:ring-offset-background"
                                >
                                    Inscription
                                </Link>
                            </>
                        ) : (
                            <>
                                <Button onClick={() => authManager.signOut()}>Déconnexion</Button>
                            </>
                        )}
                    </div>
                </div>
            </div>
        </nav>
    );
}
