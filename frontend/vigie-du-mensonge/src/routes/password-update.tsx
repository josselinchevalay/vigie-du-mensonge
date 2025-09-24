import {createFileRoute} from '@tanstack/react-router';
import {PasswordUpdateController} from "@/core/dependencies/password_update/passwordUpdateController.ts";
import {PasswordUpdate} from "@/core/components/password_update/PasswordUpdate.tsx";
import {api} from "@/core/dependencies/api.ts";
import {PasswordUpdateClient} from "@/core/dependencies/password_update/passwordUpdateClient.ts";

export const Route = createFileRoute('/password-update')({
    validateSearch: (search: { token?: string }) => ({token: search.token}),
    beforeLoad: ({search}) => {
        const client = new PasswordUpdateClient(api);
        const controller = new PasswordUpdateController(client, search.token ?? null);
        return {controller};
    },
    component: RouteComponent,
});

function RouteComponent() {
    const {controller} = Route.useRouteContext() as { controller: PasswordUpdateController };
    return <PasswordUpdate controller={controller}/>;
}
