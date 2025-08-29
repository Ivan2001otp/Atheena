import type { AdminAuthResponse, AdminRegisterRequest } from "@/models/auth";
import axios from "axios";

const BASE_URL = "http://localhost:8080/api/v1";

const axiosInstance = axios.create({
    baseURL : BASE_URL,
    withCredentials: true,
});

// Request interceptor to attach tokens at every request.
axiosInstance.interceptors.request.use(
    (config)=>{

        return config;
    },(error)=>{

        return Promise.reject(error);
    }
);


axiosInstance.interceptors.response.use(
    (response)=>{
        return response;
    },async (error) => {

    }
);



export const RegisterAdmin = async(
    payload:AdminRegisterRequest
) : Promise<AdminAuthResponse>=>{
    // post request
    const response = await axiosInstance.post(`${BASE_URL}/register`, payload);
    console.log("The status of RegisterAdmin api is ", response.status);
    console.log(response);

    if (response.status == 200) {
       return response.data;
    }

    return Promise.reject("Could not get 200 status while registering admin");
};