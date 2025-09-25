import {type RedactorArticleClient} from "@/core/dependencies/redactor/redactorArticleClient.ts";
import type {Article} from "@/core/models/article";
import {Store} from "@tanstack/react-store";

export class RedactorArticlesController {
    private readonly client: RedactorArticleClient;
    public readonly articlesStore = new Store<Article[]>([]);
    public readonly errStore = new Store(false);
    private initialized = false;

    constructor(client: RedactorArticleClient) {
        this.client = client;
        void this.init();
    }

    private async init() {
        if (this.initialized) return;

        try {
            const articles = await this.client.getArticles();
            this.articlesStore.setState(() => articles);
            this.initialized = true;
        } catch {
            this.errStore.setState(() => true);
        }
    }
}