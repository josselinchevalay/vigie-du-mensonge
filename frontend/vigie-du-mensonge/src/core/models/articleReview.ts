import type {ArticleStatus} from "@/core/models/articleStatus.ts";

export type ArticleReviewJson = {
    moderatorTag?: string;
    notes?: string;
    decision?: ArticleStatus;
}

export class ArticleReview {
    public moderatorTag?: string;
    public notes?: string;
    public decision?: ArticleStatus;

    constructor(moderatorTag?: string, notes?: string, decision?: ArticleStatus) {
        this.moderatorTag = moderatorTag;
        this.notes = notes;
        this.decision = decision;
    }

    static fromJson(json: ArticleReviewJson): ArticleReview {
        return new ArticleReview(json.moderatorTag, json.notes, json.decision);
    }
}