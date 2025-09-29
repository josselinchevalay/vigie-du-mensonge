import {z} from "zod";
import {useForm} from "react-hook-form";
import {zodResolver} from "@hookform/resolvers/zod";
import {Form, FormControl, FormField, FormItem, FormLabel, FormMessage} from "@/core/shadcn/components/ui/form.tsx";
import {Input} from "@/core/shadcn/components/ui/input.tsx";
import {Button} from "@/core/shadcn/components/ui/button.tsx";
import {useState} from "react";

const formSchema = z.object({
    email: z.email(),
});

type InquireSignUpInput = z.infer<typeof formSchema>;

export type InquireSignUpProps = {
    submitForm: (input: InquireSignUpInput) => Promise<boolean>;
};

export function InquireSignUp({submitForm}: InquireSignUpProps) {
    const [success, setSuccess] = useState(false);

    const form = useForm<InquireSignUpInput>({
        resolver: zodResolver(formSchema),
        mode: "onSubmit",
        defaultValues: {email: ''},
    });

    const onSubmit = async (values: InquireSignUpInput) => {
        const result = await submitForm(values);
        setSuccess(result);
    };

    if (success) {
        return <p className="text-center p-1">
            L'email contenant le lien d'inscription a été envoyé.
        </p>;
    }

    return <div className="mx-auto max-w-sm">
        <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4">
                <div className="space-y-1">
                    <h1 className="text-xl font-semibold">Inscription</h1>
                    <p className="text-sm text-muted-foreground">Saisissez votre adresse email pour recevoir un lien
                        d'inscription sécurisé.</p>
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
                    {form.formState.isSubmitting ? 'Envoi en cours...' : "Recevoir l'email d'inscription"}
                </Button>
            </form>
        </Form>
    </div>;
}

