import type {Article} from "@/core/models/article.ts";
import {ArticleOverviewItem} from "@/core/components/article/ArticleOverviewItem.tsx";
import React from "react";

export type ArticleOverviewGridProps = {
    articles: Article[];
    articleNavButton?: (article: Article) => React.ReactNode;
    showArticleStatus?: boolean;
    className?: string;
};

export function ArticleOverviewGrid({articles, articleNavButton, showArticleStatus, className}:
                                    ArticleOverviewGridProps) {
    return (
        <div
            className={["max-h-[70vh] overflow-auto", className].filter(Boolean).join(" ")}
            role="list"
            aria-label="articles"
        >
            <div className="flex flex-wrap justify-center gap-4">
                {articles.map((article) => (
                    <ArticleOverviewItem
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