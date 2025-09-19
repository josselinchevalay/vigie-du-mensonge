import * as React from "react";
import {useForm} from "react-hook-form";
import {z} from "zod";
import {zodResolver} from "@hookform/resolvers/zod";
import {Form, FormControl, FormField, FormItem, FormLabel, FormMessage} from "@/core/shadcn/components/ui/form";
import {Input} from "@/core/shadcn/components/ui/input";
import {Button} from "@/core/shadcn/components/ui/button";
import {authManager} from "@/core/dependencies/auth/auth-manager";
import {useNavigate} from "@tanstack/react-router";
import {toast} from "sonner";
import { Eye, EyeOff } from "lucide-react";

const signInSchema = z.object({
    email: z.email("Adresse e-mail invalide"),
    password: z.string().min(1, "Mot de passe requis"),
});

export type SignInInput = z.infer<typeof signInSchema>;

export function SignInForm() {
    const navigate = useNavigate();
    const form = useForm<SignInInput>({
        resolver: zodResolver(signInSchema),
        defaultValues: {email: "", password: ""},
        mode: "onSubmit",
    });

    const [showPassword, setShowPassword] = React.useState(false);

    const onSubmit = async (values: SignInInput) => {
        try {
            await authManager.signIn({email: values.email, password: values.password});
            await navigate({to: "/"});
        } catch (e: unknown) {
            let status: number | undefined;
            if (e && typeof e === "object" && "response" in e) {
                // @ts-expect-error ky HTTPError duck typing
                status = e.response?.status;
            }

            const message = status === 401
                ? "Identifiants invalides."
                : status === 404
                    ? "Aucun compte ne correspond à cette adresse email."
                    : "Une erreur est survenue. Veuillez réessayer.";

            toast.error(message);
        }
    };

    return (
        <div className="mx-auto w-full max-w-sm">
            <Form {...form}>
                <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4">
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
                                                <EyeOff className="h-4 w-4" aria-hidden="true" />
                                            ) : (
                                                <Eye className="h-4 w-4" aria-hidden="true" />
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
        </div>
    );
}

export default SignInForm;
