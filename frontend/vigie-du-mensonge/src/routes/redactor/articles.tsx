import {createFileRoute} from "@tanstack/react-router";
import {RedactorArticles} from "@/core/components/redactor/RedactorArticles.tsx";

export const Route = createFileRoute('/redactor/articles')({
    component: RedactorArticles,
});
