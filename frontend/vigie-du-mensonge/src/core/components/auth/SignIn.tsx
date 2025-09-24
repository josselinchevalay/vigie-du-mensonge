import {z} from "zod";
import type {SignInController} from "@/core/dependencies/auth/signInController.ts";
import {useForm} from "react-hook-form";
import {zodResolver} from "@hookform/resolvers/zod";
import * as React from "react";
import {Form, FormControl, FormField, FormItem, FormLabel, FormMessage} from "@/core/shadcn/components/ui/form.tsx";
import {Input} from "@/core/shadcn/components/ui/input.tsx";
import {Eye, EyeOff} from "lucide-react";
import {Button} from "@/core/shadcn/components/ui/button.tsx";
import {Link} from "@/core/utils/router.ts";

export type SignInProps = {
    controller: SignInController;
};

const formSchema = z.object({
    email: z.email("Adresse e-mail invalide"),
    password: z.string().min(1, "Mot de passe requis"),
});

export type SignInInput = z.infer<typeof formSchema>;

export function SignIn({controller}: SignInProps) {
    const form = useForm<SignInInput>({
        resolver: zodResolver(formSchema),
        defaultValues: {email: '', password: ""},
        mode: "onSubmit",
    });

    const [showPassword, setShowPassword] = React.useState(false);

    return (
        <div className="mx-auto w-full max-w-sm">
            <Form {...form}>
                <form onSubmit={form.handleSubmit(({email, password}) => controller.onSignIn(email, password))}
                      className="space-y-4">
                    <div className="space-y-1">
                        <h1 className="text-xl font-semibold">Connexion</h1>
                        <p className="text-sm text-muted-foreground">Connectez-vous à Vigie du mensonge</p>
                    </div>

                    <FormField
                        control={form.control}
                        name="email"
                        render={({field}) => (
                            <FormItem>
                                <FormLabel>Email</FormLabel>
                                <FormControl>
                                    <Input type="email" placeholder="vous@exemple.com"
                                           autoComplete="email" {...field} />
                                </FormControl>
                                <FormMessage/>
                            </FormItem>
                        )}
                    />

                    <FormField
                        control={form.control}
                        name="password"
                        render={({field}) => (
                            <FormItem>
                                <FormLabel>Mot de passe</FormLabel>
                                <FormControl>
                                    <div className="relative">
                                        <Input
                                            type={showPassword ? "text" : "password"}
                                            placeholder="••••••••"
                                            autoComplete="current-password"
                                            className="pr-10"
                                            {...field}
                                        />
                                        <button
                                            type="button"
                                            onClick={() => setShowPassword(v => !v)}
                                            className="absolute inset-y-0 right-0 flex items-center pr-3 text-muted-foreground hover:text-foreground"
                                            aria-label={showPassword ? "Masquer le mot de passe" : "Afficher le mot de passe"}
                                        >
                                            {showPassword ? (
                                                <EyeOff className="h-4 w-4" aria-hidden="true"/>
                                            ) : (
                                                <Eye className="h-4 w-4" aria-hidden="true"/>
                                            )}
                                        </button>
                                    </div>
                                </FormControl>
                                <FormMessage/>
                            </FormItem>
                        )}
                    />

                    <Button type="submit" disabled={form.formState.isSubmitting} className="w-full">
                        {form.formState.isSubmitting ? "Connexion…" : "Se connecter"}
                    </Button>
                </form>
            </Form>

            <br/>

            <Link to="/password-update" search={{token: undefined}} disabled={form.formState.isSubmitting}
                  className="justify-self-center rounded-md border px-3 py-2 text-sm font-medium hover:bg-accent hover:text-accent-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 focus-visible:ring-offset-background">
                Mot de passe oublié
            </Link>
        </div>
    );
}