import AppBar from "@/core/components/app-bar";
import {Toaster} from "@/core/shadcn/components/ui/sonner";
import {createRootRoute, Outlet} from '@tanstack/react-router';
import {TanStackRouterDevtools} from '@tanstack/react-router-devtools';

const RootLayout = () => (
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

export const Route = createRootRoute({component: RootLayout});