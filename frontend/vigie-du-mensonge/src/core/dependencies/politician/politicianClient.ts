import {Politician, type PoliticianJson} from "@/core/models/politician.ts";
import type {KyInstance} from "ky";
import {api} from "@/core/dependencies/api.ts";

export class PoliticianClient {
    private readonly api: KyInstance;

    constructor(api: KyInstance) {
        this.api = api;
    }

    private async _getAll(): Promise<Politician[]> {
        const res = await this.api
            .get("politicians")
            .json<PoliticianJson[]>();

        return res.map((json) => Politician.fromJson(json));
    }

    getAll = (): { queryKey: string[], queryFn: () => Promise<Politician[]> } => {
        return {
            queryKey: ["politicians"],
            queryFn: () => this._getAll(),
        };
    };
}

export const politicianClient = new PoliticianClient(api);