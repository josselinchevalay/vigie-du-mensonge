import {Link} from "@/core/utils/router.ts";
import {ArticleCardWrap} from "@/core/components/article/ArticleCardWrap.tsx";
import {SquarePen} from "lucide-react";
import type {Article} from "@/core/models/article.ts";

export type RedactorArticlesProps = {
    articles: Article[];
}

export function RedactorArticles({articles}: RedactorArticlesProps) {
    return (
        <div className="flex flex-col items-center gap-8 min-w-0 py-2">
            <Link
                to="/redactor/new-article"
                className="inline-flex items-center justify-center rounded-md border px-3 py-2 text-sm font-medium hover:bg-accent"
            >
                Ajouter un article
            </Link>
            <ArticleCardWrap articles={articles} showArticleStatus={true}
                             articleNavButton={(article) => RedactorArticleNavButton({articleId: article.id})}
            />
        </div>
    );
}

function RedactorArticleNavButton(props: { articleId: string }) {
    return (
        <Link
            to="/redactor/edit-article/$articleId"
            params={{articleId: props.articleId}}
            className="inline-flex items-center"
        >
            <SquarePen/>
        </Link>
    );
}