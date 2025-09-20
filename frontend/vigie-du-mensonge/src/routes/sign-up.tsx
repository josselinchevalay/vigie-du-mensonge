import {createFileRoute, redirect} from "@tanstack/react-router";
import {authManager} from "@/core/dependencies/auth/authManager.ts";
import {SignUpForm} from "@/core/components/auth/SignUpForm.tsx";

export const Route = createFileRoute('/sign-up')({
    beforeLoad: () => {
        const isAuthenticated = authManager.authStore.state !== null;
        if (isAuthenticated) {
            throw redirect({to: '/', replace: true});
        }
    },
    component: SignUpForm,
});
