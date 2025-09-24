import {createFileRoute} from "@tanstack/react-router";
import {ArticleFormController} from "@/core/dependencies/article/articleFormController.ts";
import {articleClient} from "@/core/dependencies/article/articleClient.ts";
import {ArticleForm} from "@/core/components/article/ArticleForm.tsx";

export const Route = createFileRoute('/article-form')({
    beforeLoad: () => {
        const controller = new ArticleFormController(articleClient);
        return {controller};
    },
    component: RouteComponent,
});

function RouteComponent() {
    const {controller} = Route.useRouteContext() as { controller: ArticleFormController };
    return <ArticleForm controller={controller}/>;
}