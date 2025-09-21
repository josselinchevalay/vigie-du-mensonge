import {createFileRoute} from '@tanstack/react-router';
import {PasswordUpdateController} from "@/core/dependencies/password_update/passwordUpdateController.ts";
import {PasswordUpdate} from "@/core/components/password_update/PasswordUpdate.tsx";

export const Route = createFileRoute('/password-update')({
    validateSearch: (search: { token?: string }) => ({token: search.token}),
    beforeLoad: ({search}) => {
        const controller = new PasswordUpdateController(search.token ?? null);
        return {controller};
    },
    component: RouteComponent,
});

function RouteComponent() {
    const {controller} = Route.useRouteContext() as { controller: PasswordUpdateController };
    return <PasswordUpdate controller={controller}/>;
}
