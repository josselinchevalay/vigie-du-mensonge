import {createFileRoute} from '@tanstack/react-router';
import {useStore} from "@tanstack/react-store";
import {politiciansManager} from "@/core/dependencies/politician/politiciansManager.ts";

export const Route = createFileRoute('/')({
    component: Index,
});

function Index() {
    const politicians = useStore(politiciansManager.politiciansStore);

    return (
        <div className="p-2">
            <h3>CLEMENT J'ATTENDS TOUJOURS LES CRITÈRES D'ACCEPTATION</h3>
            <br/>
            <>{politicians.length} politicians loaded</>
        </div>
    );
}