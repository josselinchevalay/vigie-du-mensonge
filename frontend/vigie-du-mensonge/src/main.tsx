import {StrictMode} from 'react';
import ReactDOM from 'react-dom/client';
import {createRouter, RouterProvider} from '@tanstack/react-router';
import './index.css';

// Import the generated route tree
import {routeTree} from './routeTree.gen';
import {authManager} from "@/core/dependencies/auth/authManager.ts";
import {queryClient} from "@/core/dependencies/queryClient.ts";
import {QueryClientProvider} from "@tanstack/react-query";
import {ThemeProvider} from "@/theme-provider/ThemeProvider.tsx";

// Create a new router instance
export const router = createRouter({
    routeTree,
});

// Register the router instance for type safety
declare module '@tanstack/react-router' {
    interface Register {
        router: typeof router
    }
}

// Re-run route loaders/guards whenever auth changes
const unsubscribeAuth = authManager.authStore.subscribe(() => {
    void router.invalidate();
});

// Clean up on HMR
if (import.meta.hot) {
    import.meta.hot.dispose(() => {
        unsubscribeAuth();
    });
}

// Render the app (guarded for test environments without a #root element)
const rootElement = document.getElementById('root');
if (rootElement && !rootElement.innerHTML) {
    const root = ReactDOM.createRoot(rootElement);
    root.render(
        <StrictMode>
            <ThemeProvider>
                <QueryClientProvider client={queryClient}>
                    <RouterProvider router={router}/>
                </QueryClientProvider>
            </ThemeProvider>
        </StrictMode>,
    );
}