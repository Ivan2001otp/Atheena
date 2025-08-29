export type Role = "ADMIN | SUPERVISOR"

export const ACCESS_TOKEN : string = "access_token"
export const REFRESH_TOKEN : string = "refresh_token"

export interface AdminRegisterRequest {
    name:string;
    password:string;
    email:string;
    role:Role;
}

export interface AdminAuthResponse {
    access_token : string;
    message: string;
    refresh_token: string;
    name : string;
    email : string;
    id : string;
    role:string;
}



export interface LoginRequest {
    email:string;
    password:string;
}