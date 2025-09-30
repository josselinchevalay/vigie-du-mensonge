import type {KyInstance} from "ky";
import {Article, type ArticleJson} from "@/core/models/article.ts";
import type {ArticleStatus} from "@/core/models/articleStatus.ts";

export type SaveModeratorReview = {
    decision: ArticleStatus;
    notes?: string;
}

export class ModeratorClient {
    private readonly api: KyInstance;

    constructor(api: KyInstance) {
        this.api = api;
    }

    private async _getModeratorArticles(): Promise<Article[]> {
        const res = await this.api
            .get("moderator/articles")
            .json<ArticleJson[]>();

        return res.map((json) => Article.fromJson(json));
    }

    getModeratorArticles = (): { queryKey: string[], queryFn: () => Promise<Article[]> } => {
        return {
            queryKey: ["moderator", "articles"],
            queryFn: () => this._getModeratorArticles(),
        };
    };

    private async _getPendingArticles(): Promise<Article[]> {
        const res = await this.api
            .get("moderator/articles/pending")
            .json<ArticleJson[]>();

        return res.map((json) => Article.fromJson(json));
    }

    getPendingArticles = (): { queryKey: string[], queryFn: () => Promise<Article[]> } => {
        return {
            queryKey: ["moderator", "articles", "pending"],
            queryFn: () => this._getPendingArticles(),
        };
    };

    private async _findModeratorArticlesByRef(ref: string): Promise<Article[]> {
        const res = await this.api
            .get(`moderator/articles/${ref}`)
            .json<ArticleJson[]>();

        return res.map((json) => Article.fromJson(json));
    }

    findModeratorArticlesByRef = (ref: string): { queryKey: string[], queryFn: () => Promise<Article[]> } => {
        return {
            queryKey: ["moderator", "articles", ref],
            queryFn: () => this._findModeratorArticlesByRef(ref),
        };
    };

    async claimModeratorArticle(articleID: string): Promise<void> {
        await this.api.post(`moderator/articles/${articleID}/claim`);
    }

    async saveModeratorReview(articleID: string, review: SaveModeratorReview): Promise<void> {
        await this.api.post(`moderator/articles/${articleID}/review`, {json: review});
    }
}