import type {RedactorController} from "@/core/dependencies/redactor/redactorController.ts";
import {useStore} from "@tanstack/react-store";
import {ArticleCardWrap} from "@/core/components/article/ArticleCardWrap.tsx";
import {Link} from "@/core/utils/router.ts";

export type RedactorIndexProps = {
    controller: RedactorController;
}

export function RedactorIndex({controller}: RedactorIndexProps) {
    const articles = useStore(controller.articlesStore);
    return (
        <div className="flex flex-col items-center gap-8 min-w-0 py-2">
            <Link
                to="/redactor/article-form"
                className="inline-flex items-center rounded-md border px-3 py-2 text-sm font-medium hover:bg-accent"
            >
                Ajouter un article
            </Link>
            <ArticleCardWrap articles={articles}/>
        </div>
    );
}