import { Button } from '@/components/ui/button';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { ADMIN_ID, type AdminWarehouse } from '@/models/auth';
import { GetAllWarehouse } from '@/service/auth.api';
import { motion } from 'framer-motion';
import { PlusCircle, Warehouse } from 'lucide-react';
import React, { useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom';

const InventoryPage = () => {
  const [warehouseList, setWarehouseList] = useState<AdminWarehouse[]>([]);
  

  const navigate = useNavigate();

  useEffect(() => {

    const fetchWarehouses = async() =>{

      try {
      
        const res = await GetAllWarehouse(localStorage.getItem(ADMIN_ID)!);
        setWarehouseList(res);
      } catch(error) {
      
        console.log(error);
      }

    }

    fetchWarehouses()

  }, []);

  return (
    <div
      className='min-h-screen bg-gradient-to-b from-slate-50 to-slate-100 p-8'
    >
      <div className='flex items-center justify-between mb-8'>
          <div>
            <h1 className='text-3xl font-bold tracking-tight'>Warehouses</h1>
            <p className='text-slate-600'>Total warehouses : {warehouseList.length}</p>
          </div>
      </div>
       


      <div className='grid gap-6 sm:grid-cols-2 lg:grid-cols-3'>
        {
          warehouseList.map((warehouse, index) => (
            <motion.div
              key={warehouse.id}
              initial={{opacity:0 , y:30}}
              animate={{opacity:1, y:0}}
              transition={{delay:index*0.1}}
            >
                <Card className='hover:shadow-xl transition-shadow rounded-2xl'>
                    <CardHeader  className='flex items-center gap-3'>
                        <Warehouse className='h-8 w-8 text-blue-600'/>
                        <CardTitle className='text-lg'>{warehouse.name}</CardTitle>
                    </CardHeader>

                    <CardContent>
                        <p>
                          State : {warehouse.state}
                        </p>
                        <p>
                          Country : {warehouse.country}
                        </p>

                        <Button className='mt-4 w-full cursor-pointer'
                        onClick={()=>{
                            navigate("/atheena/inv-v2", {
                              state: {
                                warehouse_id:warehouse.id,
                              },
                            })
                        }}

                        >View Inventory</Button>
                    </CardContent>
                </Card>
            </motion.div>
          ))
        }
      </div>
    </div>
  );
}

export default InventoryPage