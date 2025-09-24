import {Article, type ArticleJson} from "@/core/models/article.ts";
import type {ArticleCategory} from "@/core/models/articleCategory.ts";
import type {KyInstance} from "ky";
import {api} from "@/core/dependencies/api.ts";

export type ArticleCreateJson = {
    title: string;
    body: string;
    eventDate: Date;
    tags: string[];
    sources: string[];
    politicians: string[];
    category: ArticleCategory;
}

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

    async create(dto: ArticleCreateJson) {
        await this.api.post("articles", {json: dto});
    }
}

export const articleClient = new ArticleClient(api.extend({prefixUrl: "/articles"}));