import AppBar from "@/core/components/AppBar.tsx";
import {Toaster} from "@/core/shadcn/components/ui/sonner";
import {createRootRoute, Outlet} from '@tanstack/react-router';
import {TanStackRouterDevtools} from '@tanstack/react-router-devtools';
import { useEffect, useRef } from 'react';
import { authManager } from '@/core/dependencies/auth/authManager.ts';
import { AuthRefreshScheduler } from '@/core/dependencies/auth/authRefreshScheduler.ts';

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
            <div className={"flex min-h-dvh w-full flex-col"}>
                <AppBar/>

                <main id="main-content" role="main" className="flex-1 overflow-auto py-4">
                    <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
                        <Outlet/>
                    </div>
                </main>
            </div>

            <Toaster position="top-center" duration={3000}/>

            {import.meta.env.DEV ? <TanStackRouterDevtools/> : null}
        </>
    );
};

export const Route = createRootRoute({component: RootLayout});