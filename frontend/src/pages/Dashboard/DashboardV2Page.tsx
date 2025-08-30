import React from 'react'
import { useLocation } from 'react-router-dom'

const DashboardV2Page = () => {

  const location = useLocation();
  const adminDetails = location.state?.admin;
  return (
    <div>

       <div>{adminDetails.name}</div>
       <div>{adminDetails.id}</div>
       <div>{adminDetails.email}</div>
       <div>{adminDetails.role}</div>
    </div>
  )
}

export default DashboardV2Page