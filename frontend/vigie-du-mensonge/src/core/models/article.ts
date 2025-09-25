import {Politician, type PoliticianJson} from "@/core/models/politician.ts";
import {type ArticleCategory} from "@/core/models/articleCategory.ts";
import {formatDate} from "@/core/utils/formatDate.ts";

export type ArticleJson = {
    id: string;
    title: string;
    eventDate: string;
    updatedAt: string;
    politicians: PoliticianJson[];
    tags: string[];
    category: ArticleCategory;
}

export type ArticleDetails = {
    body: string;
    sources: string[];
}

export class Article {
    public id: string;
    public title: string;
    public eventDate: Date;
    public updatedAt: Date;
    public politicians: Politician[];
    public tags: string[];
    public category: ArticleCategory;
    public details: ArticleDetails | null = null;

    constructor(
        id: string,
        title: string,
        eventDate: Date,
        updatedAt: Date,
        politicians: Politician[],
        tags: string[],
        category: ArticleCategory,
    ) {
        this.id = id;
        this.title = title;
        this.eventDate = eventDate;
        this.updatedAt = updatedAt;
        this.politicians = politicians;
        this.tags = tags;
        this.category = category;
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
            new Date(json.eventDate),
            new Date(json.updatedAt),
            json.politicians.map(Politician.fromJson),
            json.tags,
            json.category,
        );
    }
}