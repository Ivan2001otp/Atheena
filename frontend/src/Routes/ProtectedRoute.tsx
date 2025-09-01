import React, {useState, useEffect, useRef} from 'react'
import {Navigate} from "react-router-dom";
import { getAccessToken, isTokenExpired } from '@/service/util';
import toast from 'react-hot-toast';


interface ProtectedRouteProps {
    children : React.ReactNode;
}

const ProtectedRoute = ({ children }: ProtectedRouteProps) => {
  const [isExpired, setIsExpired] = useState(false);

  useEffect(() => {
    
    const token = getAccessToken();
    if (!token || isTokenExpired(token)) {
      setIsExpired(true);
      toast.dismiss()
      toast.error("Session is Expired. Kindly Login Again");
    }
    
  }, []); // run once when component mounts

  if (isExpired) {
    return <Navigate to="/login" replace />;
  }

  return <>{children}</>;
};


export default ProtectedRoute