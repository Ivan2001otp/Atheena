
export interface Supervisor {
    _id : string;
    admin_id : string;
    email : string;
    name : string;
    phone_number : string;
    role : string;
}

export interface FetchallSupervisorResponse{
    data : Supervisor[]
    message : string;
    status : boolean;
}