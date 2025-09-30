import {createFileRoute, Outlet, redirect} from '@tanstack/react-router';
import {authManager} from "@/core/dependencies/auth/authManager.ts";
import {AdminClient} from "@/core/dependencies/admin/adminClient.ts";
import {api} from "@/core/dependencies/api.ts";

export const Route = createFileRoute('/admin')({
    beforeLoad: () => {
        const auth = authManager.authStore.state;
        if (!auth || !auth.isAdmin) {
            throw redirect({to: '/', replace: true});
        }

        const adminClient = new AdminClient(api);
        return {adminClient};
    },
    component: () => <Outlet/>,
});
