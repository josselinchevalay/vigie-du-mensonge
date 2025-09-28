import type {ModeratorClient} from "@/core/dependencies/moderator/moderatorClient.ts";
import type {Article} from "@/core/models/article.ts";
import {useState} from "react";
import {Button} from "@/core/shadcn/components/ui/button.tsx";
import {ArrowLeft, ArrowRight} from "lucide-react";
import {ArticleDisplay} from "@/core/components/article/ArticleDisplay.tsx";
import {ArticleStatuses} from "@/core/models/articleStatus.ts";
import {ModeratorArticleClaimButton} from "@/core/components/moderator/ModeratorArticleClaimButton.tsx";
import {authManager} from "@/core/dependencies/auth/authManager.ts";
import {ModeratorArticleReviewForm} from "@/core/components/moderator/ModeratorArticleReviewForm.tsx";

export type ModeratorArticlesByReferenceProps = {
    moderatorClient: ModeratorClient;
    articles: Article[];
}

export function ModeratorArticlesByReference({moderatorClient, articles}: ModeratorArticlesByReferenceProps) {
    const [index, setIndex] = useState<number>(0);
    const selected = articles[index];

    return <div className="flex flex-col items-center gap-4 min-w-0 py-2">

        {
            (selected.status === ArticleStatuses.UNDER_REVIEW && !selected.moderatorTag) &&

            <ModeratorArticleClaimButton moderatorClient={moderatorClient}
                                         articleId={selected.id}
                                         articleRef={selected.reference}/>
        }

        {
            (selected.status === ArticleStatuses.UNDER_REVIEW && selected.moderatorTag === authManager.authStore.state?.tag) &&

            <ModeratorArticleReviewForm moderatorClient={moderatorClient}
                                        articleId={selected.id}
                                        articleRef={selected.reference}/>
        }

        <h1 className="text-xl font-bold">{selected.versionLabel}</h1>

        <div className="flex flex-row justify-center gap-8">
            <Button
                disabled={index === articles.length - 1}
                onClick={() => setIndex(index + 1)}
            >
                <ArrowLeft></ArrowLeft>
            </Button>
            <Button
                disabled={index === 0}
                onClick={() => setIndex(index - 1)}
            >
                <ArrowRight></ArrowRight>
            </Button>
        </div>

        <ArticleDisplay article={articles[index]}/>

    </div>;
}