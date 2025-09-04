import React, { useEffect, useMemo, useState } from "react";
import { motion, AnimatePresence } from "framer-motion";
import { PlusCircle, Trash2, Edit3, Search } from "lucide-react";
import toast, { Toaster } from "react-hot-toast";

/* Replace these imports with your shadcn/ui path */
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogDescription,
  DialogFooter,
  DialogTrigger,
} from "@/components/ui/dialog";
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogDescription,
  AlertDialogTrigger,
} from "@/components/ui/alert-dialog";
import { Label } from "@/components/ui/label";



import { ADMIN_ID } from '@/models/auth'
import type {  Supervisor } from '@/models/supervisor';
import { FetchallSupervisorByAdminId, UpsertSupervisor } from '@/service/auth.api';

const SupervisorPage = () => {
  const [supervisorList, setSupervisorList] = useState<Supervisor[]>([]);
  const [query, setQuery] = useState("");
  const [sortBy, setSortBy] = useState<"name">("name");

  const [openAdd, setOpenAdd] = useState(false);
  const [editing, setEditing] = useState<Supervisor | null>(null);


  // Delete state
  const [deleteId, setDeleteId] = useState<string | null>(null);
  const [confirmOpen, setConfirmOpen] = useState(false);

  const [form, setForm] = useState({
    name: "",
    email: "",
    phone: "",
    role:""
  });

  const [errors, setErrors] = useState<Partial<Record<keyof typeof form, string>>>({});

  useEffect(() => {
    console.log("supervisor api triggered..!");
    
      const fetchSupervisors = async() => {
        const adminId = localStorage.getItem(ADMIN_ID)!;
        console.log("supervisor-adminid ", adminId)
   
        if (adminId) {

          try {
            const res = await FetchallSupervisorByAdminId(adminId);
            if (res.success) {
              setSupervisorList(res.data);
            }
          } catch (error) {
            console.log(error);
          }

        }
    }

    fetchSupervisors()
    
  }, []);

 const visible = useMemo(() => {
  const q = query.trim().toLowerCase();
  let arr = supervisorList.filter((supervisor) => 
    [supervisor.name, supervisor.email, supervisor.role, supervisor.phone_number || ""].join(" ").toLowerCase().includes(q)
  );

  arr = arr.sort((a,b) => a.name.localeCompare(b.name));
  return arr;
 }, [supervisorList, query, sortBy]);

 const resetForm = () => {
    setForm({name:"", email:"", phone:"", role:""});
    setErrors({});
    setEditing(null);
 };

 const openAddDialog = () => {
    resetForm();
    setOpenAdd(true);
 }

 const startEdit = (s: Supervisor) => {
    setEditing(s);
    setForm({ name: s.name, email: s.email, phone: s.phone_number || "", role: s.role });
    setOpenAdd(true);
  };

/* validation */
const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
 const validate = () => {
    const next: Partial<Record<keyof typeof form, string>> = {};
    if (!form.name.trim()) next.name = "Name is required";
    if (!form.email.trim() || !emailRegex.test(form.email)) next.email = "Valid email required";
    if (!form.role.trim()) next.role = "Role required";
    return next;
  };

  const handleSave = async() => {
    const v = validate();
    setErrors(v);

    if (Object.keys(v).length) return;

    
          if (editing) {
            const updatedSupervisor = {
                  ...editing,
                  name: form.name,
                  email: form.email,
                  phone_number: form.phone,
                  role: form.role,
                };
        
        console.log("Editing supervisor id : ", editing.id)

        for(let i = 0;i < supervisorList.length; i++) {
            if (supervisorList[i].id === editing.id) {
              console.log("Matched");
             
              supervisorList[i].name = form.name;
              supervisorList[i].email = form.email;
              supervisorList[i].phone_number = form.phone;
              supervisorList[i].role = form.role;

            }
        }
        

    
      try {
          
        const res =  await UpsertSupervisor(updatedSupervisor);

          if (res.success) {
            toast.success("Supervisor updated.");
          }
      } catch (error) {
        console.log(error);
      }
    

    } else {

      // api to update/add supervisor

      const s : Supervisor = {
        id:new Date().toISOString() + "w12",
        admin_id : localStorage.getItem(ADMIN_ID)!,
        name: form.name,
        email : form.email,
        phone_number: form.phone,
        role : form.role,
      };

      setSupervisorList((prev) => [s, ...prev]);

      try {
          s.id = ""; // just to make sure it won't break.
          const res =  await UpsertSupervisor(s);
          if (res.success) {
            toast.success("Supervisor added.");
          }
      } catch (error) {
        console.log(error);
      }


    }

    setOpenAdd(false);
    resetForm();

  };

  const confirmDelete = (id:string) => {
    setDeleteId(id);
    setConfirmOpen(true);
  }

  const doDelete =()=> {
    if (!deleteId) return;
    setSupervisorList((prev) => prev.filter((s) => s.id !== deleteId));
    setDeleteId(null);
    setConfirmOpen(false);
    toast.success("Supervisor deleted");
  };

  return (
    <div
      className="min-h-screen p-6 bg-gradient-to-b from-slate-50 to-white"
    >

        <div
          className="max-w-7xl mx-auto"
        >

          <div className="flex items-center justify-between gap-4 mb-6">
              <div className="space-y-4">
                <h1 className="text-2xl lg:text-3xl font-extrabold">Supervisors</h1>

                <p className="text-sm text-slate-500">Manage your employee</p>
              </div>


              <div className="flex items-center gap-3">
                <div className="flex items-center bg-white border rounded-lg px-3 py-1 shadow-sm">
                  <Search size={16} className="text-slate-400 mr-2"/>
                  <input
                    placeholder="Search supervisors..."
                    value={query}
                    onChange={(e) => setQuery(e.target.value)}
                    className="outline-none px-2 text-sm w-64"
                  />
                </div>


                <div className="flex items-center gap-2">
                  <select
                    value={sortBy}
                    onChange={(e) => setSortBy(e.target.value as any)}
                    className="rounded-lg border px-3 py-2 bg-white"
                  >
                    <option value="name">Name</option>
                  </select>

                  <Button onClick={openAddDialog} className="inline-flex items-center gap-2">
                    <PlusCircle className="w-4 h-4"/> Add
                  </Button>
                </div>
              </div>
          </div>

          <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
            <motion.div
              initial={{opacity:0, y:6}}
              animate={{opacity:1, y:0}}
              className="lg:col-span-1"
            >
                <Card className="p-4">
                  <CardHeader>
                    <CardTitle className="text-lg">Overview</CardTitle>
                  </CardHeader>

                  <CardContent>
                    <div className="space-y-3">
                      <div className="flex items-center justify-between">
                        <div className="text-sm text-slate-600">Total </div>
                        <div className="text-xl font-semibold">{supervisorList.length}</div>
                      </div>


                      <div className="text-sm text-slate-500">
                        Recent:{" "}
                        <span>
                          {supervisorList.slice(0,3).map((s)=>s.name).join(", ") || "-"}
                        </span>
                      </div>
                    </div>
                  </CardContent>
                </Card>
            </motion.div>


            {/* Supervisor List  */}
            <div className="lg:col-span-2">
              <Card className="p-0 overflow-hidden">
                <div className="p-4 border-b flex items-center justify-between bg-white">
                  <div className="flex items-center gap-3">
                    <h3 className="text-lg font-semibold">Team</h3>
                    <p className="text-sm text-slate-500">{visible.length} shown</p>
                  </div>
                </div>


                <div className="p-4">
                  <div className="grid gap-3">
                    <AnimatePresence>
                      {
                        visible.map((s,index) => (

                          <motion.div
                            key={index}
                            layout
                            initial={{ opacity: 0, y: 8 }}
                            animate={{ opacity: 1, y: 0 }}
                            exit={{ opacity: 0, y: -6 }}
                            transition={{ duration: 0.18 }}
                            className="group bg-white border rounded-lg p-4 flex items-start justify-between shadow-sm hover:shadow-md"
                          >

                            <div>
                              <div className="flex items-center gap-3"> 
                                  <div className="w-10 h-10 rounded-full bg-indigo-50 flex justify-center text-indigo-600 font-semibold items-center">
                                    {s.name.split(" ").map((n) => n[0]).slice(0, 2).join("")}
                                  </div>

                                  <div>
                                    <div className="font-semibold">{s.name}</div>
                                    <div className="text-sm text-slate-500">{s.email}</div>
                                  </div>
                              </div>


                              <div className="mt-2 text-sm text-slate-600 flex gap-4">
                                  <div>{s.role}</div>
                                  {
                                    s.phone_number && <div>• {s.phone_number}</div>
                                  }
                              </div>
                            </div>


                            <div className="flex items-center gap-2 opacity-0 group-hover:opacity-100 transition">
                                  <button className="p-2 rounded-lg hover:bg-slate-100"
                                    onClick={()=> startEdit(s)}
                                    aria-label="Edit"
                                  >
                                    <Edit3 className="w-4 h-4 text-slate-700"/>
                                  </button>


                                  <button
                                    onClick={()=>confirmDelete(s.id)}
                                    className="p-2 rounded-lg hover:bg-rose-50"
                                    aria-label="Delete"
                                  >
                                    <Trash2 className="w-4 h-4 text-rose-600"/>
                                  </button>
                            </div>

                          </motion.div>
                        ))
                      }
                    </AnimatePresence>

                    {
                      visible.length === 0 && (
                        <div className="py-12 text-center text-slate-500">No supervisors found</div>
                      )
                    }
                  </div>
                </div>
              </Card>
            </div>
          </div>

        </div>


        <Dialog open={openAdd} onOpenChange={(o) => {setOpenAdd(o) ;

            if (!o) resetForm();
        }}>

          <DialogContent
            aria-describedby="something"
          >
            <DialogHeader>
              <DialogTitle >
                {editing ? "Editing Supervisor" : "Add Supervisor"}
              </DialogTitle>

              <div className="text-sm text-slate-500 mt-1">
                {editing ? "Update supervisor detail" : "Add a new supervisor to your team"}
              </div>
            </DialogHeader>



            <div className="space-y-4 mt-4">
              <div className="space-y-2">
                <Label>Name</Label>
                <Input
                  value={form.name}
                  onChange={(e) => setForm({...form, name: e.target.value})}
                  placeholder="Full name"
                />
                {
                  errors.name && <div className="text-rose-600 text-sm mt-1">{errors.name}</div>
                }
              </div>


              <div className="space-y-2">
                <Label>Email</Label>
                 <Input value={form.email} onChange={(e) => setForm({ ...form, email: e.target.value })} placeholder="email@example.com" />
              {errors.email && <div className="text-rose-600 text-sm mt-1">{errors.email}</div>}
            
              </div>


                <div className="space-y-2">
                <Label>Phone (optional)</Label>
                <Input value={form.phone} onChange={(e) => setForm({ ...form, phone: e.target.value })} placeholder="+1 555 555 5555" />
              </div>

              <div className="space-y-2">
                <Label>Role</Label>
                <Input value={form.role} onChange={(e) => setForm({ ...form, role: e.target.value })} placeholder="Supervisor / Site Lead" />
                {errors.role && <div className="text-rose-600 text-sm mt-1">{errors.role}</div>}
              </div>

            </div>

            <DialogFooter>
                <div className="flex gap-2 justify-end w-full">
                  <Button variant="ghost" onClick={()=>{
                    setOpenAdd(false);
                    resetForm();
                  }}>Cancel</Button>


                  <Button onClick={handleSave}>
                    {editing ? "Update" : "Add supervisor"}
                  </Button>
                </div>
            </DialogFooter>
          </DialogContent>

        </Dialog>


        <AlertDialog open={confirmOpen} onOpenChange={(o) => { setConfirmOpen(o); if (!o) setDeleteId(null); }}>
        <AlertDialogContent>
          <AlertDialogHeader>
            <AlertDialogTitle>Delete Supervisor</AlertDialogTitle>
            <AlertDialogDescription>
              Are you sure you want to permanently remove this supervisor? This action cannot be undone.
            </AlertDialogDescription>
          </AlertDialogHeader>

          <div className="flex gap-2 justify-end mt-4">
            <AlertDialogCancel onClick={() => { setConfirmOpen(false); setDeleteId(null); }}>Cancel</AlertDialogCancel>
            <AlertDialogAction onClick={doDelete}>Delete</AlertDialogAction>
          </div>
        </AlertDialogContent>
      </AlertDialog>
    </div>
  )
}

export default SupervisorPage