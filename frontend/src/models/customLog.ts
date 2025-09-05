export interface CustomLog {
    _id : string;
    from_warehouse_name : string;
    from_warehouse_location : string;
    from_warehouse_state : string;
    from_warehouse_country : string;

    to_destination_name: string;
    to_destination_location : string;
    to_destination_state : string;
    to_destination_country : string;
    is_site : boolean;
    
    supply_name : string;
    supply_quantity : number;
    supply_unit :string;
    updated_time : string;
}

export interface FetchLogsResponse {
    data : CustomLog[];
    message: string;
    success: boolean;
}