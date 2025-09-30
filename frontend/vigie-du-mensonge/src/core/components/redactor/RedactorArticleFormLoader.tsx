import {RedactorClient} from "@/core/dependencies/redactor/redactorClient.ts";
import type {Article} from "@/core/models/article.ts";
import type {PoliticianClient} from "@/core/dependencies/politician/politicianClient.ts";
import {useQuery} from "@tanstack/react-query";
import {Spinner} from "@/core/shadcn/components/ui/spinner.tsx";
import {RedactorArticleForm} from "@/core/components/redactor/RedactorArticleForm.tsx";

export type RedactorArticleFormLoaderProps = {
    politicianClient: PoliticianClient
    redactorClient: RedactorClient
    article?: Article
    onSubmitSuccess?: (articleRef?: string) => void
};

export function RedactorArticleFormLoader({politicianClient, redactorClient, article, onSubmitSuccess}:
                                          RedactorArticleFormLoaderProps) {
    const {queryKey, queryFn} = politicianClient.getAll();

    const {data: politicians, isLoading, isError} = useQuery({
        queryKey: queryKey,
        queryFn: () => queryFn(),
        staleTime: 24 * 60 * 60 * 1000,
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

    return <RedactorArticleForm redactorClient={redactorClient} politicians={politicians ?? []} article={article}
                                onSubmitSuccess={onSubmitSuccess}/>;
}