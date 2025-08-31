export type Role = "ADMIN | SUPERVISOR"

export const ACCESS_TOKEN : string = "access_token"
export const REFRESH_TOKEN : string = "refresh_token"
export const ADMIN_EMAIL:string = "admin_email_credential@";
export const ADMIN_NAME : string = "admin_name_credential@" 
export const ADMIN_ROLE : string = "admin_role_credential@"
export const ADMIN_ID : string = "admin_id_credential@"


export interface AdminWarehouse {
    id : string;
    user_id : string;
    name : string;
    pin:string;
    address: string;
    state:string;
    country:string;
    created_at : string
}

export interface AddWarehouseRequest{
    id:string;
    user_id : string;
    name : string;
    pin:string;
    address:string;
    state:string;
    country:string;
}


export interface StandardResponse{
    message : string;
    success : boolean;
}

export interface AdminLogoutRequest {
    email : string;
    role :string;
}



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