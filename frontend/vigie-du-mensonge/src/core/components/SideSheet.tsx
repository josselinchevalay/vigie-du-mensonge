import {Button} from "@/core/shadcn/components/ui/button.tsx";
import {Sheet, SheetClose, SheetContent, SheetTitle, SheetTrigger} from "@/core/shadcn/components/ui/sheet.tsx";
import {useStore} from "@tanstack/react-store";
import {authManager} from "@/core/dependencies/auth/authManager.ts";
import {Link} from "@/core/utils/router.ts";
import {Menu} from "lucide-react";

export function SideSheet() {
    const auth = useStore(authManager.authStore);

    return <Sheet>
        <SheetTrigger asChild>
            <Button variant="ghost"><Menu/></Button>
        </SheetTrigger>
        <SheetContent side="right" className="w-64">
            <div className="flex flex-col gap-8 mt-16 mx-4">
                <SheetTitle>Navigation</SheetTitle>
                {!auth ? (
                    <>
                        <SheetClose asChild>
                            <Link
                                to="/sign-in"
                                className="inline-flex items-center justify-center rounded-md border px-3 py-2 text-sm font-medium hover:bg-accent"
                            >
                                Connexion
                            </Link>
                        </SheetClose>
                        <SheetClose asChild>
                            <Link
                                to="/sign-up" search={{token: undefined}}
                                className="inline-flex items-center justify-center rounded-md bg-primary px-3 py-2 text-sm font-medium text-primary-foreground hover:bg-primary/90"
                            >
                                Inscription
                            </Link>
                        </SheetClose>
                    </>
                ) : (
                    <>
                        {auth.isModerator &&
                            <SheetClose asChild>
                                <Link
                                    to="/moderator/articles"
                                    className="inline-flex items-center justify-center rounded-md border px-3 py-2 bg-primary-foreground text-primary text-sm font-medium hover:bg-accent"
                                >
                                    Espace modérateur
                                </Link>
                            </SheetClose>
                        }
                        {auth.isRedactor &&
                            <SheetClose asChild>
                                <Link
                                    to="/redactor/articles"
                                    className="inline-flex items-center justify-center rounded-md border px-3 py-2 bg-primary-foreground text-primary text-sm font-medium hover:bg-accent"
                                >
                                    Espace rédacteur
                                </Link>
                            </SheetClose>
                        }
                        <SheetClose asChild>
                            <Button onClick={() => authManager.signOut()}>Déconnexion</Button>
                        </SheetClose>
                    </>
                )}
            </div>
        </SheetContent>
    </Sheet>;
}