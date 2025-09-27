import {createFileRoute, Outlet, redirect} from '@tanstack/react-router';
import {authManager} from "@/core/dependencies/auth/authManager.ts";
import {ModeratorClient} from "@/core/dependencies/moderator/moderatorClient.ts";
import {api} from "@/core/dependencies/api.ts";

export const Route = createFileRoute('/moderator')({
    beforeLoad: () => {
        const auth = authManager.authStore.state;
        if (!auth || !auth.isModerator) {
            throw redirect({to: '/', replace: true});
        }

        const moderatorClient = new ModeratorClient(api);
        return {moderatorClient};
    },
    component: () => <Outlet/>,
});

