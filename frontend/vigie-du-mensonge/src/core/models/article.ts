import {Politician, type PoliticianJson} from "@/core/models/politician.ts";
import {type ArticleCategory} from "@/core/models/articleCategory.ts";
import {formatDate} from "@/core/utils/formatDate.ts";
import {type ArticleStatus, ArticleStatuses} from "@/core/models/articleStatus.ts";

export type ArticleJson = {
    id: string;
    reference: string;

    title: string;
    body?: string;

    category: ArticleCategory;
    status: ArticleStatus;

    eventDate: string;
    updatedAt: string;

    minor?: number;
    major?: number;

    sources?: string[];
    tags?: string[];
    politicians?: PoliticianJson[];
    otherVersions?: ArticleJson[];
}

export class Article {
    public id: string;
    public reference: string;

    public title: string;
    public body?: string;

    public category: ArticleCategory;
    public status?: ArticleStatus;

    public eventDate: Date;
    public updatedAt: Date;

    public minor?: number;
    public major?: number;

    public tags?: string[];
    public sources?: string[];
    public politicians?: Politician[];
    public otherVersions?: Article[];

    constructor(
        id: string,
        reference: string,
        title: string,
        body: string | undefined,
        category: ArticleCategory,
        status: ArticleStatus | undefined,
        eventDate: Date,
        updatedAt: Date,
        minor: number | undefined,
        major: number | undefined,
        tags: string[] | undefined,
        sources: string[] | undefined,
        politicians: Politician[] | undefined,
        otherVersions: Article[] | undefined,
    ) {
        this.id = id;
        this.reference = reference;
        this.title = title;
        this.body = body;
        this.category = category;
        this.status = status;
        this.eventDate = eventDate;
        this.updatedAt = updatedAt;
        this.minor = minor;
        this.major = major;
        this.tags = tags;
        this.sources = sources;
        this.politicians = politicians;
        this.otherVersions = otherVersions;
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

    public get formattedEventDate(): string {
        return formatDate(this.eventDate);
    }

    public static fromJson(json: ArticleJson): Article {

        return new Article(
            json.id,
            json.reference,
            json.title,
            json.body,
            json.category,
            json.status,
            new Date(json.eventDate),
            new Date(json.updatedAt),
            json.minor,
            json.major,
            json.tags,
            json.sources,
            json.politicians?.map(Politician.fromJson),
            json.otherVersions?.map(Article.fromJson),
        );
    }
}