import React from "react";

import { useState } from "react";
import {  Outlet, useLocation } from "react-router-dom";

import SideBar from "@/components/custom/SideBar";

const HomePage = () => {

 const [collapsed, setCollapsed] = useState(true);

  return (
    <div className="flex h-screen w-full">
      <SideBar collapsed={collapsed}/>

      <div className="flex flex-col flex-1 transition-all duration-300">
        <div className="h-14 flex items-center px-4 shadow-md bg-white">
          <button
            onClick={()=>setCollapsed(!collapsed)}
             className="px-3 py-1 rounded bg-blue-600 text-white hover:bg-blue-700 transition"

          >
            {collapsed ? "☰" : "⮜"}
          </button>
          <h1 className="ml-4 text-lg font-semibold">Atheena</h1>
        </div>


        <div className="p-4 overflow-y-auto flex-1 bg-gray-50">
          <Outlet/>
        </div>

      </div>
     
    </div>
  );
};

export default HomePage;
