import {createFileRoute} from '@tanstack/react-router';
import {useQuery} from "@tanstack/react-query";
import {RedactorArticleForm} from "@/core/components/redactor/RedactorArticleForm.tsx";
import {BasicProgress} from "@/core/components/BasicProgress.tsx";

export const Route = createFileRoute('/redactor/edit-article/$articleId')({
    beforeLoad: ({params}) => {
        const articleId = params.articleId;
        return {articleId: articleId};
    },
    component: RouteComponent,
});

function RouteComponent() {
    const {articleId} = Route.useParams();
    const redactorClient = Route.useRouteContext().redactorClient;

    const {data: article, isLoading, isError} = useQuery({
        queryKey: ["redactor", "article", articleId],
        queryFn: () => redactorClient.findArticleById(articleId),
    });

    if (isError) {
        return <div className="flex items-center justify-center h-screen">
            Une erreur est survenue. Veuillez réessayer.
        </div>;
    }

    if (isLoading) {
        return (
            <div className="flex flex-col items-center justify-center h-screen">
                Chargement en cours...
                <BasicProgress/>
            </div>
        );
    }

    return <RedactorArticleForm redactorClient={redactorClient} article={article}/>;
}
