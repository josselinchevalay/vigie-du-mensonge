import {createFileRoute} from "@tanstack/react-router";
import {useQuery} from "@tanstack/react-query";
import {Link} from "@/core/utils/router.ts";
import {ArticleOverviewItem} from "@/core/components/article/ArticleOverviewItem.tsx";
import type {Article} from "@/core/models/article.ts";
import {ArticleOverviewItemStatus} from "@/core/components/article/ArticleOverviewItemStatus.tsx";
import {Spinner} from "@/core/shadcn/components/ui/spinner.tsx";

export const Route = createFileRoute('/redactor/articles/')({
    component: RouteComponent,
});

function RouteComponent() {
    const redactorClient = Route.useRouteContext().redactorClient;

    const {queryKey, queryFn} = redactorClient.getRedactorArticles();

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

    return <div className="flex flex-col items-center gap-8 min-w-0 py-2">
        <Link
            to="/redactor/articles/new" replace={true}
            className="p-2 text-sm font-medium rounded-md hover:bg-accent"
        >
            Ajouter un article
        </Link>

        {
            (articles && articles.length > 0) &&

            <div className="flex flex-wrap justify-center gap-8 w-full">
                {articles.map((article) => (
                    <Link
                        key={article.id}
                        to="/redactor/articles/$articleRef"
                        params={{articleRef: article.reference}}
                        className="contents">
                        <ArticleOverviewItem
                            article={article}
                            header={ArticleHeader}
                        />
                    </Link>
                ))}
            </div>
        }

    </div>;
}

function ArticleHeader(article: Article) {
    return <ArticleOverviewItemStatus status={article.status}/>;
}
