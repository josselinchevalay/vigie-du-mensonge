import {redactorArticleClient, type RedactorArticleClient} from "@/core/dependencies/redactor/redactorArticleClient.ts";
import type {Article} from "@/core/models/article";
import {Store} from "@tanstack/react-store";

export class RedactorArticlesManager {
    private readonly client: RedactorArticleClient;
    public readonly articlesStore = new Store<Article[]>([]);
    public readonly errStore = new Store(false);
    private loading = false;

    constructor(client: RedactorArticleClient) {
        this.client = client;
    }

    public async loadArticles() {
        if (this.loading) {
            return;
        }

        this.loading = true;
        try {
            const articles = await this.client.getAll();
            this.articlesStore.setState(() => articles);
        } catch {
            this.errStore.setState(() => true);
        } finally {
            this.loading = false;
        }
    }
}

export const redactorArticlesManager = new RedactorArticlesManager(redactorArticleClient);