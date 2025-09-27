import type {Article} from "@/core/models/article.ts";
import {ArticleCard} from "@/core/components/article/ArticleCard.tsx";
import React from "react";

export type ArticleCardWrapProps = {
    articles: Article[];
    articleNavButton?: (article: Article) => React.ReactNode;
    showArticleStatus?: boolean;
    className?: string;
};

export function ArticleCardWrap({articles, articleNavButton, showArticleStatus, className}: ArticleCardWrapProps) {
    return (
        <div
            className={["max-h-[70vh] overflow-auto", className].filter(Boolean).join(" ")}
            role="list"
            aria-label="articles"
        >
            <div className="flex flex-wrap justify-center gap-4">
                {articles.map((article) => (
                    <ArticleCard
                        key={article.id}
                        navButton={articleNavButton}
                        article={article}
                        showStatus={showArticleStatus}
                        className="w-full sm:w-[20rem]" // 👈 fixed card width on sm+
                    />
                ))}
            </div>
        </div>
    );
}