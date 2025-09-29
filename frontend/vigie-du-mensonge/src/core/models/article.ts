import {Politician, type PoliticianJson} from "@/core/models/politician.ts";
import {type ArticleCategory} from "@/core/models/articleCategory.ts";
import {type ArticleStatus, ArticleStatuses} from "@/core/models/articleStatus.ts";
import {ArticleReview, type ArticleReviewJson} from "@/core/models/articleReview.ts";

export type ArticleJson = {
    id: string;
    reference: string;

    redactorTag?: string;
    moderatorTag?: string;

    title: string;
    body?: string;

    category: ArticleCategory;
    status: ArticleStatus;
    review?: ArticleReviewJson;

    eventDate: string;
    updatedAt: string;

    minor?: number;
    major?: number;

    sources?: string[];
    tags?: string[];
    politicians?: PoliticianJson[];
}

export class Article {
    public id: string;
    public reference: string;

    redactorTag?: string;
    moderatorTag?: string;

    public title: string;
    public body?: string;

    public category: ArticleCategory;
    public status?: ArticleStatus;
    public review?: ArticleReview;

    public eventDate: Date;
    public updatedAt: Date;

    public minor?: number;
    public major?: number;

    public tags?: string[];
    public sources?: string[];
    public politicians?: Politician[];

    constructor(
        id: string,
        reference: string,
        redactorTag: string | undefined,
        moderatorTag: string | undefined,
        title: string,
        body: string | undefined,
        category: ArticleCategory,
        status: ArticleStatus | undefined,
        review: ArticleReview | undefined,
        eventDate: Date,
        updatedAt: Date,
        minor: number | undefined,
        major: number | undefined,
        tags: string[] | undefined,
        sources: string[] | undefined,
        politicians: Politician[] | undefined,
    ) {
        this.id = id;
        this.reference = reference;
        this.redactorTag = redactorTag;
        this.moderatorTag = moderatorTag;
        this.title = title;
        this.body = body;
        this.category = category;
        this.status = status;
        this.review = review;
        this.eventDate = eventDate;
        this.updatedAt = updatedAt;
        this.minor = minor;
        this.major = major;
        this.tags = tags;
        this.sources = sources;
        this.politicians = politicians;
    }

    public get isPublished(): boolean {
        return this.status === ArticleStatuses.PUBLISHED;
    }

    public get isDraft(): boolean {
        return this.status === ArticleStatuses.DRAFT;
    }

    public get isUnderReview(): boolean {
        return this.status === ArticleStatuses.UNDER_REVIEW;
    }

    public get isChangeRequested(): boolean {
        return this.status === ArticleStatuses.CHANGE_REQUESTED;
    }

    public get isArchived(): boolean {
        return this.status === ArticleStatuses.ARCHIVED;
    }

    public get politicianIds(): string[] {
        if (!this.politicians) return [];
        return this.politicians.map(p => p.id);
    }

    public get versionLabel(): string {
        const major = this.major ?? 0;
        const minor = this.minor ?? 0;
        return `v${major}.${minor}`;
    }

    public static fromJson(json: ArticleJson): Article {
        return new Article(
            json.id,
            json.reference,
            json.redactorTag,
            json.moderatorTag,
            json.title,
            json.body,
            json.category,
            json.status,
            json.review ? ArticleReview.fromJson(json.review) : undefined,
            new Date(json.eventDate),
            new Date(json.updatedAt),
            json.minor,
            json.major,
            json.tags,
            json.sources,
            json.politicians?.map(Politician.fromJson),
        );
    }
}