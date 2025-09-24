import {Politician, type PoliticianJson} from "@/core/models/politician.ts";
import {type ArticleCategory} from "@/core/models/articleCategory.ts";

export type ArticleJson = {
    id: string;
    title: string;
    eventDate: string;
    updatedAt: string;
    politicians: PoliticianJson[];
    tags: string[];
    sources: string[];
    category: ArticleCategory;
}

export class Article {
    public id: string;
    public title: string;
    public eventDate: Date;
    public updatedAt: Date;
    public politicians: Politician[];
    public tags: string[];
    public sources: string[];
    public category: ArticleCategory;

    constructor(
        id: string,
        title: string,
        eventDate: Date,
        updatedAt: Date,
        politicians: Politician[],
        tags: string[],
        sources: string[],
        category: ArticleCategory,
    ) {
        this.id = id;
        this.title = title;
        this.eventDate = eventDate;
        this.updatedAt = updatedAt;
        this.politicians = politicians;
        this.tags = tags;
        this.sources = sources;
        this.category = category;
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
            json.category,
        );
    }
}