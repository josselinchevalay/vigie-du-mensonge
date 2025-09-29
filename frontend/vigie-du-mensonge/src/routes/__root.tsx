import {Toaster} from "@/core/shadcn/components/ui/sonner";
import {createRootRoute, Outlet} from '@tanstack/react-router';
import {TanStackRouterDevtools} from '@tanstack/react-router-devtools';
import {useEffect, useRef} from 'react';
import {authManager} from '@/core/dependencies/auth/authManager.ts';
import {AuthRefreshScheduler} from '@/core/dependencies/auth/authRefreshScheduler.ts';
import AppBar from "@/core/components/navigation/AppBar.tsx";

const RootLayout = () => {
    const schedulerRef = useRef<AuthRefreshScheduler | null>(null);

    useEffect(() => {
        // Create and start scheduler
        const scheduler = new AuthRefreshScheduler({
            getAuth: () => authManager.authStore.state,
            onRefresh: () => authManager.refresh(),
        });
        scheduler.start();
        schedulerRef.current = scheduler;

        // Reschedule whenever auth changes
        const unsub = authManager.authStore.subscribe(() => {
            scheduler.reschedule();
        });

        // Cleanup on unmount
        return () => {
            unsub();
            scheduler.stop();
            schedulerRef.current = null;
        };
    }, []);

    return (
        <>
            <a
                href="#main-content"
                className="sr-only focus:not-sr-only focus:fixed focus:top-2 focus:left-2 focus:z-50 focus:rounded focus:px-3 focus:py-2 focus:ring"
            >
                Aller au contenu
            </a>

            <div className="flex min-h-dvh w-full flex-col">
                <header><AppBar/></header>

                <main id="main-content"
                      className="flex-1 overflow-y-auto py-4">
                    <Outlet/>
                </main>
            </div>

            <Toaster position="top-center" duration={3000}/>

            {import.meta.env.DEV ? <TanStackRouterDevtools/> : null}
        </>
    );
};

export const Route = createRootRoute({component: RootLayout});