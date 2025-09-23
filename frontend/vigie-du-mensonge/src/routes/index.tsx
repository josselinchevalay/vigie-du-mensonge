import {createFileRoute} from '@tanstack/react-router';

export const Route = createFileRoute('/')({
    component: Index,
});

function Index() {
    return (
        <div className="p-2">
            <h3>CLEMENT J'ATTENDS TOUJOURS LES CRITÈRES D'ACCEPTATION</h3>
        </div>
    );
}