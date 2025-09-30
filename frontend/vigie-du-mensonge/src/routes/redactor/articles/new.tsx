import {createFileRoute} from '@tanstack/react-router';
import {navigate} from "@/core/utils/router.ts";
import {RedactorArticleFormLoader} from "@/core/components/redactor/RedactorArticleFormLoader.tsx";

export const Route = createFileRoute('/redactor/articles/new')({
    component: RouteComponent,
});

function RouteComponent() {
    const redactorClient = Route.useRouteContext().redactorClient;
    const politicianClient = Route.useRouteContext().politicianClient;
    return <RedactorArticleFormLoader redactorClient={redactorClient}
                                      onSubmitSuccess={onSubmitSuccess} politicianClient={politicianClient}/>;
}

async function onSubmitSuccess(articleRef?: string) {
    if (articleRef) {
        await navigate({
            to: "/redactor/articles/$articleRef",
            params: {articleRef: articleRef},
            replace: true
        });
    }
}
