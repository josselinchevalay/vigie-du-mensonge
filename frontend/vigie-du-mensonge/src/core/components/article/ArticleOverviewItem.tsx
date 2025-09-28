import type {Article} from "@/core/models/article.ts";
import React from "react";

export type ArticleOverviewItemProps = {
    article: Article;
    header?: (article: Article) => React.ReactNode;
    className?: string;
};

export function ArticleOverviewItem({article, header, className}: ArticleOverviewItemProps) {
    return (
        <div className={["rounded-lg border bg-white p-2 shadow-sm", className].filter(Boolean).join(" ")}
             role="article"
             aria-label={article.title}
        >
            <div className="flex flex-row items-center justify-center gap-2 mb-2">
                {header && header(article)}
            </div>
            <h3 className="text-lg font-bold leading-snug line-clamp-2">{article.title}</h3>
            {article.politicians?.length ? (
                <div className="mt-2 flex flex-wrap gap-2">
                    {article.politicians.map((pol, idx) => (
                        <span
                            key={`${pol.fullName}-${idx}`}
                            className="inline-block rounded-md border px-2 py-0.5 text-xs text-gray-700 bg-gray-50"
                        >
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
                            className="inline-block rounded-md border px-2 py-0.5 text-xs text-gray-700 bg-gray-50"
                        >
                            {tag}
                        </span>
                    ))}
                </div>
            ) : null}
        </div>
    );
}