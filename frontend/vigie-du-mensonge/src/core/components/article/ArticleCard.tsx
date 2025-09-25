import type {Article} from "@/core/models/article.ts";

export type ArticleCardProps = {
    article: Article;
    className?: string;
};

export function ArticleCard({article, className}: ArticleCardProps) {
    return (
        <div className={["rounded-lg border bg-white p-4 shadow-sm", className].filter(Boolean).join(" ")}
             role="article"
             aria-label={article.title}
        >
            <div className="mb-2">
                <h3 className="text-lg font-bold leading-snug line-clamp-2">{article.title}</h3>
            </div>
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