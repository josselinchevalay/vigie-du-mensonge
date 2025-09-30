import {createFileRoute} from '@tanstack/react-router';
import {useQuery} from "@tanstack/react-query";
import {Spinner} from "@/core/shadcn/components/ui/spinner.tsx";
import {ArticleDisplay} from "@/core/components/article/ArticleDisplay.tsx";

export const Route = createFileRoute('/articles/$articleId')({
    component: RouteComponent,
});

function RouteComponent() {
    const {articleId} = Route.useParams();
    const articleClient = Route.useRouteContext().articleClient;

    const {queryKey, queryFn} = articleClient.findById(articleId);

    const {data: article, isLoading, isError} = useQuery({
        queryKey: queryKey,
        queryFn: () => queryFn(),
    });

    if (isError) {
        return <div className="flex items-center justify-center h-screen">
            Une erreur est survenue. Veuillez réessayer.
        </div>;
    }

    if (isLoading) {
        return (
            <div className="flex flex-col gap-2 items-center justify-center h-screen">
                Chargement en cours...
                <Spinner/>
            </div>
        );
    }

    if (!article) {
        return <div className="flex items-center justify-center h-screen">
            Une erreur est survenue. Veuillez réessayer.
        </div>;
    }

    return <div className="flex flex-col items-center">
        <ArticleDisplay article={article}/>
    </div>;
}
