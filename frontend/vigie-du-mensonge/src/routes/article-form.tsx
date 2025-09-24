import {createFileRoute} from "@tanstack/react-router";

export const Route = createFileRoute('/article-form')({
    component: RouteComponent,
});

function RouteComponent() {
    return <div>ArticleForm</div>;
}