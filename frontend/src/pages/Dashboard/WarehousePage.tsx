import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { ADMIN_ID, type AddWarehouseRequest, type AdminWarehouse } from "@/models/auth";
import { AddNewWarehouse, GetAllWarehouse } from "@/service/auth.api";
import { formatDateTime, formatDateTimeV2 } from "@/service/util";
import { AnimatePresence, motion } from "framer-motion";
import {
  Building2,
  Edit3,
  Save,
  Search,
  Trash2,
  X,
} from "lucide-react";
import React, { useEffect, useMemo, useState } from "react";
import toast from "react-hot-toast";

type Warehouse = {
  user_id: string;
  name: string;
  pin?: string;
  address: string;
  state: string;
  country: string;
  created_at: number;
};

type FormState = {
  name: string;
  pin: string;
  address: string;
  state: string;
  country: string;
};

const emptyForm: FormState = {
  name: "",
  pin: "",
  address: "",
  state: "",
  country: "",
};

const WarehousePage: React.FC = () => {
  const [form, setForm] = useState<FormState>(emptyForm);
  const [warehouseList, setWarehouseList] = useState<AdminWarehouse[]>([]);

  const [editingId, setEditingId] = useState<string | null>(null);
  const [query, setQuery] = useState("");

  const [sortBy, setSortBy] = useState<string>();
  const [errors, setErrors] = useState<
    Partial<Record<keyof FormState, string>>
  >({});

  // Load or persist
  useEffect(() => {
    // fetch the warehouse from db.
    const fetchWarehouses = async() => { 
      
      try{
        const result =  await GetAllWarehouse(localStorage.getItem(ADMIN_ID)!);
        setWarehouseList(result);
      } catch (error) {
        console.log(error)
      }
    };

    fetchWarehouses();
  }, []);

  // useEffect(() => {}, [warehouseList]);

  const validate = (data: FormState) => {
    const next: Partial<Record<keyof FormState, string>> = {};
    if (!data.name.trim()) next.name = "Warehouse name is required";

    if (!data.address.trim()) next.address = "Address is required";

    if (!data.country.trim()) next.country = "Country is required";

    if (!data.state.trim()) next.state = "State is required";

    return next;
  };

  const reset = () => {
    setForm(emptyForm);
    // set editing.
    setEditingId(null);
    setErrors({});
  };

  const onSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    const v = validate(form);

    setErrors(v);
    if (Object.keys(v).length > 0) return;

    console.log("editing id is ", editingId);
    if (editingId) {
      setWarehouseList((prev) =>
        prev.map((w) =>
          w.id === editingId
            ? {
                ...w,
                
                name: form.name.trim(),
                location: form.pin.trim() || "",
                address: form.address.trim(),
                state: form.state.trim(),
                pin:form.pin.trim(),
                country: form.country.trim(),
                created_at: formatDateTimeV2( Date.now())
              }
            : w
        )
      );


    
      
    const payload: AddWarehouseRequest = {
        id: editingId,
        user_id: localStorage.getItem(ADMIN_ID)!,
        name: form.name.trim(),
        address: form.address.trim(),
        country: form.country.trim(),
        state: form.state.trim(),
        pin: form.pin.trim()
      };

      try {
        await AddNewWarehouse(payload);
        console.log("Successfully updated the warehouse.");
        toast.success("Updated successfully");
        // setWarehouseList((prev) => [nw, ...prev]);
      } catch (error) {
        toast.error("something went wrong while updating warehouse !");
      }

    } else {
      const nw: AdminWarehouse = {
        id:'',// testing
        name: form.name.trim(),
        address: form.address.trim(),
        country: form.country.trim(),
        state: form.state.trim(),
        pin: form.pin.trim(),
        user_id: localStorage.getItem(ADMIN_ID)! ,
        created_at: formatDateTimeV2( Date.now())
      };

      // save the warehouse to database.
      const payload: AddWarehouseRequest = {
        id:"",
        user_id: localStorage.getItem(ADMIN_ID)!,
        name: form.name.trim(),
        address: form.address.trim(),
        country: form.country.trim(),
        state: form.state.trim(),
        pin: form.pin.trim(),
      };

      try {
        const res = await AddNewWarehouse(payload);
        console.log("Successfully saved the warehouse.");
        toast.success(res.message);
        setWarehouseList((prev) => [nw, ...prev]);
      } catch (error) {
        toast.error("something went wrong !");
      }
    }

    reset();
  };

  const onEdit = (id: string) => {
    const w = warehouseList.find((item) => item.id === id);
    if (!w) return;

    setEditingId(id); 
    setForm({
      name: w.name,
      address: w.address,
      state: w.state,
      country: w.country,
      pin: w.pin ,
    });

  };

  const onDelete = (name: string) => {
    // need to work

    if (!confirm("Delete this warehouse ?")) return;

   // setWarehouseList((prev) => prev.filter((i) => i.name !== name));
   // if (editingId === name) reset();
  };

  const filtered = useMemo(() => {
    const q = query.toLowerCase().trim();
    let arr = warehouseList.filter((w) =>
      [w.name, w.pin ?? "", w.address, w.state, w.country]
        .join("")
        .toLowerCase()
        .includes(q)
    );

    if (sortBy === "recent") {
      arr = arr.sort((a, b) => new Date(b.created_at).getTime() - new Date(a.created_at).getTime());
    } else if (sortBy === "name") {
      arr = arr.sort((a, b) => a.name.localeCompare(b.name));
    } else if (sortBy === "state") {
      arr = arr.sort((a, b) => a.state.localeCompare(b.state));
    } else if (sortBy === "country") {
      arr = arr.sort((a, b) => a.country.localeCompare(b.country));
    }

    return arr;
  }, [warehouseList, query, sortBy]);

  return (
    <div className="min-h-screen w-full bg-gradient-to-br from-slate-50 via-white to-slate-100">
      <header className="sticky top-0 z-20 backdrop-blur bg-white/70 border-b">
        <div className="mx-auto max-w-7xl px-4 py-4 flex items-center justify-between">
          <div className="flex items-center gap-3">
            <div className="p-2 rounded-xl bg-indigo-600 text-white shadow-sm">
              <Building2 size={20} />
            </div>

            <h1 className="text-xl font-bold text-slate-800">Warehouses</h1>
          </div>

          <div className="flex items-center gap-3 w-full max-w-md">
            <div className="relative flex-1">
              <Search
                className="absolute left-3 top-1/2 -translate-y-1/2"
                size={18}
              />
              <input
                value={query}
                onChange={(e) => setQuery(e.target.value)}
                placeholder="Search by name, country, state..."
                className="w-full pl-9 pr-3 py-2 rounded-lg border bg-white/80 focus:bg-white outline-none focus:ring-2 ring-indigo-200"
              />
            </div>

            <div className="relative">
              {/*
                <select
                  value={sortBy}
                  onChange={e => setSortBy(e.target.value as any)}
                  className = "appearance-none pr-9 pl-3 py-2 rounded-lg border bg-white/80 focus:bg-white outline-none focus:ring-2 ring-indigo-200"
                >
                  <option value="recent">Sort: Recent</option>
                  <option value="name">Sort: Name</option>
                  <option value="state">Sort: State</option>
                  <option value="country">Sort: Country</option>

                </select>
              */}
              <Select onValueChange={(value) => setSortBy(value)}>
                <SelectTrigger id="sortBy">
                  <SelectValue placeholder="Sort By.." />
                </SelectTrigger>

                <SelectContent>
                  <SelectItem value="recent">Recent</SelectItem>
                  <SelectItem value="name">Warehouse-name</SelectItem>
                  <SelectItem value="state">State</SelectItem>
                  <SelectItem value="country">Country</SelectItem>
                </SelectContent>
              </Select>

              {/* <ChevronDown
                  size={16}
                  className='pointer-events-none absolute right-2 top-1/2 -translate-y-1/2 text-slate-500'
                /> */}
            </div>
          </div>
        </div>
      </header>

      <main className="mx-auto max-w-7xl px-4 py-8 grid grid-cols-1 lg:grid-cols-5 gap-8">
        <motion.section
          initial={{ opacity: 0, y: 16 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.35 }}
          className="lg:col-span-2"
        >
          <div className="rounded-2xl border bg-white shadow-sm p-6">
            <h2>{editingId ? "Edit Warehouse" : "Add Warehouse"}</h2>

            <form onSubmit={onSubmit} className="space-y-4">
              {/* warehouse name  */}
              <div>
                <label className="text-sm font-medium text-slate-700">
                  Warehouse Name *
                </label>
                <input
                  required
                  maxLength={256}
                  value={form.name}
                  onChange={(e) => setForm({ ...form, name: e.target.value })}
                  placeholder="e.g., North Hub"
                  className="mt-1 w-full rounded-lg border px-3 py-2 outline-none focus:ring-2 focus:ring-indigo-200"
                />
                {errors.name && (
                  <p className="text-xs text-red-600 mt-1">{errors.name}</p>
                )}
              </div>

              {/* address  */}
              <div>
                <label className="text-sm font-medium text-slate-700">
                  Address
                </label>
                <textarea
                  maxLength={256}
                  required
                  value={form.address}
                  onChange={(e) =>
                    setForm({ ...form, address: e.target.value })
                  }
                  rows={4}
                  className="mt-1 w-full rounded-lg border px-3 py-2 outline-none focus:ring-2 ring-indigo-200 resize-y"
                  placeholder="Street, Area, ZIP"
                />
                {errors.address && (
                  <p className="text-xs text-red-600 mt-1">{errors.address}</p>
                )}
              </div>

              {/* pincode  */}
              <div>
                <label className="text-sm font-medium text-slate-700">
                  Pin
                </label>
                <input
                  value={form.pin}
                  maxLength={8}
                  required
                  onChange={(e) => setForm({ ...form, pin: e.target.value })}
                  placeholder="e.g.,189001 or KAD313"
                  className="mt-1 w-full rounded-lg border px-3 py-2 outline-none focus:ring-2 ring-indigo-200"
                />
                {errors.pin && (
                  <p className="text-xs text-red-600 mt-1">{errors.pin}</p>
                )}
              </div>

              {/* state  */}
              <div>
                <label className="text-sm font-medium text-slate-700">
                  State
                </label>
                <input
                  maxLength={30}
                  value={form.state}
                  onChange={(e) => setForm({ ...form, state: e.target.value })}
                  placeholder="e.g.,Goa or Denver"
                  className="mt-1 w-full rounded-lg border px-3 py-2 outline-none focus:ring-2 ring-indigo-200"
                />
                {errors.state && (
                  <p className="text-xs text-red-600 mt-1">{errors.state}</p>
                )}
              </div>

              {/* country  */}
              <div>
                <label className="text-sm font-medium text-slate-700">
                  Country
                </label>
                <input
                  maxLength={30}
                  value={form.country}
                  onChange={(e) =>
                    setForm({ ...form, country: e.target.value })
                  }
                  placeholder="e.g.,189001 or KAD313"
                  className="mt-1 w-full rounded-lg border px-3 py-2 outline-none focus:ring-2 ring-indigo-200"
                />
                {errors.state && (
                  <p className="text-xs text-red-600 mt-1">{errors.state}</p>
                )}
              </div>

              {/* buttons  */}
              <div className="flex items-center gap-3 pt-2">
                <button
                  type="submit"
                  className="inline-flex items-center gap-2 rounded-lg bg-indigo-600 text-white px-4 py-2 font-medium shadow hover:bg-indigo-700 active:scale-[.99] transition"
                >
                  {editingId ? <Save size={18} /> : <Save size={18} />}
                  {editingId ? "Save Changes" : "Add Warehouse"}
                </button>

                {editingId && (
                  <button
                    type="button"
                    onClick={reset}
                    className="inline-flex items-center gap-2 rounded-lg border px-4 py-2 font-medium hover:bg-slate-50 transition"
                  >
                    <X size={18} />
                    Cancel
                  </button>
                )}
              </div>
            </form>
          </div>
        </motion.section>

        <section className="lg:col-span-3">
          <div className="grid grid-cols-1 sm:grid-cols-2 xl:grid-cols-3 gap-6">
            <AnimatePresence>
              {filtered.map((w) => (
                <motion.div
                  key={w.id}
                  
                  layout
                  initial={{ opacity: 0, y: 16, scale: 0.98 }}
                  animate={{ opacity: 1, y: 0, scale: 1 }}
                  exit={{ opacity: 0, y: 8 }}
                  transition={{ duration: 0.6 }}
                  className="group rounded-2xl border bg-white shadow-sm hover:shadow-md transition"
                >
                  <div className="relative p-5 h-full flex flex-col">
                    <div className="flex items-start justify-between gap-3">
                      <div className="flex items-center gap-2 ">
                        <div className="p-2 rounded-xl bg-indigo-50 text-indigo-600">
                          <Building2 size={20} />
                        </div>
                        <div>
                          <h3 className="font-semibold text-slate-900 leading-tight">
                            {w.name}
                          </h3>
                          <p className="text-xs text-slate-500 mt-2">
                            Added {formatDateTime(w.created_at)}
                          </p>
                        </div>

                        <div className="absolute z-20 bottom-0 right-0 gap-1 opacity-0 group-hover:opacity-100 transition ">
                        <button
                          onClick={() => onEdit(w.id)}
                          className="p-2 rounded-lg hover:bg-slate-100"
                          aria-label="Edit"
                        >
                          <Edit3 size={18} />
                        </button>
                        <button
                          onClick={() => onDelete(w.name)}
                          className="p-2 rounded-lg hover:bg-red-50 text-red-600"
                          aria-label="Delete"
                        >
                          <Trash2 size={18} />
                        </button>
                      </div>
                      </div>

                      
                    </div>
                  </div>
                </motion.div>
              ))}
            </AnimatePresence>
          </div>

          {filtered.length === 0 && (
            <motion.div
              initial={{ opacity: 0, y: 6 }}
              animate={{ opacity: 1, y: 0 }}
              className="text-center text-slate-500 mt-10"
            >
              No warehouses yet. Add your first one on the left!
            </motion.div>
          )}
        </section>
      </main>
    </div>
  );
};

export default WarehousePage;
