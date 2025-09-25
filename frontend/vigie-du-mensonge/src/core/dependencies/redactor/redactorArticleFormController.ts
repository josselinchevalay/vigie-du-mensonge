import {toast} from "@/core/utils/toast.ts";
import type {RedactorArticleJson, RedactorClient} from "@/core/dependencies/redactor/redactorClient.ts";
import {navigate} from "@/core/utils/router.ts";

export class RedactorArticleFormController {
    private readonly client: RedactorClient;

    constructor(client: RedactorClient) {
        this.client = client;
    }

    async onSubmit(json: RedactorArticleJson): Promise<boolean> {
        try {
            await this.client.createDraftArticle(json);
            toast.success('Votre article a été créé.');
            void navigate({to: '/redactor', replace: true});
            return true;
        } catch {
            toast.error('Une erreur est survenue. Veuillez réessayer.');
            return false;
        }
    }
}