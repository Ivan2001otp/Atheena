import { ACCESS_TOKEN, REFRESH_TOKEN } from "@/models/auth";
import Cookies from "js-cookie";
import {jwtDecode} from "jwt-decode";
import { number } from "motion-dom";

interface JwtPayload {
    exp:number;
    email:string;
    role:string;
}

export const getAccessToken = ()  => Cookies.get(ACCESS_TOKEN);
export const getRefreshToken = () => Cookies.get(REFRESH_TOKEN);
export const clearAuth = () => {
    Cookies.remove(ACCESS_TOKEN);
    Cookies.remove(REFRESH_TOKEN);
}

// use this method to parse date while adding item to your ui state.
export function formatDateTimeV2(date: Date | number) : string {
    const d = typeof date==="number" ? new Date(date) : date;

    return d.toLocaleDateString("en-US", {
        month:"short",
        day:"2-digit",
        year:"numeric",
    }).replace(/ /g," ");
}

// use this to parse the date time value from backend api
export function formatDateTime(dateTimeStr : string) : string {
    const date = new Date(dateTimeStr);

    return date.toLocaleDateString("en-GB", {
        day : "2-digit",
        month:"short",
        year:"numeric",
    }).replace(/ /g," ");
}

export const isTokenExpired = (token:string) : boolean => {
    try {
        const decoded = jwtDecode<JwtPayload>(token);
        /*
        {
            email : "sidehustle681@gmail.com",
            expiry : 1756359504 ,
            role : "ADMIN"
        }
        */
        const result = decoded.exp * 1000 < Date.now();;
        if (result) {
            console.log("token is expired");
        } else {
            console.log("token is not expired");;
        }

        return decoded.exp * 1000 < Date.now();
    } catch(error: any) {
        console.log("Something went wrong while decoding jwt token");
        console.log(error); 
        return true;
    }
}