export interface OrderItem {
    order_id:string;
    material_name : string;
    quantity: number;

    unit:string;
    order_type:string;
    current_status:string;
    trackers : OrderTracker[]

}

export interface OrderTracker {
    order_status : string;
    created_time : string; 
}

export const  DELIVERED : string = "DELIVERED";
export const IN_TRANSIT : string = "IN_TRANSIT";
export const ORDER_PLACED : string = "ORDER_PLACED";
export const OUT_FOR_DELIVERY : string = "OUT_FOR_DELIVERY";

export interface OrderResponse {
    success : boolean;
    message: string;
    data : OrderItem[]
}