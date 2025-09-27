import type {Article} from "@/core/models/article.ts";
import React from "react";
import {ArticleStatusDisplay} from "@/core/components/article/ArticleStatusDisplay.tsx";

export type ArticleCardProps = {
    article: Article;
    navButton?: (article: Article) => React.ReactNode;
    showStatus?: boolean;
    className?: string;
};

export function ArticleCard({article, navButton, showStatus, className}: ArticleCardProps) {
    return (
        <>

            <div className={["rounded-lg border bg-white p-2 shadow-sm", className].filter(Boolean).join(" ")}
                 role="article"
                 aria-label={article.title}
            >
                <div className="flex flex-row gap-2 mb-2">
                    {showStatus && ArticleStatusDisplay({status: article.status})}
                </div>
                <div className="flex flex-row gap-2 mb-2">
                    {navButton && navButton(article)}
                    <h3 className="text-lg font-bold leading-snug line-clamp-2">{article.title}</h3>
                </div>
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
        </>
    );
}