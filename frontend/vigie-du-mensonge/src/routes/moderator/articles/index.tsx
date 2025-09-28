import {createFileRoute} from '@tanstack/react-router';
import {useQuery} from "@tanstack/react-query";
import {BasicProgress} from "@/core/components/BasicProgress.tsx";
import {Link} from "@/core/utils/router.ts";
import {ArticleOverviewGrid} from "@/core/components/article/ArticleOverviewGrid.tsx";
import type {Article} from "@/core/models/article.ts";

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

                : <ArticleOverviewGrid articles={articles}
                                       articleHeader={ArticleHeader}/>
        }

    </div>;
}

function ArticleHeader(article: Article) {
    return <div className="flex sm:flex-row flex-col items-center justify-center gap-2">
        <Link
            to="/moderator/articles/$articleRef"
            params={{articleRef: article.reference}}
            className="inline-flex items-center rounded-md border px-3 py-2 text-sm font-medium hover:bg-accent"
        >
            review
        </Link>
    </div>;
}