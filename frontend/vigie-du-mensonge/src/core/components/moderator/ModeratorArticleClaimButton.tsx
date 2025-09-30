import type {ModeratorClient} from "@/core/dependencies/moderator/moderatorClient.ts";
import {useMutation, useQueryClient} from "@tanstack/react-query";
import {toast} from "@/core/utils/toast.ts";
import {Button} from "@/core/shadcn/components/ui/button.tsx";

export type ModeratorArticleClaimButtonProps = {
    moderatorClient: ModeratorClient;
    articleId: string;
    articleRef: string;
};

export function ModeratorArticleClaimButton({moderatorClient, articleId, articleRef}:
                                            ModeratorArticleClaimButtonProps) {
    const queryClient = useQueryClient();

    const mutation = useMutation({
        mutationFn: async () => {
            return moderatorClient.claimModeratorArticle(articleId);
        },
        onSuccess: async () => {
            toast.success("Vous avez revendiqué la modération de cet article.");
            await queryClient.invalidateQueries({queryKey: ["moderator", "articles"]});
            await queryClient.invalidateQueries({queryKey: ["moderator", "articles", "pending"]});
            await queryClient.invalidateQueries({queryKey: ["moderator", "articles", articleRef]});
        },
        onError: () => {
            toast.error("Une erreur est survenue. Veuillez réessayer.");
        },
    });

    async function onClaimArticle() {
        await mutation.mutateAsync();
    }

    return <Button onClick={onClaimArticle}>Revendiquer la modération de cet article</Button>;
}