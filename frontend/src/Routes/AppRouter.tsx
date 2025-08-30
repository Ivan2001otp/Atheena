import LoginPage from "@/pages/Auth/LoginPage";
import RegisterPage from "@/pages/Auth/RegisterPage";
import DashboardPage from "@/pages/Dashboard/HomePage";

import NotFound from "@/pages/Others/NotFound";
import { Routes, Route } from "react-router-dom";
import ProtectedRoute from "./ProtectedRoute";
import AccessDeniedPage from "@/pages/Auth/AccessDeniedPage";
import DashboardV2Page from "@/pages/Dashboard/DashboardV2Page";
import InventoryPage from "@/pages/Dashboard/InventoryPage";
import WarehousePage from "@/pages/Dashboard/WarehousePage";
import SupervisorPage from "@/pages/Dashboard/SupervisorPage";
import ProfilePage from "@/pages/Dashboard/ProfilePage";
import ReportPage from "@/pages/Dashboard/ReportPage";
import SitesPage from "@/pages/Dashboard/SitesPage";
import TransactionPage from "@/pages/Dashboard/TransactionPage";

export default function AppRoutes() {
  return (
    <Routes>
      {/* public routes  */}
      <Route path="/register" element={<RegisterPage />} />
      <Route path="/login" element={<LoginPage />} />
      <Route path="/access_denied" element={<AccessDeniedPage />} />

      {/* protected routes  */}
      <Route
        path="/dashboard-v1"
        element={
          <ProtectedRoute>
            <DashboardPage />
          </ProtectedRoute>
        }
      >
        <Route
          path="dashboard-v2"
          element={
            <ProtectedRoute>
              <DashboardV2Page />
            </ProtectedRoute>
          }
        />

        <Route
          path="inventory"
          element={
            <ProtectedRoute>
              <InventoryPage />
            </ProtectedRoute>
          }
        />

        <Route
          path="warehouses"
          element={
            <ProtectedRoute>
              <WarehousePage />
            </ProtectedRoute>
          }
        />

        <Route
          path="supervisors"
          element={
            <ProtectedRoute>
              <SupervisorPage />
            </ProtectedRoute>
          }
        />
        <Route
          path="transactions"
          element={
            <ProtectedRoute>
              <TransactionPage />
            </ProtectedRoute>
          }
        />

        <Route
          path="sites"
          element={
            <ProtectedRoute>
              <SitesPage />
            </ProtectedRoute>
          }
        />

        <Route
          path="reports"
          element={
            <ProtectedRoute>
              <ReportPage />
            </ProtectedRoute>
          }
        />

        <Route
          path="profile"
          element={
            <ProtectedRoute>
              <ProfilePage />
            </ProtectedRoute>
          }
        />
      </Route>

      {/* catch all  */}
      <Route path="*" element={<NotFound />} />
    </Routes>
  );
}
