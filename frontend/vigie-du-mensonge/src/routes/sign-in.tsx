import {createFileRoute, redirect} from "@tanstack/react-router";
import {authManager} from "@/core/dependencies/auth/authManager.ts";
import { SignInForm } from "@/core/components/auth/SignInForm";

export const Route = createFileRoute('/sign-in')({
    beforeLoad: () => {
        const isAuthenticated = authManager.authStore.state !== null;
        if (isAuthenticated) {
            throw redirect({to: '/', replace: true});
        }
    },
    component: SignInForm,
});