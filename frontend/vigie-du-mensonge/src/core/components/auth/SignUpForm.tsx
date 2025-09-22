import * as React from "react";
import {useForm} from "react-hook-form";
import {z} from "zod";
import {zodResolver} from "@hookform/resolvers/zod";
import {Form, FormControl, FormField, FormItem, FormLabel, FormMessage} from "@/core/shadcn/components/ui/form";
import {Input} from "@/core/shadcn/components/ui/input";
import {Button} from "@/core/shadcn/components/ui/button";
import {authManager} from "@/core/dependencies/auth/authManager.ts";
import {useNavigate} from "@/core/utils/router";
import {toast} from "@/core/utils/toast";
import {Eye, EyeOff} from "lucide-react";
import {HTTPError} from "ky";

// Schema aligned with AuthClient.signUp requirements
const signUpSchema = z
    .object({
        email: z.email("Adresse e-mail invalide"),
        password: z
            .string()
            .min(8, "Au moins 8 caractères")
            .regex(/[a-z]/, "Au moins une lettre minuscule (a-z)")
            .regex(/[A-Z]/, "Au moins une lettre majuscule (A-Z)")
            .regex(/[0-9]/, "Au moins un chiffre (0-9)")
            .regex(/[^A-Za-z0-9]/, "Au moins un caractère spécial (ex: &!$?;:#@)"),
        confirmPassword: z.string(),
    })
    .refine((data) => data.password === data.confirmPassword, {
        message: "Les mots de passe ne correspondent pas",
        path: ["confirmPassword"],
    });

export type SignUpInput = z.infer<typeof signUpSchema>;

export function SignUpForm() {
    const navigate = useNavigate();
    const form = useForm<SignUpInput>({
        resolver: zodResolver(signUpSchema),
        defaultValues: {email: "", password: "", confirmPassword: ""},
        mode: "onSubmit",
    });

    const [showPassword, setShowPassword] = React.useState(false);

    const onSubmit = async (values: SignUpInput) => {
        try {
            await authManager.signUp({
                email: values.email,
                password: values.password,
            });

            // Navigate to home (or another page if needed later)
            await navigate({to: "/email-verification", search: {token: undefined}});
        } catch (e) {
            let status: number | undefined;
            if (e instanceof HTTPError) {
                status = e.response.status;
            }
            const message = status === 409
                ? 'Cette adresse email est déjà associée à un compte.'
                : 'Une erreur est survenue. Veuillez réessayer.';
            toast.error(message);
        }
    };

    return (
        <div className="mx-auto w-full max-w-sm">
            <Form {...form}>
                <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4">
                    <div className="space-y-1">
                        <h1 className="text-xl font-semibold">Créer un compte</h1>
                        <p className="text-sm text-muted-foreground">Rejoignez Vigie du mensonge</p>
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
                                            autoComplete="new-password"
                                            className="pr-10"
                                            {...field}
                                        />
                                    </div>
                                </FormControl>
                                <FormMessage/>
                            </FormItem>
                        )}
                    />

                    <FormField
                        control={form.control}
                        name="confirmPassword"
                        render={({field}) => (
                            <FormItem>
                                <FormLabel>Confirmer le mot de passe</FormLabel>
                                <FormControl>
                                    <div className="relative">
                                        <Input
                                            type={showPassword ? "text" : "password"}
                                            placeholder="••••••••"
                                            autoComplete="new-password"
                                            className="pr-10"
                                            {...field}
                                        />
                                    </div>
                                </FormControl>
                                <FormMessage/>
                            </FormItem>
                        )}
                    />

                    <div className="justify-self-center">
                        <button type="button" onClick={() => setShowPassword((visibility) => !visibility)}>
                            {showPassword ? <EyeOff/> : <Eye/>}
                        </button>
                    </div>

                    <Button type="submit" disabled={form.formState.isSubmitting} className="w-full">
                        {form.formState.isSubmitting ? "Création…" : "Créer le compte"}
                    </Button>
                </form>
            </Form>
        </div>
    );
}
