import { Button } from '@/components/ui/button';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Input } from '@/components/ui/input';
import type { AddInventoryRequest, InventoryItem } from '@/models/inventory';
import { AddInventoryOfSpecificWarehouse, FetchInventoryByWarehouseId } from '@/service/auth.api';
import { AnimatePresence, motion } from 'framer-motion';
import React, { useEffect, useState } from 'react'
import toast from 'react-hot-toast';
import { useLocation } from 'react-router-dom'

const InventoryV2Page = () => {

    const location = useLocation();
    const {warehouse_id} = location.state || "dfasf"

    const [inventoryItems, setInventoryItems] = useState<InventoryItem[]>([])
    const [form, setForm] = useState({name:"", quantity:"", units:"", reason:""});
    const [search, setSearch] = useState("");
    useEffect(()=>{

      const fetchInventory = async() => {

        if (warehouse_id === "dfasf") {
          window.location.href = "/login";
        }

        try {
          const result = await FetchInventoryByWarehouseId(warehouse_id);
          toast.success(result.message);
          setInventoryItems(result.data);
        } catch (error) {
          console.log(error);
        }
      };

      fetchInventory();
    },[warehouse_id]);

    const handleChange = (e : React.ChangeEvent<HTMLInputElement>) => {
      setForm({...form, [e.target.name]:e.target.value});
    };

    const handleSubmit = async(e: React.FormEvent) => {
      e.preventDefault();

      if (!form.name || !form.quantity || !form.quantity) return;


      const newItem : InventoryItem= {
         id: String(Date.now()),
        warehouse_id: String(Date.now())+"w1",
        name: form.name,
        quantity: Number(form.quantity),
        units: form.units,
        reason: form.reason,
        created_at: new Date().toISOString(),
        updated_at: new Date().toISOString(),
      };

      setInventoryItems([newItem, ...inventoryItems]);
      setForm({name:"", quantity:"", units:"", reason:""});


      try {
        const payload : AddInventoryRequest = {
          name : form.name,
          warehouse_id : warehouse_id,
          quantity : Number(form.quantity),
          reason: form.reason,
        }
         const res = await AddInventoryOfSpecificWarehouse(payload);

         if (res.success) {
          toast.success(res.message); 
         }
      } catch (error ) {
        console.log(error);
      }
    };

    const filteredItems = inventoryItems.filter((item) => 
      item.name.toLowerCase().includes(search.toLowerCase())
    );



  return (
    <div className='p-4 grid lg:grid-cols-2 gap-6'>

      <Card className='shadow-xl border-none'>
          <CardHeader>
            <CardTitle className='text-xl font-bold'>Add Inventory Item</CardTitle>
          </CardHeader>

          <CardContent>
            <form onSubmit={handleSubmit} className='space-y-4'>
              <Input placeholder='Item name' name='name' value={form.name} onChange={handleChange}/>
              <Input placeholder='Quantity' name='quantity' value={form.quantity} onChange={handleChange}/>

              <Input placeholder='Units (e.g., kg, bags' name='units' value={form.units} onChange={handleChange}/>
              <Input placeholder='Reason (optional)' name='reason' value={form.reason} onChange={handleChange} />

              <motion.div whileTap={{scale:0.95}} whileHover={{scale:1.05}}>
                <Button type="submit" className='w-full'>Add Inventory</Button>
              </motion.div>

            </form>
          </CardContent>
      </Card>

      {/* inventory list  */}
      <Card>
        <CardHeader>
          <CardTitle className='text-xl font-bold'>Inventory List</CardTitle>
        </CardHeader>
        <CardContent>
            <Input
              placeholder='Search items...'
              value={search}
              onChange={(e) => setSearch(e.target.value)}
              className='mb-4'
            />


            <div className='max-h-96 overflow-y-auto'>
              <AnimatePresence>
                {
                  filteredItems.map((item,index)=>(
                    <motion.div
                      key={index}
                      initial={{ opacity: 0, y: 10 }}
                  animate={{ opacity: 1, y: 0 }}
                  exit={{ opacity: 0, y: -10 }}
                  className="p-4 mb-2 rounded-lg bg-gray-50 shadow-sm hover:bg-gray-100 transition cursor-pointer"
               
                    >
                        <p>{item.name}</p>
                        <p>{item.quantity} {item.units} - {item.reason || "No reason provided"}</p>
                    </motion.div>
                  ))
                }
              </AnimatePresence>

              {
                filteredItems.length === 0  && (
                  <p className='text-center py-6 text-gray-500'>No items found</p>
                )
              }
            </div>
        </CardContent>
      </Card>
    </div>
  )
}

export default InventoryV2Page