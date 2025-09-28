import type {Article} from "@/core/models/article.ts";
import {ArticleOverviewItem} from "@/core/components/article/ArticleOverviewItem.tsx";
import React from "react";

export type ArticleOverviewGridProps = {
    articles: Article[];
    articleHeader?: (article: Article) => React.ReactNode;
    className?: string;
};

export function ArticleOverviewGrid({articles, articleHeader, className}:
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
                        article={article}
                        header={articleHeader}
                        className="sm:w-[30rem]"
                    />
                ))}
            </div>
        </div>
    );
}