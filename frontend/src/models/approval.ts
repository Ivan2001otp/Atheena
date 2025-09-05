export interface Approval {
    id: string;
    from_warehouse_name : string;
    from_warehouse_location : string;

    from_warehouse_state : string;
    from_warehouse_country : string;
    status : string;

    reason : string;
    supply_name : string;
    supervisor_name : string;

    updated_time : string;
}

export interface ApprovalResponse {
    message : string;
    success : boolean;
    data :  Approval[];
}