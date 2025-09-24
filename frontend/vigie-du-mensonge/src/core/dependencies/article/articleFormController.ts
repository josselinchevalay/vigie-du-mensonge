import {articleClient, type ArticleCreateJson} from "@/core/dependencies/article/articleClient.ts";
import {toast} from "@/core/utils/toast.ts";

export class ArticleFormController {
    async onSubmit(json: ArticleCreateJson): Promise<boolean> {
        try {
            await articleClient.create(json);
            toast.success('Votre article a été créé.');
            return true;
        } catch {
            toast.error('Une erreur est survenue. Veuillez réessayer.');
            return false;
        }
    }
}