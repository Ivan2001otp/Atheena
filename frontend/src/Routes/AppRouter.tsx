import LoginPage from "@/pages/Auth/LoginPage";
import RegisterPage from "@/pages/Auth/RegisterPage";
import DashboardPage from "@/pages/Dashboard/dashboardPage";
import dashboardPage from "@/pages/Dashboard/dashboardPage";
import NotFound from "@/pages/Others/NotFound";
import { Routes, Route } from "react-router-dom";


export default function AppRoutes() {
    return (
        
            <Routes>
                {/* public routes  */}
                <Route path="/register" element={<RegisterPage/>}/>
                <Route path="/login" element={<LoginPage/>}/>

                {/* protected routes  */}
                <Route path="/dashboard" element={<DashboardPage/>}/>

                {/* catch all  */}
                <Route path="*" element={<NotFound/>}/>
            </Routes>
    )
}