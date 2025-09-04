import { ADMIN_ID } from '@/models/auth'
import type {  Supervisor } from '@/models/supervisor';
import { FetchallSupervisorByAdminId } from '@/service/auth.api';
import React, { useEffect, useState } from 'react'

const SupervisorPage = () => {
  const [apiResponse, setApiResponse] = useState<Supervisor[]>([]);

  useEffect(() => {
    console.log("supervisor api triggered..!");
    
      const fetchSupervisors = async() => {
        const adminId = localStorage.getItem(ADMIN_ID)!;
        console.log("supervisor-adminid ", adminId)
   
        if (adminId) {

          try {
            const res = await FetchallSupervisorByAdminId(adminId);
            if (res.status) {
              setApiResponse(res.data);
            }
          } catch (error) {
            console.log(error);
          }

        }
    }

    fetchSupervisors()
    
  }, []);

  return (
    <div>
      {
        apiResponse.map((item,index) => (
          <h1 className='text-black' key={index}>{item.name}</h1>
        ))
      }
    </div>
  )
}

export default SupervisorPage