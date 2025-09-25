import type {RedactorClient} from "@/core/dependencies/redactor/redactorClient.ts";
import type {Article} from "@/core/models/article";
import {Store} from "@tanstack/react-store";

export class RedactorController {
    private readonly client: RedactorClient;
    public readonly articlesStore = new Store<Article[]>([]);
    public readonly errStore = new Store(false);

    constructor(client: RedactorClient) {
        this.client = client;
        void this.init();
    }

    private async init() {
        try {
            const articles = await this.client.getArticles();
            this.articlesStore.setState(() => articles);
        } catch {
            this.errStore.setState(() => true);
        }
    }
}