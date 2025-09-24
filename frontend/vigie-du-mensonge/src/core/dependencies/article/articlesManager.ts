import {ArticleClient} from "@/core/dependencies/article/articleClient.ts";
import {Store} from "@tanstack/react-store";
import type {Article} from "@/core/models/article.ts";

class ArticlesManager {
    private readonly client = new ArticleClient();
    public readonly articlesStore = new Store<Article[]>([]);
    public readonly errStore = new Store(false);

    constructor() {
        void this.init();
    }

    async init() {
        try {
            const articles = await this.client.getAll();
            this.articlesStore.setState(() => articles);
        } catch {
            this.errStore.setState(() => true);
        }
    }
}

export const articlesManager = new ArticlesManager();