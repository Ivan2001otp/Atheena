export interface KPI {
    title : string;
    value : string;
    description : string;
}

export interface Item {
    id : string;
    name : string;
    location : string;
    stock : number;
}

export interface TransactionDataPoint {
    date : string;
    inbound : number;
    outbound : number;
}

export interface LocationDataPoint {
    name : string;
    count : number
}

export interface  StatusDataPoint {
    name : string;
    value : number;
}


export const kpiData: KPI[] = [
  {
    title: "Total Inventory Value",
    value: "$1.2M",
    description: "Across all locations",
  },
  {
    title: "Total Items in Stock",
    value: "15,489",
    description: "Unique items",
  },
  {
    title: "Out-of-Stock Items",
    value: "12",
    description: "Requires immediate reorder",
  },
  {
    title: "Number of Locations",
    value: "5",
    description: "Warehouses and stores",
  },
];

export const lowStockItems: Item[] = [
  { id: "item-001", name: "Cement Bags", location: "Warehouse A", stock: 5 },
  { id: "item-002", name: "Steel Rods", location: "Warehouse B", stock: 8 },
  { id: "item-003", name: "Concrete Blocks", location: "Site 1", stock: 10 },
  { id: "item-004", name: "Safety Helmets", location: "HQ", stock: 2 },
];

export const popularItems: Item[] = [
  { id: "item-005", name: "Wood Planks", location: "Site 2", stock: 500 },
  { id: "item-006", name: "Copper Wiring", location: "Warehouse C", stock: 350 },
  { id: "item-007", name: "Insulation Foam", location: "Site 1", stock: 200 },
  { id: "item-008", name: "Plastic Pipes", location: "Warehouse A", stock: 150 },
];

export const transactionTrendsData: TransactionDataPoint[] = [
  { date: "Jan", inbound: 50, outbound: 30 },
  { date: "Feb", inbound: 70, outbound: 55 },
  { date: "Mar", inbound: 65, outbound: 75 },
  { date: "Apr", inbound: 90, outbound: 80 },
  { date: "May", inbound: 120, outbound: 110 },
  { date: "Jun", inbound: 150, outbound: 140 },
];

export const inventoryByLocationData: LocationDataPoint[] = [
  { name: "Warehouse A", count: 400 },
  { name: "Warehouse B", count: 300 },
  { name: "Warehouse C", count: 300 },
  { name: "Site 1", count: 200 },
  { name: "Site 2", count: 250 },
];

export const inventoryStatusData: StatusDataPoint[] = [
  { name: "In Stock", value: 80 },
  { name: "On Reorder", value: 15 },
  { name: "Out of Stock", value: 5 },
];

export const PIE_COLORS = ["#0088FE", "#FFBB28", "#FF8042"];
