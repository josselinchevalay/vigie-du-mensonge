import {StrictMode} from 'react';
import ReactDOM from 'react-dom/client';
import {createRouter, RouterProvider} from '@tanstack/react-router';

const mq = window.matchMedia('(prefers-color-scheme: dark)');
document.documentElement.classList.toggle('dark', mq.matches);

import './index.css';

// Import the generated route tree
import {routeTree} from './routeTree.gen';

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

// Render the app (guarded for test environments without a #root element)
const rootElement = document.getElementById('root');
if (rootElement && !rootElement.innerHTML) {
    const root = ReactDOM.createRoot(rootElement);
    root.render(
        <StrictMode>
            <RouterProvider router={router} />
        </StrictMode>,
    );
}