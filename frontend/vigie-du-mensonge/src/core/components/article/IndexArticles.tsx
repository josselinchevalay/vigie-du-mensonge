import {useStore} from "@tanstack/react-store";
import {authManager} from "@/core/dependencies/auth/authManager.ts";
import {Button} from "@/core/shadcn/components/ui/button.tsx";

export function IndexArticles() {
    const auth = useStore(authManager.authStore);

    return (
        <div className="mx-auto w-full max-w-sm">
            <div className="flex flex-col items-center justify-center gap-4">
                {auth?.isRedactor &&
                    <Button>
                        Ajouter un article
                    </Button>
                }
                <h3>CLEMENT J'ATTENDS TOUJOURS LES CRITÈRES D'ACCEPTATION</h3>
            </div>
        </div>
    );
}