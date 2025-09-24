import type {KyInstance} from "ky";
import {Article, type ArticleJson} from "@/core/models/article.ts";

export class MeClient {
    private readonly api: KyInstance;

    constructor(api: KyInstance) {
        this.api = api;
    }

    async getArticles(): Promise<Article[]> {
        const res = await this.api
            .get("me/articles")
            .json<ArticleJson[]>();

        return res.map((json) => Article.fromJson(json));
    }
}