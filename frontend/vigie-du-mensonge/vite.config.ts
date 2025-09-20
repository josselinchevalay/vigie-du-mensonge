/// <reference types="vitest/config" />
import {defineConfig} from 'vite';
import react from '@vitejs/plugin-react-swc';
import {tanstackRouter} from '@tanstack/router-plugin/vite';
import tailwindcss from "@tailwindcss/vite";
import path from "path";

// https://vitejs.dev/config/
export default defineConfig({
    test: {
        globals: true, // optional; lets you use describe/it/expect without imports
        environment: 'jsdom',
        setupFiles: './src/test/setupTests.ts',
        css: true, // allows importing css (Tailwind)
        coverage: {
            provider: 'v8',
            reportsDirectory: './coverage',
        },
    },
    plugins: [
        // Please make sure that '@tanstack/router-plugin' is passed before '@vitejs/plugin-react'
        tailwindcss(),
        tanstackRouter({
            target: 'react',
            autoCodeSplitting: true,
            routesDirectory: 'src/routes',
            generatedRouteTree: 'src/routeTree.gen.ts',
        }),
        react(),
        // ...,
    ],
    // Shadcn
    resolve: {
        alias: {
            "@": path.resolve(__dirname, "./src"),
        },
    },
});