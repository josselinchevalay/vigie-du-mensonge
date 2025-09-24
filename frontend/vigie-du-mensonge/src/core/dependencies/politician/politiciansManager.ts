import type {Politician} from "@/core/models/politician.ts";
import {Store} from "@tanstack/react-store";
import {politicianClient} from "@/core/dependencies/politician/politicianClient.ts";

class PoliticiansManager {
    public readonly politiciansStore = new Store<Politician[]>([]);
    public readonly errStore = new Store(false);

    constructor() {
        void this.init();
    }

    async init() {
        try {
            const politicians = await politicianClient.getAll();
            this.politiciansStore.setState(() => politicians);
        } catch {
            this.errStore.setState(() => true);
        }
    }
}

export const politiciansManager = new PoliticiansManager();