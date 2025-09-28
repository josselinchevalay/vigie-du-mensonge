import {createFileRoute} from '@tanstack/react-router';
import {useQuery} from "@tanstack/react-query";
import {BasicProgress} from "@/core/components/BasicProgress.tsx";
import {Link} from "@/core/utils/router.ts";
import type {Article} from "@/core/models/article.ts";
import {formatDateFR} from "@/core/utils/formatDate.ts";
import {Separator} from "@/core/shadcn/components/ui/separator.tsx";
import {ArticleOverviewItem} from "@/core/components/article/ArticleOverviewItem.tsx";

export const Route = createFileRoute('/moderator/articles/pending')({
    component: RouteComponent,
});

function RouteComponent() {
    const moderatorClient = Route.useRouteContext().moderatorClient;

    const {data: articles, isLoading, isError} = useQuery({
        queryKey: ["moderator", "articles", "pending"],
        queryFn: () => moderatorClient.getPendingArticles(),
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
            to="/moderator/articles" replace={true}
            className="inline-flex items-center justify-center rounded-md border px-3 py-2 text-sm font-medium hover:bg-accent"
        >
            Voir les articles sous votre modération
        </Link>
        {
            (!articles || articles.length === 0)

                ? <p>Aucun article n'est actuellement en attente de modération.</p>

                : <div className="flex flex-wrap justify-center gap-4 w-full">
                    {articles.map((article) => (
                        <Link
                            key={article.id}
                            to="/moderator/articles/$articleRef"
                            params={{articleRef: article.reference}}
                            className="contents">
                            <ArticleOverviewItem
                                article={article}
                                header={ArticleHeader}
                                showStatus={true}
                            />
                        </Link>
                    ))}
                </div>
        }

    </div>;
}

function ArticleHeader(article: Article) {
    return (
        <div className="flex flex-col items-center gap-2 w-full">
            <p className="text-sm">Rédigé par {article.redactorTag}</p>
            <p className="text-sm">En attente depuis le {formatDateFR(article.updatedAt)}</p>
            <Separator/>
        </div>
    );
}