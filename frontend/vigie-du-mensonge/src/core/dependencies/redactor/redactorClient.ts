import type {KyInstance} from "ky";
import {Article, type ArticleJson} from "@/core/models/article.ts";

export type SaveRedactorArticle = {
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

    private async _getRedactorArticles(): Promise<Article[]> {
        const res = await this.api
            .get("redactor/articles")
            .json<ArticleJson[]>();

        return res.map((json) => Article.fromJson(json));
    }

    getRedactorArticles = (): { queryKey: string[], queryFn: () => Promise<Article[]> } => {
        return {
            queryKey: ["redactor", "articles"],
            queryFn: () => this._getRedactorArticles(),
        };
    };

    async saveArticle(publish: boolean, dto: SaveRedactorArticle): Promise<string> {
        const res = await this.api
            .post(`redactor/articles?publish=${publish}`, {json: dto})
            .json<{ articleReference: string }>();

        return res.articleReference;
    }

    private async _findRedactorArticlesByRef(ref: string): Promise<Article[]> {
        const res = await this.api
            .get(`redactor/articles/${ref}`)
            .json<ArticleJson[]>();

        return res.map((json) => Article.fromJson(json));
    }

    findRedactorArticlesByRef = (ref: string): { queryKey: string[], queryFn: () => Promise<Article[]> } => {
        return {
            queryKey: ["redactor", "articles", ref],
            queryFn: () => this._findRedactorArticlesByRef(ref),
        };
    };
}
