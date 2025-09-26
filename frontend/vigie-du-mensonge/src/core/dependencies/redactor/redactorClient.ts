import type {KyInstance} from "ky";
import {Article, type ArticleJson} from "@/core/models/article.ts";

export type RedactorArticleJson = {
    id?: string;
    title: string;
    body: string;
    eventDate: Date;
    tags: string[];
    sources: string[];
    politicians: string[];
    category: string;
}

export class RedactorClient {
    private readonly api: KyInstance;

    constructor(api: KyInstance) {
        this.api = api;
    }

    async getAllArticles(): Promise<Article[]> {
        const res = await this.api
            .get("redactor/articles")
            .json<ArticleJson[]>();

        return res.map((json) => Article.fromJson(json));
    }

    async saveArticleDraft(dto: RedactorArticleJson): Promise<void> {
        await this.api.post("redactor/articles/drafts", {json: dto});
    }

    async createArticleDraft(dto: RedactorArticleJson): Promise<string> {
        const res = await this.api
            .post("redactor/articles", {json: dto})
            .json<{ articleID: string }>();

        return res.articleID;
    }

    async findArticleById(articleId: string): Promise<Article> {
        const res = await this.api
            .get(`redactor/articles/${articleId}`)
            .json<ArticleJson>();

        return Article.fromJson(res);
    }

    async updateArticle(articleId: string, dto: RedactorArticleJson): Promise<void> {
        await this.api.put(`redactor/articles/${articleId}`, {json: dto});
    }
}
