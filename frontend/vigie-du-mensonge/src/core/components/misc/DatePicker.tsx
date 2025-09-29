import React from "react";
import {Input} from "@/core/shadcn/components/ui/input.tsx";
import {Popover, PopoverContent, PopoverTrigger} from "@/core/shadcn/components/ui/popover.tsx";
import {Button} from "@/core/shadcn/components/ui/button.tsx";
import {CalendarIcon} from "lucide-react";
import {Calendar} from "@/core/shadcn/components/ui/calendar.tsx";
import {fmtDate} from "@/core/utils/fmtDate.ts";
import {fr} from "date-fns/locale";

function isValidDate(date: Date | undefined) {
    if (!date) {
        return false;
    }
    return !isNaN(date.getTime());
}

export function DatePicker({date, setDate}: { date: Date | undefined, setDate: (date: Date | undefined) => void }) {
    const [open, setOpen] = React.useState(false);

    const [month, setMonth] = React.useState<Date | undefined>(date);
    const [value, setValue] = React.useState(fmtDate(date ?? new Date().setHours(0, 0, 0, 0)));

    return (
        <div className="relative flex gap-2">
            <Input
                id="date"
                value={value}
                placeholder="dd/mm/yyyy"
                className="bg-background pr-10"
                onChange={(e) => {
                    const date = new Date(e.target.value);
                    setValue(e.target.value);
                    if (isValidDate(date)) {
                        setDate(date);
                        setMonth(date);
                    }
                }}
                onKeyDown={(e) => {
                    if (e.key === "ArrowDown") {
                        e.preventDefault();
                        setOpen(true);
                    }
                }}
            />
            <Popover open={open} onOpenChange={setOpen}>
                <PopoverTrigger asChild>
                    <Button
                        id="date-picker"
                        variant="ghost"
                        className="absolute top-1/2 right-2 size-6 -translate-y-1/2"
                    >
                        <CalendarIcon className="size-3.5"/>
                    </Button>
                </PopoverTrigger>
                <PopoverContent
                    className="w-auto overflow-hidden p-0"
                    align="end"
                    alignOffset={-8}
                    sideOffset={10}
                >
                    <Calendar
                        mode="single"
                        selected={date}
                        captionLayout="dropdown"
                        month={month}
                        onMonthChange={setMonth}
                        locale={fr}
                        disabled={(date) => date > new Date()}
                        onSelect={(date) => {
                            setOpen(false);
                            if (date) {
                                setDate(date);
                                setValue(fmtDate(date));
                            }
                        }}
                    />
                </PopoverContent>
            </Popover>
        </div>
    );
}
