import {createFileRoute} from '@tanstack/react-router';
import {useQuery} from "@tanstack/react-query";
import {BasicProgress} from "@/core/components/BasicProgress.tsx";
import {ModeratorArticlesByReference} from "@/core/components/moderator/ModeratorArticlesByReference.tsx";

export const Route = createFileRoute('/moderator/articles/$articleRef')({
    beforeLoad: ({params}) => {
        const articleRef = params.articleRef;
        return {articleRef: articleRef};
    },
    component: RouteComponent,
});

function RouteComponent() {
    const {articleRef} = Route.useParams();
    const moderatorClient = Route.useRouteContext().moderatorClient;

    const {data: articles, isLoading, isError} = useQuery({
        queryKey: ["moderator", "articles", articleRef],
        queryFn: () => moderatorClient.findModeratorArticlesByRef(articleRef),
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

    if (!articles || articles.length === 0) {
        return (
            <div className="flex flex-col items-center justify-center h-screen">
                Une erreur est survenue. Veuillez réessayer.
            </div>
        );
    }

    return <ModeratorArticlesByReference articles={articles} moderatorClient={moderatorClient}/>;
}
