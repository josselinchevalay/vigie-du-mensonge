import {createFileRoute} from '@tanstack/react-router';
import {IndexArticles} from "@/core/components/article/IndexArticles.tsx";

export const Route = createFileRoute('/')({
    component: IndexArticles,
});