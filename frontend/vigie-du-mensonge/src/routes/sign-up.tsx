import {createFileRoute, redirect} from "@tanstack/react-router";
import {authManager} from "@/core/dependencies/auth/auth-manager.ts";
import {SignUp} from "@/core/components/auth/SignUp.tsx";

export const Route = createFileRoute('/sign-up')({
    beforeLoad: () => {
        const isAuthenticated = authManager.authStore.state !== null;
        if (isAuthenticated) {
            throw redirect({to: '/', replace: true});
        }
    },
    component: SignUp,
});
