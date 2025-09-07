import ApprovalCard from '@/components/custom/ApprovalCard';
import OrderCard from '@/components/custom/OrderCard';

import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs';
import type { Approval } from '@/models/approval';
import { ADMIN_ID } from '@/models/auth';
import type { OrderItem } from '@/models/order';
import { FetchAllApprovals, FetchOrders } from '@/service/auth.api';
import React, { useEffect, useState } from 'react'

const TransactionPage = () => {
  const [activeTab, setActiveTab] = useState<string>("approvals");

  const [approvalList, setApprovalList] = useState<Approval[]>([]);
  const [orders, setOrders] = useState<OrderItem[]>([]);

  useEffect(()=>{

    const fetchApprovalResponse = async() => {
       const adminId = localStorage.getItem(ADMIN_ID!);
       if (adminId){
          try {
            const res = await FetchAllApprovals(adminId);
            if (res.success) {
              setApprovalList(res.data);
            }

          } catch (error) {
            console.log(error);
          }
      }
    }

    const fetchOrderResponse = async() => {
       
        const adminId = localStorage.getItem(ADMIN_ID)!;

        if (adminId) {

          try {

            const res = await FetchOrders(adminId);
            if (res.success) {
              setOrders(res.data);
            }

          } catch (error) {
            console.log(error);
          }
        }
    }

    fetchOrderResponse();
    fetchApprovalResponse();
  }, []);


  return (


     <div className="min-h-screen flex flex-col">
      <main className="container mx-auto p-4 md:p-8 flex-grow font-sans">
        <div className="flex justify-between items-center mb-6">
          <h1 className="text-3xl font-bold tracking-tight text-gray-900">Supply Chain Dashboard</h1>
        </div>
        <Tabs defaultValue="approvals" onValueChange={setActiveTab}>
          <TabsList className="grid w-full grid-cols-2">
            <TabsTrigger value="approvals">Approvals</TabsTrigger>
            <TabsTrigger value="orders">Orders</TabsTrigger>
          </TabsList>
          <TabsContent value="approvals" className="mt-4">
            <div className="space-y-4">
              {approvalList.map((approval,index) => (
                <ApprovalCard key={index} approval={approval} />
              ))}
            </div>
          </TabsContent>
          <TabsContent value="orders" className="mt-4">
            <div className="space-y-4">
              {orders.map((order) => (
                <OrderCard key={order.order_id} order={order} />
              ))}
            </div>
          </TabsContent>
        </Tabs>
      </main>
    </div>

  )
}

export default TransactionPage