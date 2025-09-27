import type {KyInstance} from "ky";
import {Article, type ArticleJson} from "@/core/models/article.ts";

export type RedactorArticleJson = {
    id?: string;
    title: string;
    body: string;
    eventDate: Date;
    tags: string[];
    sources: string[];
    politicianIds: string[];
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

    async saveArticle(publish: boolean, dto: RedactorArticleJson): Promise<void> {
        await this.api.post(`redactor/articles?publish=${publish}`, {json: dto});
    }

    async findArticleByRef(ref: string): Promise<Article> {
        const res = await this.api
            .get(`redactor/articles/${ref}`)
            .json<ArticleJson>();

        return Article.fromJson(res);
    }
}
