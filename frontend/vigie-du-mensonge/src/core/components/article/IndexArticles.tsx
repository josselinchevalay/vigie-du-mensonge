import {useStore} from "@tanstack/react-store";
import {authManager} from "@/core/dependencies/auth/authManager.ts";
import {Link} from "@/core/utils/router.ts";

export function IndexArticles() {
    const auth = useStore(authManager.authStore);

    return (
        <>
            <div className="mx-auto w-full max-w-sm">
                <div className="flex flex-col items-center justify-center gap-4">
                    {auth?.isRedactor &&
                        <Link to="/redactor/article-form">
                            Ajouter un article
                        </Link>
                    }
                    <h3>CLEMENT J'ATTENDS TOUJOURS LES CRITÈRES D'ACCEPTATION</h3>

                </div>

                <br/>


            </div>

            <div className="w-full flex flex-col sm:flex-row gap-4 items-center justify-between">
                <div className="flex flex-col">
                    <img src="/adso.jpg" alt="adso" className="w-100 h-100"/>
                    <h3>Le plus beau</h3>
                </div>

                <h3>{"ADSO >>> HIPPIAS"}</h3>

                <div className="flex flex-col">
                    <img src="/hippias.PNG" alt="hippias" className="w-100 h-100"/>
                    <h3>Il pu le seum</h3>
                </div>
            </div>
        </>
    );
}