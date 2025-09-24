export const ArticleCategories = {
    FALSEHOOD: 'FALSEHOOD',
    LIE: 'LIE',
} as const;

export type ArticleCategory = keyof typeof ArticleCategories;

export const ArticleCategoryLabels: Record<ArticleCategory, string> = {
    FALSEHOOD: 'Contre-vérité',
    LIE: 'Mensonge',
};