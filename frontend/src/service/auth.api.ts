import { ACCESS_TOKEN, type AddWarehouseRequest, type AdminAuthResponse, type AdminLogoutRequest, type AdminRegisterRequest, type AdminWarehouse, type LoginRequest, type SiteModel, type StandardResponse } from "@/models/auth";

import axios from "axios";
import { clearAuth, getAccessToken, getRefreshToken } from "./util";
import JsCookies from "js-cookie"
import toast from "react-hot-toast";


const BASE_URL = "http://localhost:8080/api/v1";

const axiosInstance = axios.create({
    baseURL: BASE_URL,
    withCredentials: true,
});

// Request interceptor to attach tokens at every request.
axiosInstance.interceptors.request.use(
    (config) => {
 const token = getAccessToken();

        if (token) {
            console.log("token exists...");
            config.headers["Authorization"] = `Bearer ${token}`;
        }

        return config;

    }, (error) => {
        console.log("something wrong , request interceptor");
        return Promise.reject(error);
    }
);



// Response interceptor
axiosInstance.interceptors.response.use(
    (response) => {


        return response;
    }, async (error) => {
        console.log("something went wrong in response - interceptor.");
        const originalRequest = error.config;
        console.log("error response status is ", error.response?.status);


        if (error.response?.status === 404) {
            originalRequest._retry = false;
        } else if (error.response?.status === 401) {
            originalRequest._retry = false;
            toast.error("Invalid Credentials");
            window.location.href="/login";
            // window.location.href = "/access_denied";
        }
        else if (error.response?.status == 403) {
            originalRequest._retry = false;
            console.log("Setting retry to false : ", originalRequest._retry);
            const cookie_refresh_token = getRefreshToken()

            if (cookie_refresh_token) {
                try {
                    const payload = {
                        refresh_token: cookie_refresh_token,
                    };

                    const res = await axiosInstance.post("/refresh-token", payload);
                    console.log(res);
                    JsCookies.set(ACCESS_TOKEN, res.data.access_token, { expires: 0.0104, secure: true, sameSite: "Strict" });


                    originalRequest.headers["Authorization"] = `Bearer ${res.data.access_token}`;
                    return axiosInstance(originalRequest);
                } catch (error) {
                    clearAuth();
                    setTimeout(
                        () => {
                            console.log("Session expired. Kindly login again")
                            // toast.error("Session expired. Kindly login again")
                        }, 2000
                    );

                    window.location.href = "/login";
                    return Promise.reject(error);
                }
            }
        }
    }
);


export const LoginAdmin = async (
    payload: LoginRequest
): Promise<AdminAuthResponse> => {

    try {
        const response = await axiosInstance.post(`${BASE_URL}/login`, payload);
        console.log("The status of LoginAdmin api is ", response.status);
        console.log(response);

        if (response.status === 200) {
            return response.data;
        }


    } catch (error: any) {
        console.log("UI - something went wrong during login")
        console.log("status is ", error)
        console.log(error)

    }
    return Promise.reject("Could not get 200 status while logging-in admin")
};


export const RegisterAdmin = async(
    payload:AdminRegisterRequest
) : Promise<AdminAuthResponse>=>{
    // post request
    const response = await axiosInstance.post(`${BASE_URL}/register`, payload);
    console.log("The status of RegisterAdmin api is ", response.status);


    if (response.status == 200) {
        return response.data;
    }

    return Promise.reject("Could not get 200 status while registering admin");
};

export const LogoutAdmin = async(
    payload : AdminLogoutRequest
): Promise<StandardResponse>=>{

    console.log("logout request body")
    console.log(payload);
    try {
        const response = await axiosInstance.post(`${BASE_URL}/logout`,{}, {
            params:{
                "email" : payload.email,
                "role" : payload.role
            },
        });
        // console.log(response);
        console.log("The status of LogoutAdmin api is ", response.status);

        if (response.status === 200) {
            return response.data;
        }

    } catch (error:any) {
        console.log("UI - something went wrong during logout")
        console.log(error)
    }
    
    return Promise.reject("Could not get 200 status while logout admin")

}


// Construction site apis
export const AddNewConstructionSite = async (payload: SiteModel): Promise<StandardResponse> =>{

    try {

        const response = await axiosInstance.post(`${BASE_URL}/add_construction_site`, payload);
        
        console.log(response);
        console.log("The status of Add-Ware-house api is ", response.status);

        if (response.status === 200) {
            return response.data;
        }

    } catch (error) {
        console.log("Something went wrong while adding new warehouse !");
        console.log(error);
    }

    return Promise.reject("Could not get 200 status while adding new warehouse")

}

export const GetAllConstructionSites = async (adminIdStr:string) : Promise<SiteModel[]> => {

    try {
        const res = await axiosInstance.get(`${BASE_URL}/get_construction_sites/${adminIdStr}`)
        console.log(res)

        if (res.status ===200)return res.data;

    } catch (error) {
        console.log("Something went wrong while fetching all construction sites...");
        console.log(error);
    }

    return Promise.reject("Could not get 200 status while fetching all construction sites.")

}


// Warehouse apis
export const GetAllWarehouse = async(adminIdStr : string
) : Promise<AdminWarehouse[]> => {


    try {
        const response = await axiosInstance.get(`${BASE_URL}/get_warehouses/${adminIdStr}`,)
        console.log(response);

        if (response.status === 200) {
            return response.data;
        }

    }catch(error) { 
        console.log("Something went wrong while fetching warehouse by admin id");
        console.log(error);
    }

    return Promise.reject("Could not get 200 status while fetching new warehouses")

}

export const AddNewWarehouse = async(
    payload : AddWarehouseRequest
) : Promise<StandardResponse> =>{ 
    
    try {
        const response = await axiosInstance.post(`${BASE_URL}/add_warehouse`, payload);

        console.log(response);
        console.log("The status of Add-Ware-house api is ", response.status);

        if (response.status === 200) {
            return response.data;
        }


    } catch (error) {
        console.log("Something went wrong while adding new warehouse !");
        console.log(error);
    }
    
    return Promise.reject("Could not get 200 status while adding new warehouse")
}