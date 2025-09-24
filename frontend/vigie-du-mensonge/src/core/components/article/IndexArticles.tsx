import {useStore} from "@tanstack/react-store";
import {authManager} from "@/core/dependencies/auth/authManager.ts";
import {Link} from "@/core/utils/router.ts";

export function IndexArticles() {
    const auth = useStore(authManager.authStore);

    return (
        <div className="mx-auto w-full max-w-sm">
            <div className="flex flex-col items-center justify-center gap-4">
                {auth?.isRedactor &&
                    <Link to="/article-form">
                        Ajouter un article
                    </Link>
                }
                <h3>CLEMENT J'ATTENDS TOUJOURS LES CRITÈRES D'ACCEPTATION</h3>
            </div>
        </div>
    );
}