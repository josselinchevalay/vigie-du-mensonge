import type {KyInstance} from "ky";
import {Article, type ArticleJson} from "@/core/models/article.ts";
import {api} from "@/core/dependencies/api.ts";

export type RedactorArticleJson = {
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

    async getArticles(): Promise<Article[]> {
        const res = await this.api
            .get("redactor/articles")
            .json<ArticleJson[]>();

        return res.map((json) => Article.fromJson(json));
    }

    async createDraftArticle(dto: RedactorArticleJson) {
        await this.api.post("redactor/articles", {json: dto});
    }
}

export const meClient = new RedactorClient(api);