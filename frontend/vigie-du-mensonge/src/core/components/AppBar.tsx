import {authManager} from "@/core/dependencies/auth/authManager.ts";
import {useStore} from "@tanstack/react-store";
import {Button} from "../shadcn/components/ui/button";
import {Link} from "@/core/utils/router.ts";
import {SideSheet} from "@/core/components/SideSheet.tsx";


export default function AppBar() {
    const auth = useStore(authManager.authStore);

    return (
        <nav
            className="sticky top-0 z-40 w-full border-b bg-background">
            <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
                <div className="flex flex-row min-h-14 items-center justify-between">
                    <Link to="/" className="py-2 text-base font-semibold text-foreground">
                        Vigie du mensonge
                    </Link>

                    <div className="md:hidden">
                        <SideSheet/>
                    </div>

                    <div className="hidden md:flex items-center gap-2 min-w-0 py-2">
                        {!auth ? (
                            <>
                                <Link
                                    to="/sign-in"
                                    className="inline-flex items-center justify-center rounded-md border px-3 py-2 text-sm font-medium hover:bg-accent"
                                >
                                    Connexion
                                </Link>
                                <Link
                                    to="/sign-up" search={{token: undefined}}
                                    className="inline-flex items-center justify-center rounded-md bg-primary px-3 py-2 text-sm font-medium text-primary-foreground hover:bg-primary/90"
                                >
                                    Inscription
                                </Link>
                            </>
                        ) : (
                            <>
                                {auth.isRedactor &&
                                    <>
                                        <Link
                                            to="/moderator/articles"
                                            className="inline-flex items-center justify-center rounded-md border px-3 py-2 text-sm font-medium hover:bg-accent"
                                        >
                                            Espace modérateur
                                        </Link>
                                        <Link
                                            to="/redactor/articles"
                                            className="inline-flex items-center justify-center rounded-md border px-3 py-2 text-sm font-medium hover:bg-accent"
                                        >
                                            Espace rédacteur
                                        </Link>
                                    </>
                                }
                                <Button onClick={() => authManager.signOut()}>Déconnexion</Button>
                            </>
                        )}
                    </div>
                </div>
            </div>
        </nav>
    );
}
