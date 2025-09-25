import {createFileRoute, Outlet, redirect} from "@tanstack/react-router";
import {authManager} from "@/core/dependencies/auth/authManager.ts";

export const Route = createFileRoute('/redactor')({
    beforeLoad: () => {
        const auth = authManager.authStore.state;
        if (!auth || !auth.isRedactor) {
            throw redirect({to: '/', replace: true});
        }
    },
    component: () => <Outlet/>,
});