import {createFileRoute} from "@tanstack/react-router";
import {MeClient} from "@/core/dependencies/me/meClient.ts";
import {api} from "@/core/dependencies/api.ts";
import {MeController} from "@/core/dependencies/me/meController.ts";
import {Me} from "@/core/components/me/Me.tsx";

export const Route = createFileRoute('/me')({
    beforeLoad: () => {
        const client = new MeClient(api);
        const controller = new MeController(client);
        return {controller};
    },
    component: RouteComponent,
});

function RouteComponent() {
    const {controller} = Route.useRouteContext() as { controller: MeController };
    return <Me controller={controller}/>;
}