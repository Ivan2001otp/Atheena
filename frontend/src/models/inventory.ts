export interface InventoryItem {
    id : string;
    warehouse_id : string;
    name : string;
    quantity : number;
    units : string;
    reason:string;
    created_at : string;
    updated_at:string;
}

export interface AddInventoryRequest {
    warehouse_id : string;
    name : string;
    quantity : number;
    reason : string;
}

export interface InventoryResponse {
    message: string;
    success: boolean;
    data : InventoryItem[]
}