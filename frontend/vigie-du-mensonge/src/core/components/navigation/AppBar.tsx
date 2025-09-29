import {authManager} from "@/core/dependencies/auth/authManager.ts";
import {useStore} from "@tanstack/react-store";
import {Button} from "../../shadcn/components/ui/button.tsx";
import {Link} from "@/core/utils/router.ts";
import {SideSheet} from "@/core/components/navigation/SideSheet.tsx";
import {ThemeToggle} from "@/core/components/misc/ThemeToggle.tsx";

export default function AppBar() {
    const auth = useStore(authManager.authStore);

    return <nav className="sticky top-0 z-40 w-full border-b bg-card shadow-md">

        <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">

            <div className="flex flex-row min-h-14 items-center justify-between">

                <div className="md:flex flex-row items-center gap-2 min-w-0 p-4">
                    <Link to="/" className="font-semibold text-lg">
                        Vigie du mensonge
                    </Link>

                    {auth?.isModerator &&
                        <Link
                            to="/moderator/articles"
                            className="hidden md:flex p-2 text-sm font-medium rounded-md hover:bg-accent"
                        >
                            Espace modérateur
                        </Link>
                    }
                    {auth?.isRedactor &&
                        <Link
                            to="/redactor/articles"
                            className="hidden md:flex p-2 text-sm font-medium rounded-md hover:bg-accent"
                        >
                            Espace rédacteur
                        </Link>
                    }
                </div>

                <div className="md:hidden">
                    <SideSheet/>
                </div>

                <div className="hidden md:flex flex-row items-center gap-2 min-w-0 p-2">

                    <ThemeToggle/>

                    {!auth ?
                        <>
                            <Link
                                to="/sign-in"
                                className="p-2 text-sm font-medium rounded-md hover:bg-accent"
                            >
                                Connexion
                            </Link>
                            <Link
                                to="/sign-up" search={{token: undefined}}
                                className="p-2 text-sm font-medium rounded-md hover:bg-accent"
                            >
                                Inscription
                            </Link>
                        </>
                        : <Button
                            onClick={authManager.signOut}>Déconnexion</Button>
                    }

                </div>

            </div>

        </div>

    </nav>;
}
