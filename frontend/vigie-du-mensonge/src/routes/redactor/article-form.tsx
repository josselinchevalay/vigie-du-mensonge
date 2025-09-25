import {createFileRoute, redirect} from "@tanstack/react-router";
import {RedactorArticleFormController} from "@/core/dependencies/redactor/redactorArticleFormController.ts";
import {RedactorArticleForm} from "@/core/components/redactor/RedactorArticleForm.tsx";
import {redactorArticleClient} from "@/core/dependencies/redactor/redactorArticleClient.ts";
import {redactorArticlesManager} from "@/core/dependencies/redactor/redactorArticlesManager.ts";
import {useStore} from "@tanstack/react-store";
import {BasicProgress} from "@/core/components/BasicProgress.tsx";

export const Route = createFileRoute('/redactor/article-form')({
    validateSearch: (search: { articleID?: string }) => ({articleID: search.articleID}),
    beforeLoad: ({search}) => {
        const articleID = search.articleID;

        if (articleID) {
            const article = redactorArticlesManager.articlesStore.state
                .find(a => a.id === articleID);

            if (!article) {
                throw redirect({to: '/redactor/articles', replace: true});
            }

            const controller = new RedactorArticleFormController(redactorArticleClient, article);
            return {controller};
        }

        const controller = new RedactorArticleFormController(redactorArticleClient, null);
        return {controller};
    },
    component: RouteComponent,
});

function RouteComponent() {
    const {controller} = Route.useRouteContext() as { controller: RedactorArticleFormController };

    const loading = useStore(controller.loadingStore);
    const err = useStore(controller.errStore);

    if (loading) {
        return (
            <div className="flex flex-col items-center justify-center h-screen">
                Chargement en cours...
                <BasicProgress/>
            </div>
        );
    }

    if (err) {
        return <div className="flex items-center justify-center h-screen">
            Une erreur est survenue. Veuillez réessayer.
        </div>;
    }

    return <RedactorArticleForm controller={controller}/>;
}