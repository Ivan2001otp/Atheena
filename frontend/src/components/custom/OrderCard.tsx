import { DELIVERED, IN_TRANSIT, ORDER_PLACED, OUT_FOR_DELIVERY, type OrderItem } from '@/models/order'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Separator } from "@/components/ui/separator";
import { Badge } from "@/components/ui/badge";
import { format } from "date-fns";

const OrderCard = ({order} : {order : OrderItem}) => {

  const isCompleted = (status : string) => {
    const trackIndex = order.trackers.findIndex(t => t.order_status === status);
    const currIndex = order.trackers.findIndex(t => t.order_status === order.current_status);

    return trackIndex <= currIndex;
  };

  const statusColors = {
    DELIVERED:'bg-green-100 text-green-800',
    IN_TRANSIT : 'bg-blue-100 text-blue-80',
    ORDER_PLACED : 'bg-yellow-100 text-yellow-800',
    OUT_FOR_DELIVERY : 'bg-purple-100 text-purple-800',
  }


  return (
     <Card className="hover:shadow-lg transition-shadow duration-300 animate-in fade-in-0 slide-in-from-top-4">
      <CardHeader>
        <div className="flex justify-between items-start">
          <CardTitle className="text-xl">{order.material_name} - {order.order_type}</CardTitle>
          <Badge variant="secondary" className={`capitalize ${statusColors[order.current_status as keyof typeof statusColors] || 'bg-gray-100 text-gray-800'}`}>
            {order.current_status.replace(/_/g, ' ')}
          </Badge>
        </div>
        <CardDescription>
          Order ID: {order.order_id}
        </CardDescription>
      </CardHeader>
      <CardContent>
        <div className="grid grid-cols-2 gap-4 mb-6">
          <div className="text-sm">
            <p className="text-muted-foreground">Quantity</p>
            <p className="font-medium">{order.quantity} {order.unit}</p>
          </div>
          <div className="text-sm">
            <p className="text-muted-foreground">Type</p>
            <p className="font-medium">{order.order_type}</p>
          </div>
        </div>
        <Separator className="mb-6" />
        
        {/* Animated Tracker Stepper UI */}
        <div className="flex flex-col space-y-4">
          <h3 className="text-lg font-semibold">Order Progress</h3>
          {order.trackers.map((tracker, index) => (
            <div key={index} className="flex items-center space-x-4 relative">
              {/* Stepper Dot */}
              <div className={`w-3 h-3 rounded-full transition-all duration-300 transform scale-100 ${isCompleted(tracker.order_status) ? 'bg-green-500' : 'bg-gray-300'}`}></div>
              {/* Vertical Line */}
              {index < order.trackers.length - 1 && (
                <div className={`absolute left-1.5 top-3.5 w-0.5 h-full -z-10 transition-all duration-300 ${isCompleted(order.trackers[index+1].order_status) ? 'bg-green-500' : 'bg-gray-300'}`}></div>
              )}
              {/* Tracker Details */}
              <div className="flex-1">
                <p className={`font-medium ${isCompleted(tracker.order_status) ? 'text-gray-900' : 'text-gray-500'}`}>
                  {tracker.order_status.replace(/_/g, ' ')}
                </p>
                <p className="text-xs text-muted-foreground">
                  {format(new Date(tracker.created_time), "PPP 'at' p")}
                </p>
              </div>
            </div>
          ))}
        </div>
      </CardContent>
    </Card>
  )
}

export default OrderCard