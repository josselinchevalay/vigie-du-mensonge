import type {ArticleReview} from "@/core/models/articleReview.ts";
import {type ArticleStatus, ArticleStatusLabels} from "@/core/models/articleStatus.ts";
import {File, FileCheck, FileClock, FileWarning, FileX, MessageSquareText} from "lucide-react";
import React from "react";

export type ArticleReviewCardProps = {
    review: ArticleReview;
};

const statusIcon: Record<ArticleStatus, React.ReactNode> = {
    UNDER_REVIEW: <FileClock/>,
    CHANGE_REQUESTED: <FileWarning/>,
    PUBLISHED: <FileCheck/>,
    ARCHIVED: <FileX/>,
    DRAFT: <File/>,
};

export function ArticleReviewCard({review}: ArticleReviewCardProps) {
    if (!review) return null;

    const decision = review.decision as ArticleStatus | undefined;

    return (
        <section
            className="w-full sm:max-w-2xl rounded-lg border bg-secondary/60 text-secondary-foreground shadow-sm p-3 sm:p-4"
            aria-label="Review du modérateur"
        >
            <header className="flex items-center justify-between gap-2 mb-3">
                <div className="flex items-center gap-2">
                    <MessageSquareText className="h-4 w-4"/>
                    <h3 className="text-sm font-semibold">Review du modérateur</h3>
                </div>

                {decision && (
                    <div className="flex items-center gap-2">
                        {statusIcon[decision]}
                        <span className="text-sm font-bold">
                            {ArticleStatusLabels[decision]}
                        </span>
                    </div>
                )}
            </header>

            {review.moderatorTag && (
                <div className="mb-2">
                    <span className="inline-flex items-center rounded-md border bg-background px-2 py-0.5 text-xs">
                        {review.moderatorTag}
                    </span>
                </div>
            )}

            {review.notes && (
                <div className="mt-1">
                    <p
                        className="text-sm leading-relaxed whitespace-pre-wrap break-words"
                        title={review.notes}
                    >
                        {review.notes}
                    </p>
                </div>
            )}
        </section>
    );
}