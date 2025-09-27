import {File, FileCheck, FileClock, FileQuestionMark, FileWarning, FileX} from "lucide-react";
import {type ArticleStatus, ArticleStatusLabels} from "@/core/models/articleStatus.ts";
import React from "react";

const statusConfig: Record<ArticleStatus, { icon: React.ReactNode; label: string }> = {
    UNDER_REVIEW: {icon: <FileClock/>, label: ArticleStatusLabels.UNDER_REVIEW},
    CHANGE_REQUESTED: {icon: <FileWarning/>, label: ArticleStatusLabels.CHANGE_REQUESTED},
    PUBLISHED: {icon: <FileCheck/>, label: ArticleStatusLabels.PUBLISHED},
    ARCHIVED: {icon: <FileX/>, label: ArticleStatusLabels.ARCHIVED},
    DRAFT: {icon: <File/>, label: ArticleStatusLabels.DRAFT},
};

export function ArticleStatusDisplay({status}: { status: ArticleStatus; }) {

    const {icon, label} = statusConfig[status] ?? {icon: <FileQuestionMark/>, label: "Inconnu"};

    return (
        <div className="flex flex-row items-center gap-2 mb-2">
            {icon}
            <span className="text-sm text-muted-foreground">{label}</span>
        </div>
    );
}