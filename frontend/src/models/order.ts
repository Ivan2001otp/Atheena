export interface OrderItem {
    order_id:string;
    material_name : string;
    quantity: number;

    unit:string;
    order_type:string;
    current_status:string;
    order_trackers : OrderTracker[]
}

export interface OrderTracker {
    order_status : string;
    created_time : string; 
}

export interface OrderResponse {
    success : boolean;
    message: string;
    data : OrderItem[]
}