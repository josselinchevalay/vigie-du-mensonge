import {Politician, type PoliticianJson} from "@/core/models/politician.ts";
import {type ArticleCategory} from "@/core/models/articleCategory.ts";
import {formatDate} from "@/core/utils/formatDate.ts";
import type {ArticleStatus} from "@/core/models/articleStatus.ts";

export type ArticleJson = {
    id: string;
    title: string;
    body: string | undefined;
    sources: string[] | undefined;
    eventDate: string;
    updatedAt: string;
    politicians: PoliticianJson[];
    tags: string[];
    category: ArticleCategory;
    status: ArticleStatus;
}

export class Article {
    public id: string;
    public title: string;
    public body: string | undefined;
    public sources: string[] | undefined;
    public eventDate: Date;
    public updatedAt: Date;
    public politicians: Politician[];
    public tags: string[];
    public category: ArticleCategory;
    public status: ArticleStatus;

    constructor(
        id: string,
        title: string,
        body: string | undefined,
        sources: string[] | undefined,
        eventDate: Date,
        updatedAt: Date,
        politicians: Politician[],
        tags: string[],
        category: ArticleCategory,
        status: ArticleStatus,
    ) {
        this.id = id;
        this.title = title;
        this.body = body;
        this.sources = sources;
        this.eventDate = eventDate;
        this.updatedAt = updatedAt;
        this.politicians = politicians;
        this.tags = tags;
        this.category = category;
        this.status = status;
    }

    public get politicianIds(): string[] {
        return this.politicians.map(p => p.id);
    }

    public get formattedEventDate(): string {
        return formatDate(this.eventDate);
    }

    public static fromJson(json: ArticleJson): Article {
        return new Article(
            json.id,
            json.title,
            json.body,
            json.sources,
            new Date(json.eventDate),
            new Date(json.updatedAt),
            json.politicians.map(Politician.fromJson),
            json.tags,
            json.category,
            json.status,
        );
    }
}