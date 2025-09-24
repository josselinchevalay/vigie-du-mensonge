import {Article, type ArticleJson} from "@/core/models/article.ts";
import {api} from "@/core/dependencies/api.ts";

export class ArticleClient {
    async getAll(): Promise<Article[]> {
        const res = await api
            .get("articles")
            .json<ArticleJson[]>();

        return res.map((json) => Article.fromJson(json));
    }
}
