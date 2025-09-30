import {QueryClient} from "@tanstack/react-query";

export const queryClient = new QueryClient();

queryClient.removeQueries({
    predicate: (query) => {
        const key = query.queryKey[0];
        return key === "redactor" || key === "moderator" || key === "admin";
    },
});