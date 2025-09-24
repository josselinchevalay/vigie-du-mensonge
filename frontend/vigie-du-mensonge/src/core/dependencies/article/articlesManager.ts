import {Store} from "@tanstack/react-store";
import type {Article} from "@/core/models/article.ts";
import {articleClient} from "@/core/dependencies/article/articleClient.ts";

class ArticlesManager {
    public readonly articlesStore = new Store<Article[]>([]);
    public readonly errStore = new Store(false);

    constructor() {
        void this.init();
    }

    private async init() {
        try {
            const articles = await articleClient.getAll();
            this.articlesStore.setState(() => articles);
        } catch {
            this.errStore.setState(() => true);
        }
    }
}

export const articlesManager = new ArticlesManager();