import {createFileRoute} from "@tanstack/react-router";
import {RedactorArticleFormController} from "@/core/dependencies/redactor/redactorArticleFormController.ts";
import {RedactorArticleForm} from "@/core/components/redactor/RedactorArticleForm.tsx";
import {redactorArticleClient} from "@/core/dependencies/redactor/redactorArticleClient.ts";

export const Route = createFileRoute('/redactor/article-form')({
    beforeLoad: () => {
        const controller = new RedactorArticleFormController(redactorArticleClient);
        return {controller};
    },
    component: RouteComponent,
});

function RouteComponent() {
    const {controller} = Route.useRouteContext() as { controller: RedactorArticleFormController };
    return <RedactorArticleForm controller={controller}/>;
}