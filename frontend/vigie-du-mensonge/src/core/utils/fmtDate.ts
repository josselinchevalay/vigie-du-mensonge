import {format} from "date-fns";
import {fr} from "date-fns/locale";

export function fmtDate(date: Date | number, fmt: string = "dd/MM/yyyy") {
    return format(date, fmt, {locale: fr});
}