import type {Article} from "@/core/models/article.ts";
import {ArticleCard} from "@/core/components/article/ArticleCard.tsx";
import React from "react";

export type ArticleCardWrapProps = {
    articles: Article[];
    articleNavButton?: (article: Article) => React.ReactNode;
    className?: string;
};

export function ArticleCardWrap({articles, articleNavButton, className}: ArticleCardWrapProps) {
    return (
        <div className={["max-h-[70vh] overflow-auto", className].filter(Boolean).join(" ")}
             role="list"
             aria-label="articles"
        >
            <div className="flex flex-wrap gap-4 justify-center">
                {articles?.map((article) => (
                    <ArticleCard
                        key={article.id}
                        navButton={articleNavButton}
                        article={article}
                        className="w-full sm:w-[calc(50%-0.5rem)] lg:w-[calc(33.333%-0.666rem)]"
                    />
                ))}
            </div>
        </div>
    );
}