import type {Article} from "@/core/models/article.ts";
import {ArticleCategoryLabels} from "@/core/models/articleCategory.ts";
import {fmtDate} from "@/core/utils/fmtDate.ts";

export function ArticleDisplay({article}: { article: Article }) {
    return (
        <article className="mx-auto w-full max-w-5xl px-4 sm:px-6 lg:px-8 py-6" aria-label={article.title}>
            <header className="mb-6">
                <div className="mb-3 flex flex-wrap items-center gap-2 text-sm text-gray-600">
                    <span className="text-sm font-bold">{ArticleCategoryLabels[article.category]}</span>
                    <span aria-label="date de l'événement">le {fmtDate(article.eventDate)}</span>
                    <span className="hidden sm:inline" aria-hidden>
                        ·
                    </span>
                    <span className="text-xs sm:text-sm text-gray-500" aria-label="dernière mise à jour">
                        Mis à jour le {new Intl.DateTimeFormat("fr-FR", {dateStyle: "medium"}).format(article.updatedAt)}
                    </span>
                </div>
                <h1 className="text-2xl sm:text-3xl md:text-4xl font-bold leading-tight text-gray-900">
                    {article.title}
                </h1>
                {article.politicians?.length ? (
                    <div className="mt-3 flex flex-wrap gap-2">
                        {article.politicians.map((pol, idx) => (
                            <span
                                key={`${pol.id}-${idx}`}
                                className="inline-block rounded-md border px-2 py-0.5 text-xs text-gray-700 bg-gray-50"
                            >
                                {pol.fullName}
                            </span>
                        ))}
                    </div>
                ) : null}
            </header>

            <div className="grid grid-cols-1 md:grid-cols-12 gap-6">
                <section className="md:col-span-8">
                    {article.body ? (
                        <div className="prose prose-sm sm:prose sm:max-w-none text-gray-800 whitespace-pre-line">
                            {article.body}
                        </div>
                    ) : (
                        <p className="text-gray-500 italic">Aucun contenu disponible.</p>
                    )}
                </section>

                <aside className="md:col-span-4">
                    {article.tags?.length ? (
                        <div className="mb-6">
                            <h2 className="mb-2 text-sm font-semibold text-gray-700">Tags</h2>
                            <div className="flex flex-wrap gap-2">
                                {article.tags.map((tag, idx) => (
                                    <span
                                        key={`${tag}-${idx}`}
                                        className="inline-block rounded-md border px-2 py-0.5 text-xs text-gray-700 bg-gray-50"
                                    >
                                        {tag}
                                    </span>
                                ))}
                            </div>
                        </div>
                    ) : null}

                    {article.sources?.length ? (
                        <div className="mb-2">
                            <h2 className="mb-2 text-sm font-semibold text-gray-700">Sources</h2>
                            <ul className="space-y-2">
                                {article.sources.map((src, idx) => (
                                    <li key={`${src}-${idx}`} className="break-words">
                                        <a
                                            href={src}
                                            target="_blank"
                                            rel="noopener noreferrer"
                                            className="text-blue-600 hover:text-blue-700 hover:underline"
                                        >
                                            {src}
                                        </a>
                                    </li>
                                ))}
                            </ul>
                        </div>
                    ) : null}
                </aside>
            </div>
        </article>
    );
}