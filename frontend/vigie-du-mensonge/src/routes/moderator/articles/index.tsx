import {createFileRoute} from '@tanstack/react-router';
import {useQuery} from "@tanstack/react-query";
import {BasicProgress} from "@/core/components/misc/BasicProgress.tsx";
import {Link} from "@/core/utils/router.ts";
import {ArticleOverviewItem} from "@/core/components/article/ArticleOverviewItem.tsx";

export const Route = createFileRoute('/moderator/articles/')({
    component: RouteComponent,
});

function RouteComponent() {
    const moderatorClient = Route.useRouteContext().moderatorClient;

    const {data: articles, isLoading, isError} = useQuery({
        queryKey: ["moderator", "articles"],
        queryFn: () => moderatorClient.getModeratorArticles(),
        staleTime: 60 * 60 * 1000,
    });

    if (isError) {
        return (
            <div className="flex items-center justify-center h-screen">
                Une erreur est survenue. Veuillez réessayer.
            </div>
        );
    }

    if (isLoading) {
        return (
            <div className="flex flex-col items-center justify-center h-screen">
                Chargement en cours...
                <BasicProgress/>
            </div>
        );
    }

    return <div className="flex flex-col items-center gap-8 min-w-0 py-2">
        <Link
            to="/moderator/articles/pending" replace={true}
            className="inline-flex items-center justify-center rounded-md border px-3 py-2 text-sm font-medium hover:bg-accent"
        >
            Voir les articles en attente de modération
        </Link>

        {
            (!articles || articles.length === 0)

                ? <p>Aucun article n'est actuellement sous votre modération.</p>

                : <div className="flex flex-wrap justify-center gap-4 w-full">
                    {articles.map((article) => (
                        <Link
                            key={article.id}
                            to="/moderator/articles/$articleRef"
                            params={{articleRef: article.reference}}
                            className="contents">
                            <ArticleOverviewItem article={article}/>
                        </Link>
                    ))}
                </div>
        }

    </div>;
}