
export interface Supervisor {
    id: string;
    admin_id: string;
    email: string;
    name: string;
    phone_number: string;
    role: string;
}

export interface CreateSupervisorRequest {
    name: string;
    admin_id: string;
    email: string;
    phone_number: string;
    role: string;
}

export interface FetchallSupervisorResponse {
    data: Supervisor[]
    message: string;
    success: boolean;
}