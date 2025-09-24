import {Politician, type PoliticianJson} from "@/core/models/politician.ts";

export type ArticleJson = {
    id: string;
    title: string;
    eventDate: string;
    updatedAt: string;
    politicians: PoliticianJson[];
    tags: string[];
    sources: string[];
}

export class Article {
    public id: string;
    public title: string;
    public eventDate: Date;
    public updatedAt: Date;
    public politicians: Politician[];
    public tags: string[];
    public sources: string[];

    constructor(
        id: string,
        title: string,
        eventDate: Date,
        updatedAt: Date,
        politicians: Politician[],
        tags: string[],
        sources: string[],
    ) {
        this.id = id;
        this.title = title;
        this.eventDate = eventDate;
        this.updatedAt = updatedAt;
        this.politicians = politicians;
        this.tags = tags;
        this.sources = sources;
    }

    public static fromJson(json: ArticleJson): Article {
        const politicians = json.politicians.map(Politician.fromJson);
        const tags = [
            ...politicians.map(p => p.fullName),
            ...json.tags,
        ];
        return new Article(
            json.id,
            json.title,
            new Date(json.eventDate),
            new Date(json.updatedAt),
            politicians,
            tags,
            json.sources,
        );
    }
}