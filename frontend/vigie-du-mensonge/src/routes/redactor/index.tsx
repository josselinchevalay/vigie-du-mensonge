import {createFileRoute} from "@tanstack/react-router";
import {RedactorClient} from "@/core/dependencies/redactor/redactorClient.ts";
import {api} from "@/core/dependencies/api.ts";
import {RedactorController} from "@/core/dependencies/redactor/redactorController.ts";
import {RedactorIndex} from "@/core/components/redactor/RedactorIndex.tsx";

export const Route = createFileRoute('/redactor/')({
    beforeLoad: () => {
        const client = new RedactorClient(api);
        const controller = new RedactorController(client);
        return {controller};
    },
    component: RouteComponent,
});

function RouteComponent() {
    const {controller} = Route.useRouteContext() as { controller: RedactorController };
    return <RedactorIndex controller={controller}/>;
}