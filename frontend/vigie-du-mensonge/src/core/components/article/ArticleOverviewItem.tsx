import type {Article} from "@/core/models/article.ts";
import React from "react";
import {ArticleCategoryLabels} from "@/core/models/articleCategory.ts";
import {fmtDate} from "@/core/utils/fmtDate.ts";

export type ArticleOverviewItemProps = {
    article: Article;
    header?: (article: Article) => React.ReactNode;
};

export function ArticleOverviewItem({article, header}: ArticleOverviewItemProps) {
    return (
        <div
            className="flex flex-col gap-1 w-4/5 sm:w-1/3 lg:w-1/4 p-2 rounded-lg border bg-secondary text-secondary-foreground shadow-xl hover:shadow-2xl"
            role="article"
            aria-label={article.title}>

            <div className="flex justify-center">
                {header && header(article)}
            </div>

            <div className="flex flex-row items-center justify-center gap-2 my-1">
                <span className="text-sm font-bold">{ArticleCategoryLabels[article.category]}</span>
                <span className="text-sm">le {fmtDate(article.eventDate)}</span>
            </div>

            <h3 className="text-lg font-bold overflow-hidden text-clip">{article.title}</h3>

            {article.politicians?.length ? (
                <div className="mt-2 flex flex-wrap gap-2">
                    {article.politicians.map((pol, idx) => (
                        <span
                            key={`${pol.fullName}-${idx}`}
                            className="inline-block rounded-md border px-2 py-0.5 text-sm text-foreground bg-background">
                            {pol.fullName}
                        </span>
                    ))}
                </div>
            ) : null}

            {article.tags?.length ? (
                <div className="mt-2 flex flex-wrap gap-2">
                    {article.tags.map((tag, idx) => (
                        <span
                            key={`${tag}-${idx}`}
                            className="inline-block rounded-md border px-2 py-0.5 text-xs text-foreground bg-background"
                        >
                            {tag}
                        </span>
                    ))}
                </div>
            ) : null}

        </div>
    );
}