import {Moon, Sun} from "lucide-react";
import {Button} from "@/core/shadcn/components/ui/button.tsx";
import {useTheme} from "@/core/components/theme/theme-utils.tsx";

export function ThemeToggle() {
    const {theme, setTheme} = useTheme();

    if (theme === "dark") {
        return <Button variant="ghost"
                       onClick={() => setTheme("light")}>
            <Sun/>
        </Button>;
    }

    return <Button variant="ghost"
                   onClick={() => setTheme("dark")}>
        <Moon/>
    </Button>;
}