import React from "react";
import {Progress} from "@/core/shadcn/components/ui/progress.tsx";

export function BasicProgress() {
    const [progress, setProgress] = React.useState(13);
    React.useEffect(() => {
        const timer = setTimeout(() => setProgress(66), 500);
        return () => clearTimeout(timer);
    }, []);
    return <Progress value={progress} className="w-[60%]"/>;
}