import {createFileRoute} from "@tanstack/react-router";
import {RedactorArticleForm} from "@/core/components/redactor/RedactorArticleForm.tsx";

export const Route = createFileRoute('/redactor/new-draft')({
    component: RouteComponent,
});

function RouteComponent() {
    const redactorClient = Route.useRouteContext().redactorClient;
    return <RedactorArticleForm redactorClient={redactorClient}/>;
}