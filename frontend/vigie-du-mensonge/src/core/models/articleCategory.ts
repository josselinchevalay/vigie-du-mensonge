export const ArticleCategoryValues = {
    FALSEHOOD: 'FALSEHOOD',
    LIE: 'LIE',
} as const;

export type ArticleCategory = keyof typeof ArticleCategoryValues;

export const ArticleCategoryLabels: Record<ArticleCategory, string> = {
    FALSEHOOD: 'Contre-vérité',
    LIE: 'Mensonge',
};