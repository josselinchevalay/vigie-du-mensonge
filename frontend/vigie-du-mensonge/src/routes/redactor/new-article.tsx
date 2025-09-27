import {createFileRoute} from "@tanstack/react-router";
import {RedactorArticleForm} from "@/core/components/redactor/RedactorArticleForm.tsx";

export const Route = createFileRoute('/redactor/new-article')({
    component: RouteComponent,
});

function RouteComponent() {
    const redactorClient = Route.useRouteContext().redactorClient;
    return <RedactorArticleForm redactorClient={redactorClient}/>;
}