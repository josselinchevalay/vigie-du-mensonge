import {createFileRoute} from '@tanstack/react-router';
import {RedactorArticleForm} from "@/core/components/redactor/RedactorArticleForm.tsx";
import {navigate} from "@/core/utils/router.ts";

export const Route = createFileRoute('/redactor/articles/new')({
    component: RouteComponent,
});

function RouteComponent() {
    const redactorClient = Route.useRouteContext().redactorClient;
    return <RedactorArticleForm redactorClient={redactorClient}
                                onSubmitSuccess={onSubmitSuccess}/>;
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
