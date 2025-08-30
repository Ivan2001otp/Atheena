import { ACCESS_TOKEN, REFRESH_TOKEN } from "@/models/auth";
import Cookies from "js-cookie";
import {jwtDecode} from "jwt-decode";

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