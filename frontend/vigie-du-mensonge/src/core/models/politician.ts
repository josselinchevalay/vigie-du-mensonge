export type PoliticianJson = {
    id: string;
    fullName: string;
};

export class Politician {
    public id: string;
    public fullName: string;

    constructor(id: string, fullName: string) {
        this.id = id;
        this.fullName = fullName;
    }

    public static fromJson(json: PoliticianJson): Politician {
        return new Politician(json.id, json.fullName);
    }
}