import { ADMIN_EMAIL, ADMIN_ID, ADMIN_NAME, ADMIN_ROLE } from '@/models/auth';
import { useState } from "react";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import {
  ResponsiveContainer,
  BarChart,
  Bar,
  XAxis,
  YAxis,
  Tooltip,
  Legend,
  AreaChart,
  Area,
  PieChart,
  Pie,
  Cell,
} from "recharts";
import { kpiData, lowStockItems, popularItems, transactionTrendsData } from '@/models/dashboard';


const DashboardV2Page = () => {

  return (
    <div className='flex min-h-screen w-full flex-col bg-gray-100 p-4 font-sans sm:p-6 lg:p-8'>
      <header className='mb-8'>
        <h1 className='text-3xl font-bold tracking-tight text-gray-900'>Dashboard</h1>
      
        <p className='text-lg text-gray-500'>A live overview of all inventory accross your locations.</p>
      </header>

      <section className='grid gap-6 sm:grid-cols-2 md:grid-cols-4'> 
        {
          kpiData.map((kpi, index) => (
            <Card
              key={index}
              className='rounded-xl transition-transform duration-300 hover:scale-[1.02] animate-in fade-in-0 zoom-in-95'
            >
              <CardHeader className='flex flex-row items-center justify-between space-y-0 pb-2'>
                <CardTitle className='text-sm  font-medium'>
                  {kpi.title}
                </CardTitle>
              </CardHeader>
              
              <CardContent>
                <div className='text-2xl font-bold'>{kpi.value}</div>
                <p className='text-xs text-muted-foreground'>{kpi.description}</p>
              </CardContent>
            </Card>
          ))
        }
      </section>

      <section className='mt-8 '>
          <Card>
            <CardHeader>
              <CardTitle>Recent Transaction Trends</CardTitle>
              <CardDescription>
                Inbound vs Outbound transaction over last 6 months.
              </CardDescription>

              <CardContent className='h-[300px]'>
                <ResponsiveContainer width="100%" height="100%">
                    <AreaChart data={transactionTrendsData}>

                        <XAxis  dataKey="date"/>
                        <YAxis />
                        <Tooltip/>
                        <Legend/>

                        <Area 
                          type="monotone"
                          dataKey="inbound"
                          stroke='#8884d8'
                          fill='#8884d8'
                        />

                        <Area
                        type="monotone"
                        dataKey="outbound"
                        stroke="#82ca9d"
                        fill="#82ca9d"
                      />
                    </AreaChart>
                </ResponsiveContainer>
              </CardContent>
            </CardHeader>
          </Card>
      </section>

      <section className='mt-8 grid gap-6 md:grid-cols-2'>
          <Card className='rounded-xl animate-in fade-in-0 slide-in-from-left-4'>
            <CardHeader>
              <CardTitle>Low Stock Items</CardTitle>
              <CardDescription>Items that need to be reordered soon.</CardDescription>
            </CardHeader>

             <CardContent>
            <div className="space-y-4">
              {lowStockItems.map((item) => (
                <div key={item.id} className="flex items-center justify-between">
                  <div>
                    <p className="font-medium">{item.name}</p>
                    <p className="text-sm text-muted-foreground">
                      {item.location}
                    </p>
                  </div>
                  <Badge variant="destructive">
                    Stock: {item.stock}
                  </Badge>
                </div>
              ))}
            </div>
          </CardContent>
          </Card>

          <Card className="rounded-xl animate-in fade-in-0 slide-in-from-right-4">
          <CardHeader>
            <CardTitle>Top Popular Items</CardTitle>
            <CardDescription>
              Most frequently transacted items this month.
            </CardDescription>
          </CardHeader>
          <CardContent>
            <div className="space-y-4">
              {popularItems.map((item) => (
                <div key={item.id} className="flex items-center justify-between">
                  <div>
                    <p className="font-medium">{item.name}</p>
                    <p className="text-sm text-muted-foreground">
                      {item.location}
                    </p>
                  </div>
                  <Badge variant="secondary">
                    Stock: {item.stock}
                  </Badge>
                </div>
              ))}
            </div>
          </CardContent>
        </Card>
      </section>
    </div>
  )
}

export default DashboardV2Page