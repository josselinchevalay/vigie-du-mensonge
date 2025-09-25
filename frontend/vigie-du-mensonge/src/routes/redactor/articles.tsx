import {createFileRoute} from "@tanstack/react-router";
import {RedactorArticles} from "@/core/components/redactor/RedactorArticles.tsx";
import {redactorArticleClient} from "@/core/dependencies/redactor/redactorArticleClient.ts";
import {RedactorArticlesController} from "@/core/dependencies/redactor/redactorArticlesController.ts";

export const Route = createFileRoute('/redactor/articles')({
    beforeLoad: () => {
        const controller = new RedactorArticlesController(redactorArticleClient);
        return {controller};
    },
    component: RouteComponent,
});

function RouteComponent() {
    const {controller} = Route.useRouteContext() as { controller: RedactorArticlesController };
    return <RedactorArticles controller={controller}/>;
}