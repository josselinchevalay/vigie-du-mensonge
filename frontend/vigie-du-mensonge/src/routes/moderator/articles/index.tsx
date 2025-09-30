import {createFileRoute} from '@tanstack/react-router';
import {useQuery} from "@tanstack/react-query";
import {Link} from "@/core/utils/router.ts";
import {ArticleOverviewItem} from "@/core/components/article/ArticleOverviewItem.tsx";
import {Spinner} from "@/core/shadcn/components/ui/spinner.tsx";

export const Route = createFileRoute('/moderator/articles/')({
    component: RouteComponent,
});

function RouteComponent() {
    const moderatorClient = Route.useRouteContext().moderatorClient;

    const {queryKey, queryFn} = moderatorClient.getModeratorArticles();

    const {data: articles, isLoading, isError} = useQuery({
        queryKey: queryKey,
        queryFn: () => queryFn(),
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
            <div className="flex flex-col gap-2 items-center justify-center h-screen">
                Chargement en cours...
                <Spinner/>
            </div>
        );
    }

    return <div className="flex flex-col items-center gap-8 min-w-0 p-2">
        <Link
            to="/moderator/articles/pending" replace={true}
            className="p-2 text-sm font-medium rounded-md hover:bg-accent"
        >
            Voir les articles en attente de modération
        </Link>

        {
            (!articles || articles.length === 0)

                ? <p className="italic">Aucun article n'est actuellement sous votre modération.</p>

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