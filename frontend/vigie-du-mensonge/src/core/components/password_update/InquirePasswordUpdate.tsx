import {z} from "zod";
import {useForm} from "react-hook-form";
import {zodResolver} from "@hookform/resolvers/zod";
import {Form, FormControl, FormField, FormItem, FormLabel, FormMessage} from "@/core/shadcn/components/ui/form.tsx";
import {Input} from "@/core/shadcn/components/ui/input.tsx";
import {Button} from "@/core/shadcn/components/ui/button.tsx";
import {useState} from "react";

export type InquirePasswordUpdateProps = {
    submitForm: (input: InquirePasswordUpdateInput) => Promise<boolean>;
}

const formSchema = z.object({
    email: z.email(),
});

type InquirePasswordUpdateInput = z.infer<typeof formSchema>;

export function InquirePasswordUpdate({submitForm}: InquirePasswordUpdateProps) {
    const [success, setSuccess] = useState(false);

    const form = useForm<InquirePasswordUpdateInput>({
        resolver: zodResolver(formSchema),
        defaultValues: {email: ""},
        mode: "onSubmit",
    });

    const onSubmit = async (values: InquirePasswordUpdateInput) => {
        const result = await submitForm(values);
        setSuccess(result);
    };

    if (success) {
        return <p className="text-center p-1">
            L'email contenant le lien de modification a été envoyé.
        </p>;
    }

    return (
        <div className="mx-auto max-w-sm">
            <Form {...form}>
                <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4">
                    <div className="space-y-1">
                        <h1 className="text-xl font-semibold">Modification votre mot de passe</h1>
                        <p className="text-sm text-muted-foreground">Saisissez votre adresse email pour recevoir un lien
                            de modification sécurisé.</p>
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

                    <Button type="submit" disabled={form.formState.isSubmitting} className="w-full">
                        {form.formState.isSubmitting ? 'Envoi en cours...' : "Recevoir l'email de modification"}
                    </Button>
                </form>
            </Form>
        </div>
    );
}