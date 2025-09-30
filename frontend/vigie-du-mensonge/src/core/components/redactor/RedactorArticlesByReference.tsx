import type {Article} from "@/core/models/article.ts";
import {useState} from "react";
import {Button} from "@/core/shadcn/components/ui/button.tsx";
import {ArrowLeft, ArrowRight, Eye, SquarePen} from "lucide-react";
import {ArticleDisplay} from "@/core/components/article/ArticleDisplay.tsx";
import {ArticleStatuses, ArticleStatusLabels} from "@/core/models/articleStatus.ts";
import type {RedactorClient} from "@/core/dependencies/redactor/redactorClient.ts";
import {
    Dialog,
    DialogClose,
    DialogContent,
    DialogFooter,
    DialogHeader,
    DialogTitle,
    DialogTrigger
} from "@/core/shadcn/components/ui/dialog.tsx";
import {ArticleReviewCard} from "@/core/components/article/ArticleReviewCard.tsx";
import {RedactorArticleFormLoader} from "@/core/components/redactor/RedactorArticleFormLoader.tsx";
import type {PoliticianClient} from "@/core/dependencies/politician/politicianClient.ts";

export type RedactorArticlesByReferenceProps = {
    redactorClient: RedactorClient;
    politicianClient: PoliticianClient;
    articles: Article[];
}

export function RedactorArticlesByReference({redactorClient, politicianClient, articles}:
                                            RedactorArticlesByReferenceProps) {
    const [index, setIndex] = useState<number>(0);
    const [editMode, setEditMode] = useState<boolean>(false);

    const selected = articles[index];

    const editable = selected.status === ArticleStatuses.DRAFT || selected.status === ArticleStatuses.CHANGE_REQUESTED;

    return <div className="flex flex-col items-center gap-4 min-w-0 py-2">

        <div className="flex flex-row justify-center gap-8">
            <Button variant="ghost"
                    disabled={index === articles.length - 1}
                    onClick={() => setIndex(index + 1)}>
                <ArrowLeft></ArrowLeft>
            </Button>

            <h1 className="text-xl font-bold">{selected.versionLabel}</h1>

            <Button variant="ghost"
                    disabled={index === 0}
                    onClick={() => setIndex(index - 1)}
            >
                <ArrowRight></ArrowRight>
            </Button>
        </div>

        {!editable &&
            <div className="mb-4 p-4 rounded-md border bg-primary text-primary-foreground">
                {`Les articles dont le statut est [ ${ArticleStatusLabels[selected!.status!]} ] ne peuvent pas être modifiés.`}
            </div>
        }

        {selected.review && <div className="px-2">
            <ArticleReviewCard review={selected.review}/>
        </div>}

        {editable &&
            (
                !editMode
                    ? <Button variant="ghost"
                              onClick={() => setEditMode(true)}><SquarePen/></Button>
                    : <Dialog>
                        <DialogTrigger asChild>
                            <Button variant="ghost"><Eye/></Button>
                        </DialogTrigger>
                        <DialogContent aria-describedby={undefined}>
                            <DialogHeader>
                                <DialogTitle>Toute modification non enregistrée sera perdue.</DialogTitle>
                            </DialogHeader>
                            <DialogFooter>
                                <DialogClose asChild>
                                    <Button onClick={() => setEditMode(false)}>Passer en mode lecteur</Button>
                                </DialogClose>
                                <DialogClose asChild>
                                    <Button>Rester en mode édition</Button>
                                </DialogClose>
                            </DialogFooter>
                        </DialogContent>
                    </Dialog>
            )
        }

        {editMode ? <RedactorArticleFormLoader article={selected} redactorClient={redactorClient}
                                               politicianClient={politicianClient}
                                               onSubmitSuccess={() => setEditMode(false)}/> :
            <ArticleDisplay article={articles[index]}/>}
    </div>;
}