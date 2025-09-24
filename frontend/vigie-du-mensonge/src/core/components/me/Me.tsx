import type {MeController} from "@/core/dependencies/me/meController.ts";
import {useStore} from "@tanstack/react-store";

export type MeProps = {
    controller: MeController;
}

export function Me({controller}: MeProps) {
    const articles = useStore(controller.articlesStore);
    return <div>articles: {articles.length}</div>;
}