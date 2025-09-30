import {Article, type ArticleJson} from "@/core/models/article.ts";
import type {KyInstance} from "ky";
import {api} from "@/core/dependencies/api.ts";

export class ArticleClient {
    private readonly api: KyInstance;

    constructor(api: KyInstance) {
        this.api = api;
    }

    private async _getAll(): Promise<Article[]> {
        const res = await this.api
            .get("articles")
            .json<ArticleJson[]>();

        return res.map((json) => Article.fromJson(json));
    }

    getAll = (): { queryKey: string[], queryFn: () => Promise<Article[]> } => {
        return {
            queryKey: ["articles"],
            queryFn: () => this._getAll(),
        };
    };

    private async _findById(id: string): Promise<Article> {
        const res = await this.api
            .get(`articles/${id}`)
            .json<ArticleJson>();

        return Article.fromJson(res);
    }

    findById = (id: string): { queryKey: string[], queryFn: () => Promise<Article> } => {
        return {
            queryKey: ["articles", id],
            queryFn: () => this._findById(id),
        };
    };
}

export const articleClient = new ArticleClient(api);