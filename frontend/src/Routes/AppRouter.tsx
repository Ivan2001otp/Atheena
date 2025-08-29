import LoginPage from "@/pages/Auth/LoginPage";
import RegisterPage from "@/pages/Auth/RegisterPage";
import DashboardPage from "@/pages/Dashboard/dashboardPage";
import NotFound from "@/pages/Others/NotFound";
import { Routes, Route } from "react-router-dom";
import ProtectedRoute from "./ProtectedRoute";
import AccessDeniedPage from "@/pages/Auth/AccessDeniedPage";


export default function AppRoutes() {
    return (
        
            <Routes>
                {/* public routes  */}
                <Route path="/register" element={<RegisterPage/>}/>
                <Route path="/login" element={<LoginPage/>}/>
                <Route path="/access_denied" element={<AccessDeniedPage/>}/>

                {/* protected routes  */}
                <Route path="/dashboard" element={
                    <ProtectedRoute>
                          <DashboardPage/>
                    </ProtectedRoute>
                    }/>

                {/* catch all  */}
                <Route path="*" element={<NotFound/>}/>
            </Routes>
    )
}