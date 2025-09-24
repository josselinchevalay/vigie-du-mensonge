import {createFileRoute} from '@tanstack/react-router';
import {useStore} from "@tanstack/react-store";
import {articlesManager} from "@/core/dependencies/article/articlesManager.ts";

export const Route = createFileRoute('/')({
    component: Index,
});

function Index() {
    const articles = useStore(articlesManager.articlesStore);

    return (
        <div className="p-2">
            <h3>CLEMENT J'ATTENDS TOUJOURS LES CRITÈRES D'ACCEPTATION</h3>
            <br/>
            <>{articles.length} articles loaded</>
        </div>
    );
}