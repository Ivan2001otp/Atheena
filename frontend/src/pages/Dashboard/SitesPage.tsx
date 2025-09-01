import React, { useEffect, useState } from "react";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle, DialogTrigger } from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { motion } from "framer-motion";
import { Trash2, PlusCircle } from "lucide-react";
import { ADMIN_ID, type SiteModel } from "@/models/auth";
import { formatDateTimeV2 } from "@/service/util";
import { AddNewConstructionSite, GetAllConstructionSites } from "@/service/auth.api";
import toast from "react-hot-toast";




const SitesPage = () => {

  const [constructionSiteList, setConstructionSites] = useState<SiteModel[]>([]);
  const [formData, setForm] = useState({
    name:"",
    address:"",
    country:"",
    state:"",
  });
  
  const [deleteSiteId, setDeleteSiteId] = useState<string | null>(null);

  // Load or persist
    useEffect(() => {
      // fetch the warehouse from db.
      const fetchWarehouses = async() => { 
        
        try{
          if(localStorage.getItem(ADMIN_ID)){
            const response = await GetAllConstructionSites(localStorage.getItem(ADMIN_ID)!);
            setConstructionSites(response);
          }
        } catch (error) {
          console.log(error)
        }
      };
  
      fetchWarehouses();
    }, []);
  

  const addSite = async() => {
    
    if (formData.address.trim().length === 0) return;
    if (formData.country.trim().length === 0)return;
    if (formData.name.trim().length === 0) return;
    if (formData.state.trim().length === 0 ) return;

    let item : SiteModel = {
      id:Date.now().toString(),
      user_id: localStorage.getItem(ADMIN_ID)!,
      address: formData.address,
      country: formData.country,
      name: formData.name,
      state: formData.state,
      updated_time : formatDateTimeV2(Date.now())
    }

    setConstructionSites([...constructionSiteList, item ]);
    setForm({
      name:"",
      address:"",
      country:"",
      state:"",
    })

    try {
      const res = await AddNewConstructionSite(item)
      toast.success(res.message);
    } catch (error) {
      console.log("something went wrong while adding new construction site.");
      console.log(error);
    }
  };

  const confirmDeleteSite = (id: string) => {
    setDeleteSiteId(id);
    deleteSite() 
  }

  const deleteSite = () => {
    if (deleteSiteId !== null) {
      setConstructionSites(constructionSiteList.filter((item)=> item.id !== deleteSiteId.toString()));
      setDeleteSiteId(null);
    }
  };

  const handleOnChange = (key:string, value:string) => {
    setForm({...formData, [key]:value});
  };


  return (
    <div className="min-h-screen bg-gray-50 p-6">

      <div className="flex justify-between items-center mb-8">
        <h1 className="text-3xl font-bold text-gray-800">
          🏗️Construction Sites
        </h1>
        <Dialog>
          <DialogTrigger asChild>
            <Button className="flex items-center gap-2">
              <PlusCircle className="w-5 h-5"/> Add Site
            </Button>
          </DialogTrigger>
          <DialogContent>
            <DialogHeader>
              <DialogTitle>Add New Site</DialogTitle>
              <DialogDescription>Enter the site name to add it your list.</DialogDescription>
            </DialogHeader>



            <Input
              placeholder="Site name..."
              value={formData.name}
              onChange={(e)=>handleOnChange("name", e.target.value)}
              required
            />

            <Input
              placeholder="Site address..."
              value={formData.address}
              onChange={(e)=>handleOnChange("address", e.target.value)}
              required
            />

            <Input
              placeholder="State..."
              value={formData.state}
              onChange={(e)=>handleOnChange("state", e.target.value)}
              required
            />

            <Input
              placeholder="Country..."
              value={formData.country}
              onChange={(e)=>handleOnChange("country", e.target.value)}
              required
            />
            <DialogFooter>
              <Button className="cursor-pointer" onClick={addSite}>Add Site</Button>
            </DialogFooter>
          </DialogContent>
        </Dialog>
      </div>




      <p className="mb-4 text-gray-600 font-semibold text-  xl">Sites contract : {constructionSiteList.length}</p>
      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
        {
          constructionSiteList.map((site, index) => (
            <motion.div
              key={site.id}
              initial={{opacity:0, y:20}}
              animate={{opacity:1, y:0}}
              transition={{delay:index* 0.1}}
            >
                <Card className="shadow-lg hover:shadow-xl transition rounded-2xl border">
                  <CardHeader className="flex justify-between items-center">
                    <CardTitle className="text-lg font-medium">{site.name}</CardTitle>
                    <Button
                      variant="destructive"
                      size="icon"
                      onClick={()=>confirmDeleteSite(site.id)}
                    >
                      <Trash2 className="w-4 h-4"/>
                    </Button>
                  </CardHeader>

                  <CardContent>
                    <p>Monitoring site Materials...</p>
                  </CardContent>
                </Card>
            </motion.div>
          ))
        }
      </div>
    </div>
  )
}

export default SitesPage