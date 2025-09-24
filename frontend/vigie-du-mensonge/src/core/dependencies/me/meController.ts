import type {MeClient} from "@/core/dependencies/me/meClient.ts";
import type {Article} from "@/core/models/article";
import {Store} from "@tanstack/react-store";

export class MeController {
    private readonly client: MeClient;
    public readonly articlesStore = new Store<Article[]>([]);
    public readonly errStore = new Store(false);

    constructor(client: MeClient) {
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