import {createFileRoute} from '@tanstack/react-router';
import {useQuery} from "@tanstack/react-query";
import {Spinner} from "@/core/shadcn/components/ui/spinner.tsx";
import {Link} from "@/core/utils/router.ts";
import {ArticleOverviewItem} from "@/core/components/article/ArticleOverviewItem.tsx";

export const Route = createFileRoute('/')({
    component: RouteComponent,
});

function RouteComponent() {
    const articleClient = Route.useRouteContext().articleClient;

    const {data: articles, isLoading, isError} = useQuery({
        queryKey: ["articles"],
        queryFn: () => articleClient.getAll(),
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

        {
            (articles && articles.length > 0) &&

            <div className="flex flex-wrap justify-center gap-8 w-full">
                {articles.map((article) => (
                    <Link
                        key={article.id}
                        to="/articles/$articleId"
                        params={{articleId: article.id}}
                        className="contents">
                        <ArticleOverviewItem article={article}/>
                    </Link>
                ))}
            </div>
        }

    </div>;
}