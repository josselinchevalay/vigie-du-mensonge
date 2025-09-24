export const ArticleStatuses = {
    UNDER_REVIEW: 'UNDER_REVIEW',
    CHANGE_REQUESTED: 'CHANGE_REQUESTED',
    DRAFT: 'DRAFT',
    ARCHIVED: 'ARCHIVED',
    PUBLISHED: 'PUBLISHED',
};

export type ArticleStatus = keyof typeof ArticleStatuses;

export const ArticleStatusLabels: Record<ArticleStatus, string> = {
    UNDER_REVIEW: 'Examen en cours',
    CHANGE_REQUESTED: 'Changements demandés',
    DRAFT: 'Brouillon',
    ARCHIVED: 'Archivé',
    PUBLISHED: 'Publié',
};

