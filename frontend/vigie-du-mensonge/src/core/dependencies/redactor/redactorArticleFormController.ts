import {toast} from "@/core/utils/toast.ts";
import type {RedactorArticleClient, RedactorArticleJson} from "@/core/dependencies/redactor/redactorArticleClient.ts";
import {navigate} from "@/core/utils/router.ts";
import type {Article} from "@/core/models/article.ts";
import {Store} from "@tanstack/react-store";

export class RedactorArticleFormController {
    private readonly client: RedactorArticleClient;
    public readonly originalArticle: Article | null;
    public readonly loadingStore = new Store(true);
    public readonly errStore = new Store(false);

    constructor(client: RedactorArticleClient, originalArticle: Article | null) {
        this.client = client;
        this.originalArticle = originalArticle;

        if (originalArticle) {
            void this.loadOriginalArticleDetails();
        } else {
            this.loadingStore.setState(() => false);
        }
    }

    private async loadOriginalArticleDetails() {
        if (!this.originalArticle) return;

        try {
            this.originalArticle.details = await this.client.findDetails(this.originalArticle.id);
        } catch {
            this.errStore.setState(() => true);
        } finally {
            this.loadingStore.setState(() => false);
        }
    }

    async onSubmit(json: RedactorArticleJson): Promise<boolean> {
        try {
            if (this.originalArticle) {
                await this.client.update(this.originalArticle.id, json);
                toast.success('Votre article a été modifié.');
            } else {
                await this.client.create(json);
                toast.success('Votre article a été créé.');
            }
            void navigate({to: '/redactor/articles', replace: true});
            return true;
        } catch {
            toast.error('Une erreur est survenue. Veuillez réessayer.');
            return false;
        }
    }
}