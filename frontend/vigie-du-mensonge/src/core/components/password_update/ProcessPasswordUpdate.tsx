import {z} from "zod";
import * as React from "react";
import {useState} from "react";
import {useForm} from "react-hook-form";
import {zodResolver} from "@hookform/resolvers/zod";
import {Form, FormControl, FormField, FormItem, FormLabel, FormMessage} from "@/core/shadcn/components/ui/form.tsx";
import {Input} from "@/core/shadcn/components/ui/input.tsx";
import {Eye, EyeOff} from "lucide-react";
import {Button} from "@/core/shadcn/components/ui/button.tsx";

export type ProcessPasswordUpdateProps = {
    submitForm: (input: ProcessPasswordUpdateInput) => Promise<boolean>;
}

const formSchema = z.object({
    password: z
        .string()
        .min(8, "Au moins 8 caractères")
        .regex(/[a-z]/, "Au moins une lettre minuscule (a-z)")
        .regex(/[A-Z]/, "Au moins une lettre majuscule (A-Z)")
        .regex(/[0-9]/, "Au moins un chiffre (0-9)")
        .regex(/[^A-Za-z0-9]/, "Au moins un caractère spécial (ex: &!$?;:#@)"),
    confirmPassword: z.string(),
}).refine((data) => data.password === data.confirmPassword, {
    message: "Les mots de passe ne correspondent pas",
    path: ["confirmPassword"],
});

export type ProcessPasswordUpdateInput = z.infer<typeof formSchema>;

export function ProcessPasswordUpdate({submitForm}: ProcessPasswordUpdateProps) {
    const [success, setSuccess] = useState(false);

    const form = useForm<ProcessPasswordUpdateInput>({
        resolver: zodResolver(formSchema),
        defaultValues: {password: "", confirmPassword: ""},
        mode: "onSubmit",
    });

    const [showPassword, setShowPassword] = React.useState(false);

    const onSubmit = async (values: ProcessPasswordUpdateInput) => {
        const result = await submitForm(values);
        setSuccess(result);
    };

    if (success) {
        return <>
            Votre mot de passe a été mis à jour.
        </>;
    }

    return (
        <div className="mx-auto w-full max-w-sm">
            <Form {...form}>
                <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4">
                    <div className="space-y-1">
                        <h1 className="text-xl font-semibold">Saisissez votre nouveau mot de passe</h1>
                    </div>

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
                        {form.formState.isSubmitting ? "Envoi cours…" : "Modifier le mot de passe"}
                    </Button>
                </form>
            </Form>
        </div>
    );
}