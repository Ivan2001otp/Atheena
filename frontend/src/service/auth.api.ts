import { ACCESS_TOKEN, type AdminAuthResponse, type AdminRegisterRequest, type LoginRequest } from "@/models/auth";
import axios from "axios";
import { clearAuth, getAccessToken, getRefreshToken } from "./util";
import JsCookies from "js-cookie" 
import toast from "react-hot-toast";

const BASE_URL = "http://localhost:8080/api/v1";

const axiosInstance = axios.create({
    baseURL : BASE_URL,
    withCredentials: true,
});

// Request interceptor to attach tokens at every request.
axiosInstance.interceptors.request.use(
    (config)=>{
        const token = getAccessToken();
        
        if (token) {
            console.log("token exists...");
            config.headers["Authorization"]= `Bearer ${token}`;
        } else {
            console.log("token does not exists...");
        }
        
        return config;

    },(error)=>{
        console.log("something wrong , request interceptor");
        return Promise.reject(error);
    }
);



// Response interceptor
axiosInstance.interceptors.response.use(
    (response)=>{

        
        return response;
    },async (error) => {
        console.log("something went wrong in response - interceptor.");
        const originalRequest = error.config;
        if (error.response?.status === 404) {
            originalRequest._retry = false;
        }  else if(error.response?.status === 403) {
            originalRequest._retry = false;
            toast.error("You are not allowed to access protected resources.");
            window.location.href = "/access_denied";
        } 
        else if (error.response?.status == 401) {
            originalRequest._retry = false;
            const cookie_refresh_token = getRefreshToken()

            if (cookie_refresh_token) {
                try {
                    const payload = {
                        refresh_token : cookie_refresh_token,
                    };

                    const res = await axiosInstance.post("/refresh-token", payload);
                    console.log(res);
                    JsCookies.set(ACCESS_TOKEN, res.data.access_token, {expires: 0.0104,httpOnly:true, secure: true, sameSite: "Strict"});
                
                    
                    originalRequest.headers["Authorization"] = `Bearer ${res.data.access_token}`;
                    return axiosInstance(originalRequest);
                } catch (error) {
                    clearAuth();
                    setTimeout(
                        ()=>{
                            toast.error("Session expired. Kindly login again")
                        }, 2000
                    );

                    window.location.href="/login";
                    return Promise.reject(error);
                }
            }
        }
    }   
);



export const LoginAdmin = async(
    payload : LoginRequest
) : Promise<AdminAuthResponse> => {

    const response = await axiosInstance.post(`${BASE_URL}/login`, payload);
    console.log("The status of LoginAdmin api is ", response.status);
    console.log(response);

    if (response.status == 200) {
        return response.data;
    }

    return Promise.reject("Could not get 200 status while logging-in admin")
}

export const RegisterAdmin = async(
    payload:AdminRegisterRequest
) : Promise<AdminAuthResponse>=>{
    // post request
    const response = await axiosInstance.post(`${BASE_URL}/register`, payload);
    console.log("The status of RegisterAdmin api is ", response.status);
    // console.log(response);

    if (response.status == 200) {
       return response.data;
    }

    return Promise.reject("Could not get 200 status while registering admin");
};