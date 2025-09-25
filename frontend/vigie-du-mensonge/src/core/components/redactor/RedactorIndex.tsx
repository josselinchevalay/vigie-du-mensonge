import type {RedactorController} from "@/core/dependencies/redactor/redactorController.ts";
import {useStore} from "@tanstack/react-store";
import {ArticleCardWrap} from "@/core/components/article/ArticleCardWrap.tsx";

export type MeProps = {
    controller: RedactorController;
}

export function RedactorIndex({controller}: MeProps) {
    const articles = useStore(controller.articlesStore);
    return <ArticleCardWrap articles={articles}/>;
}