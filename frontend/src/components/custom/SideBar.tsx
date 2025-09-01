import { Navigate, NavLink } from "react-router-dom";

import {
  Home,
  Boxes,
  Users,
  ClipboardList,
  BarChart3,
  User,
  LucideConstruction,
  WarehouseIcon,
  LogOut
} from "lucide-react";

import {
  Tooltip,
  TooltipContent,
  TooltipTrigger,
} from "@/components/ui/tooltip"

import { motion } from "framer-motion";
import { Button } from "../ui/button";

import { LogoutAdmin } from "@/service/auth.api";
import { ADMIN_EMAIL, ADMIN_ROLE, type AdminLogoutRequest } from "@/models/auth";
import toast from "react-hot-toast";
import { clearAuth } from "@/service/util";

interface SidebarProps1 {
  collapsed: boolean;
}

export default function Sidebar({ collapsed}: SidebarProps1) {
  const menuItems = [
    { name: "Dashboard", path: "/atheena/dashboard-v2", icon: <Home /> },
    { name: "Inventory", path: "/atheena/inventory", icon: <Boxes /> },
    { name: "Warehouses", path: "/atheena/warehouses", icon: <WarehouseIcon /> },
    {
      name: "Supervisors",
      path: "/atheena/supervisors",
      icon: <Users />,
    },
    {
      name: "Transactions",
      path: "/atheena/transactions",
      icon: <ClipboardList />,
    },
    { name: "Sites", path: "/atheena/sites", icon: <LucideConstruction /> },
    { name: "Reports", path: "/atheena/reports", icon: <BarChart3 /> },
    { name: "Profile", path: "/atheena/profile", icon: <User /> },
  ];


  const handleLogout = async() => {

    const payload : AdminLogoutRequest = {
      email:(localStorage.getItem(ADMIN_EMAIL) as string) ,
      role:(localStorage.getItem(ADMIN_ROLE) as string)
    };

    try {
      const res =  await LogoutAdmin(payload)
      toast.success(res.message);
      clearAuth()

      setTimeout(()=>window.location.href = "/login",1000)
      
    } catch (errro :any){
      toast.error("Something went wrong (logout) !");
    }
  };

  return (
    <motion.div
      animate={{ width: collapsed ? "60px" : "220px" }}
      className="bg-slate-900 text-white h-full shadow-lg flex flex-col transition-all duration-300"
    >
      {!collapsed ? (

        <motion.div
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.8, ease: "easeOut" }}
          className="flex place-items-center"
        >
          <img
            src="/atheena_logo.png"
            alt="atheena-logo"
            width={120}
            height={120}
          />
          <span className="text-xl font-bold">Atheena</span>
        </motion.div>

      ) : (
        <br />
      )}


      <nav className="mt-2 space-y-0.5">
        {menuItems.map((item) => (
          <NavLink
            key={item.name}
            to={item.path}
            className={({ isActive }) =>
              `block px-4 py-3 rounded hover:bg-blue-600 transistion ${
                isActive ? "bg-blue-900 font-bold" : ""
              }`
            }
          >
            {collapsed ? (
              <Tooltip>
                <TooltipTrigger asChild>
                    <div>{item.icon}</div>
                </TooltipTrigger>
                <TooltipContent>
                  <p>{item.name}</p>
                </TooltipContent>
              </Tooltip>
              
            ) : (
              <div className="flex space-x-4">
                <div>{item.icon}</div>

                <div>{item.name}</div>
              </div>
            )}
          </NavLink>
        ))}
      </nav>

      {!collapsed ?
      <motion.div
        initial={{ opacity: 0, x: -20 }}
          animate={{ opacity: 1, x: 0 }}
          transition={{ duration: 1, ease: "easeOut" }}
          className="p-4"
      >
        <Button
        className="bg-white hover:scale-105 text-black hover:border-2 hover:border-white transition-all duration-300 hover:bg-black hover:text-white hover:font-bold "
        size="lg"
        onClick={handleLogout}
        >
         <LogOut size={40}/> Logout
        </Button>
      </motion.div>
      :``}
    </motion.div>
  );
}
