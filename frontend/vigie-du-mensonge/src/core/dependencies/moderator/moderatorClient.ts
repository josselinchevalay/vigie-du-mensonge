import type {KyInstance} from "ky";
import {Article, type ArticleJson} from "@/core/models/article.ts";

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
}