import type {KyInstance} from "ky";
import {Article, type ArticleJson} from "@/core/models/article.ts";
import type {ArticleStatus} from "@/core/models/articleStatus.ts";

export type ModeratorReview = {
    decision: ArticleStatus;
    notes?: string;
}

export class ModeratorClient {
    private readonly api: KyInstance;

    constructor(api: KyInstance) {
        this.api = api;
    }

    async getModeratorArticles(): Promise<Article[]> {
        const res = await this.api
            .get("moderator/articles")
            .json<ArticleJson[]>();

        return res.map((json) => Article.fromJson(json));
    }

    async getPendingArticles(): Promise<Article[]> {
        const res = await this.api
            .get("moderator/articles/pending")
            .json<ArticleJson[]>();

        return res.map((json) => Article.fromJson(json));
    }

    async findModeratorArticlesByRef(ref: string): Promise<Article[]> {
        const res = await this.api
            .get(`moderator/articles/${ref}`)
            .json<ArticleJson[]>();

        return res.map((json) => Article.fromJson(json));
    }

    async claimModeratorArticle(articleID: string): Promise<void> {
        await this.api.post(`moderator/articles/${articleID}/claim`);
    }

    async reviewModeratorArticle(articleID: string, review: ModeratorReview): Promise<void> {
        await this.api.post(`moderator/articles/${articleID}/review`, {json: review});
    }
}