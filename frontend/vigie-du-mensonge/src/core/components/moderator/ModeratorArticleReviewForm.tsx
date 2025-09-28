import type {ModeratorClient} from "@/core/dependencies/moderator/moderatorClient.ts";
import {useMutation, useQueryClient} from "@tanstack/react-query";
import {z} from "zod";
import {ArticleStatuses, ArticleStatusLabels} from "@/core/models/articleStatus.ts";
import {useForm} from "react-hook-form";
import {zodResolver} from "@hookform/resolvers/zod";
import {toast} from "@/core/utils/toast.ts";
import {Form, FormControl, FormField, FormItem, FormLabel, FormMessage} from "@/core/shadcn/components/ui/form";
import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue} from "@/core/shadcn/components/ui/select";
import {Button} from "@/core/shadcn/components/ui/button";

export type ModeratorArticleReviewFormProps = {
    moderatorClient: ModeratorClient;
    articleId: string;
    articleRef: string;
};

const AllowedDecisions = [
    ArticleStatuses.ARCHIVED,
    ArticleStatuses.CHANGE_REQUESTED,
    ArticleStatuses.PUBLISHED,
] as const;

const formSchema = z
    .object({
        decision: z.enum(AllowedDecisions),
        notes: z
            .string()
            .min(30, "Au moins 30 caractères")
            .max(200, "Maximum 200 caractères")
            .optional(),
    })
    .refine(
        (data) => data.decision === ArticleStatuses.PUBLISHED || data.notes !== undefined,
        {
            message: "Vous devez fournir des notes pour expliquer votre décision.",
            path: ["notes"],
        }
    );

export type ModeratorArticleReviewFormInput = z.infer<typeof formSchema>;

export function ModeratorArticleReviewForm({moderatorClient, articleId, articleRef}: ModeratorArticleReviewFormProps) {
    const form = useForm<ModeratorArticleReviewFormInput>({
        resolver: zodResolver(formSchema),
        mode: "onSubmit",
        defaultValues: {
            decision: ArticleStatuses.CHANGE_REQUESTED,
            notes: undefined,
        },
    });

    const queryClient = useQueryClient();

    const mutation = useMutation({
        mutationFn: async (input: ModeratorArticleReviewFormInput) => {
            return moderatorClient.reviewModeratorArticle(articleId, {
                decision: input.decision,
                notes: input.notes,
            });
        },
        onSuccess: () => {
            toast.success("Votre review a été enregistrée.");
            void queryClient.invalidateQueries({queryKey: ["moderator", "articles", articleRef]});
            void queryClient.invalidateQueries({queryKey: ["moderator", "articles"]});
        },
        onError: () => {
            toast.error("Une erreur est survenue. Veuillez réessayer.");
        },
    });

    async function onSubmit(input: ModeratorArticleReviewFormInput) {
        mutation.mutate(input);
    }

    return (
        <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)} className="flex flex-col gap-4">
                <FormField
                    control={form.control}
                    name="decision"
                    render={({field}) => (
                        <FormItem>
                            <FormLabel>Décision</FormLabel>
                            <FormControl>
                                <Select onValueChange={field.onChange} value={field.value}
                                        disabled={mutation.isPending}>
                                    <SelectTrigger className="min-w-52">
                                        <SelectValue placeholder="Votre décision"/>
                                    </SelectTrigger>
                                    <SelectContent>
                                        {AllowedDecisions.map((value) => (
                                            <SelectItem key={value} value={value}>
                                                {ArticleStatusLabels[value as keyof typeof ArticleStatusLabels]}
                                            </SelectItem>
                                        ))}
                                    </SelectContent>
                                </Select>
                            </FormControl>
                            <FormMessage/>
                        </FormItem>
                    )}
                />

                <FormField
                    control={form.control}
                    name="notes"
                    render={({field}) => (
                        <FormItem>
                            <FormLabel>
                                Notes {form.getValues("decision") !== ArticleStatuses.PUBLISHED &&
                                <span className="text-muted-foreground">(obligatoire)</span>}
                            </FormLabel>
                            <FormControl>
                <textarea
                    {...field}
                    rows={4}
                    className="border-input focus-visible:border-ring focus-visible:ring-ring/50 dark:bg-input/30 dark:hover:bg-input/50 w-full rounded-md border bg-transparent px-3 py-2 text-sm shadow-xs outline-none focus-visible:ring-[3px] disabled:cursor-not-allowed disabled:opacity-50"
                    placeholder="Expliquez brièvement votre décision (30 à 200 caractères, optionnel si vous choisissez [Publié])"
                    minLength={30}
                    maxLength={200}
                    disabled={mutation.isPending}
                />
                            </FormControl>
                            <FormMessage/>
                        </FormItem>
                    )}
                />

                <div className="flex items-center gap-2">
                    <Button type="submit" disabled={mutation.isPending}>
                        {mutation.isPending ? "Envoi..." : "Enregistrer la review"}
                    </Button>
                    <Button
                        type="button"
                        variant="ghost"
                        onClick={() => form.reset()}
                        disabled={mutation.isPending}
                    >
                        Réinitialiser
                    </Button>
                </div>
            </form>
        </Form>
    );
}