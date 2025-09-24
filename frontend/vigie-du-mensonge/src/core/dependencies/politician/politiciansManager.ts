import type {Politician} from "@/core/models/politician.ts";
import {Store} from "@tanstack/react-store";
import {PoliticianClient, politicianClient} from "@/core/dependencies/politician/politicianClient.ts";

class PoliticiansManager {
    private readonly client: PoliticianClient;

    public readonly politiciansStore = new Store<Politician[]>([]);
    public readonly errStore = new Store(false);

    constructor(client: PoliticianClient) {
        this.client = client;
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

export const politiciansManager = new PoliticiansManager(politicianClient);