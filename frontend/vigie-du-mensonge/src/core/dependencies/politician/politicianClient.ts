import {Politician, type PoliticianJson} from "@/core/models/politician.ts";
import type {KyInstance} from "ky";
import {api} from "@/core/dependencies/api.ts";

export class PoliticianClient {
    private readonly api: KyInstance;

    constructor(api: KyInstance) {
        this.api = api;
    }

    async getAll(): Promise<Politician[]> {
        const res = await this.api
            .get("politicians")
            .json<PoliticianJson[]>();

        return res.map((json) => Politician.fromJson(json));
    }
}

export const politicianClient = new PoliticianClient(api);