import {Politician, type PoliticianJson} from "@/core/models/politician.ts";
import {api} from "@/core/dependencies/api.ts";

export class PoliticianClient {
    async getAll(): Promise<Politician[]> {
        const res = await api
            .get("politicians")
            .json<PoliticianJson[]>();

        return res.map((json) => Politician.fromJson(json));
    }
}