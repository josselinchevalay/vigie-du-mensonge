import {createFileRoute, Outlet, redirect} from "@tanstack/react-router";
import {authManager} from "@/core/dependencies/auth/authManager.ts";
import {RedactorClient} from "@/core/dependencies/redactor/redactorClient.ts";
import {api} from "@/core/dependencies/api.ts";

export const Route = createFileRoute('/redactor')({
    beforeLoad: () => {
        const auth = authManager.authStore.state;
        if (!auth || !auth.isRedactor) {
            throw redirect({to: '/', replace: true});
        }

        const redactorClient = new RedactorClient(api);
        return {redactorClient};
    },
    component: () => <Outlet/>,
});