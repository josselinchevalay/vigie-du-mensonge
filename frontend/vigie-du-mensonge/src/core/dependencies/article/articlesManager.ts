import {Store} from "@tanstack/react-store";
import type {Article} from "@/core/models/article.ts";
import {ArticleClient, articleClient} from "@/core/dependencies/article/articleClient.ts";

class ArticlesManager {
    private readonly client: ArticleClient;

    public readonly articlesStore = new Store<Article[]>([]);
    public readonly errStore = new Store(false);

    constructor(client: ArticleClient) {
        this.client = client;
        void this.init();
    }

    private async init() {
        try {
            const articles = await this.client.getAll();
            this.articlesStore.setState(() => articles);
        } catch {
            this.errStore.setState(() => true);
        }
    }
}

export const articlesManager = new ArticlesManager(articleClient);