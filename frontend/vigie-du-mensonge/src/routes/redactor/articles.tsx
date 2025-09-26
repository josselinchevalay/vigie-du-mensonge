import {createFileRoute} from "@tanstack/react-router";
import {RedactorArticles} from "@/core/components/redactor/RedactorArticles.tsx";
import {useQuery} from "@tanstack/react-query";
import {BasicProgress} from "@/core/components/BasicProgress.tsx";

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

    return <RedactorArticles articles={articles!}/>;
}