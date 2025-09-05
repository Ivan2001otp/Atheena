import React from 'react'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { format } from "date-fns";
import type { Approval } from '@/models/approval';

interface ApprovalCardProps {
    approval : Approval
}

const ApprovalCard = ({approval} : ApprovalCardProps) => {
  return (
    <Card className='hover:shadow-lg transition-shadow duration-300'>
        <CardHeader>
            <CardTitle className='text-xl'>
                {approval.supply_name}
            </CardTitle>
            <CardDescription>
               Asked by **{approval.supervisor_name}**
            </CardDescription>
        </CardHeader>


        <CardContent>
            <div className='grid grid-cols-1 md:grid-cols-2 gap-4'>
                <div>
                    <p className='text-sm text-muted-foreground'>From Warehouse</p>
                    <p className='font-medium'>{approval.from_warehouse_name}</p>
                </div>

                <div>
                    <p className="text-sm text-muted-foreground">Location</p>
                    <p className="font-medium">
                    {approval.from_warehouse_location}, {approval.from_warehouse_state}, {approval.from_warehouse_country}
                    </p>
                </div>

                 <div>
                    <p className="text-sm text-muted-foreground">Status</p>
                    <p className={`font-semibold capitalize ${approval.status === 'completed' ? 'text-green-500' : 'text-yellow-500'}`}>
                    {approval.status}
                    </p>
                </div>

                 <div>
                    <p className="text-sm text-muted-foreground">Reason</p>
                    <p className="font-medium italic">"{approval.reason}"</p>
                </div>
            </div>
            <p className="mt-4 text-xs text-right text-muted-foreground">
                Last Updated: {format(new Date(approval.updated_time), 'PPp')}
            </p>
        </CardContent>
    </Card>
  )
}

export default ApprovalCard