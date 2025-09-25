import type {KyInstance} from "ky";
import {Article, type ArticleDetails, type ArticleJson} from "@/core/models/article.ts";
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

export class RedactorArticleClient {
    private readonly api: KyInstance;

    constructor(api: KyInstance) {
        this.api = api;
    }

    async getAll(): Promise<Article[]> {
        const res = await this.api
            .get("redactor/articles")
            .json<ArticleJson[]>();

        return res.map((json) => Article.fromJson(json));
    }

    async create(dto: RedactorArticleJson): Promise<string> {
        const res = await this.api
            .post("redactor/articles", {json: dto})
            .json<{ articleID: string }>();

        return res.articleID;
    }

    async findDetails(articleID: string): Promise<ArticleDetails> {
        return await this.api
            .get(`redactor/articles/${articleID}/details`)
            .json<ArticleDetails>();
    }

    async update(articleID: string, dto: RedactorArticleJson): Promise<void> {
        await this.api.put(`redactor/articles/${articleID}`, {json: dto});
    }
}

export const redactorArticleClient = new RedactorArticleClient(api);