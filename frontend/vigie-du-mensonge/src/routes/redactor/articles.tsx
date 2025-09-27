import {createFileRoute} from "@tanstack/react-router";
import {useQuery} from "@tanstack/react-query";
import {BasicProgress} from "@/core/components/BasicProgress.tsx";
import {Link} from "@/core/utils/router.ts";
import {SquarePen} from "lucide-react";
import {ArticleOverviewGrid} from "@/core/components/article/ArticleOverviewGrid.tsx";

export const Route = createFileRoute('/redactor/articles')({
    component: RouteComponent,
});

function RouteComponent() {
    const redactorClient = Route.useRouteContext().redactorClient;

    const {data: articles, isLoading, isError} = useQuery({
        queryKey: ["redactor", "articles"],
        queryFn: () => redactorClient.getAllArticles(),
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
            to="/redactor/new-article"
            className="inline-flex items-center justify-center rounded-md border px-3 py-2 text-sm font-medium hover:bg-accent"
        >
            Ajouter un article
        </Link>
        <ArticleOverviewGrid articles={articles!} showArticleStatus={true}
                             articleNavButton={(article) => RedactorArticleNavButton({articleRef: article.reference})}
        />
    </div>;
}

function RedactorArticleNavButton(props: { articleRef: string }) {
    return (
        <Link
            to="/redactor/edit-article/$articleRef"
            params={{articleRef: props.articleRef}}
            className="inline-flex items-center"
        >
            <SquarePen/>
        </Link>
    );
}