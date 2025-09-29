import {Article, type ArticleJson} from "@/core/models/article.ts";
import type {KyInstance} from "ky";
import {api} from "@/core/dependencies/api.ts";

export class ArticleClient {
    private readonly api: KyInstance;

    constructor(api: KyInstance) {
        this.api = api;
    }

    async getAll(): Promise<Article[]> {
        const res = await this.api
            .get("articles")
            .json<ArticleJson[]>();

        return res.map((json) => Article.fromJson(json));
    }

    async findById(id: string): Promise<Article> {
        const res = await this.api
            .get(`articles/${id}`)
            .json<ArticleJson>();

        return Article.fromJson(res);
    }
}

export const articleClient = new ArticleClient(api);