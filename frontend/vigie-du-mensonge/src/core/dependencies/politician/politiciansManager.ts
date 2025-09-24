import {PoliticianClient} from "@/core/dependencies/politician/politicianClient.ts";
import type {Politician} from "@/core/models/politician.ts";
import {Store} from "@tanstack/react-store";

class PoliticiansManager {
    private readonly client = new PoliticianClient();
    public readonly politiciansStore = new Store<Politician[]>([]);
    public readonly errStore = new Store(false);

    constructor() {
        void this.init();
    }

    async init() {
        try {
            const politicians = await this.client.getAll();
            this.politiciansStore.setState(() => politicians);
        } catch {
            this.errStore.setState(() => true);
        }
    }
}

export const politiciansManager = new PoliticiansManager();