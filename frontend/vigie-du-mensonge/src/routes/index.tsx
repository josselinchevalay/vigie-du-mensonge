import {createFileRoute} from '@tanstack/react-router';
import {useQuery} from "@tanstack/react-query";
import {Spinner} from "@/core/shadcn/components/ui/spinner.tsx";
import {Link} from "@/core/utils/router.ts";
import {ArticleOverviewItem} from "@/core/components/article/ArticleOverviewItem.tsx";
import {Separator} from "@/core/shadcn/components/ui/separator.tsx";

export const Route = createFileRoute('/')({
    component: RouteComponent,
});

function RouteComponent() {
    const articleClient = Route.useRouteContext().articleClient;

    const {queryKey, queryFn} = articleClient.getAll();

    const {data: articles, isLoading, isError} = useQuery({
        queryKey: queryKey,
        queryFn: () => queryFn(),
        staleTime: 24 * 60 * 60 * 1000,
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

    return <div className="flex flex-col items-center gap-4 min-w-0 py-2">

        <div className="flex flex-col gap-4 max-w-7xl px-4">
            <p>
                Ce site a pour mission de recenser, analyser et documenter les contre-vérités, manipulations et discours
                officiels souvent contestables des gouvernements successifs sous la présidence d’Emmanuel Macron. Notre
                objectif est d’offrir aux citoyens une source fiable, factuelle et accessible pour comprendre les enjeux
                réels derrière les déclarations publiques, afin de favoriser un débat éclairé et critique. Ici, vous
                trouverez des entrées détaillées accompagnées de citations, faits et commentaires pour décrypter les
                discours politiques et mieux armer votre regard face à la communication gouvernementale.
            </p>
            <Separator/>
        </div>

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