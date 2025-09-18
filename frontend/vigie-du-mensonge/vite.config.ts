import {defineConfig} from 'vite';
import react from '@vitejs/plugin-react-swc';
import {tanstackRouter} from '@tanstack/router-plugin/vite';
import tailwindcss from "@tailwindcss/vite";
import path from "path";

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [
        // Please make sure that '@tanstack/router-plugin' is passed before '@vitejs/plugin-react'
        tailwindcss(),
        tanstackRouter({
            target: 'react',
            autoCodeSplitting: true,
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