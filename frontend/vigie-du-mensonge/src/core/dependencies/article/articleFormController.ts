import {ArticleClient, type ArticleCreateJson} from "@/core/dependencies/article/articleClient.ts";
import {toast} from "@/core/utils/toast.ts";

export class ArticleFormController {
    private readonly client: ArticleClient;

    constructor(client: ArticleClient) {
        this.client = client;
    }

    async onSubmit(json: ArticleCreateJson): Promise<boolean> {
        try {
            await this.client.create(json);
            toast.success('Votre article a été créé.');
            return true;
        } catch {
            toast.error('Une erreur est survenue. Veuillez réessayer.');
            return false;
        }
    }
}