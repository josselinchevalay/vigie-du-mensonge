import {createFileRoute} from "@tanstack/react-router";
import {RedactorArticleFormController} from "@/core/dependencies/redactor/redactorArticleFormController.ts";
import {RedactorArticleForm} from "@/core/components/redactor/RedactorArticleForm.tsx";
import {meClient} from "@/core/dependencies/redactor/redactorClient.ts";

export const Route = createFileRoute('/redactor/article-form')({
    beforeLoad: () => {
        const controller = new RedactorArticleFormController(meClient);
        return {controller};
    },
    component: RouteComponent,
});

function RouteComponent() {
    const {controller} = Route.useRouteContext() as { controller: RedactorArticleFormController };
    return <RedactorArticleForm controller={controller}/>;
}