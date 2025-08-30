import { useState } from "react";
import { motion } from "framer-motion";
import JsCookies from "js-cookie";
import { useNavigate } from "react-router-dom";

import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Button } from "@/components/ui/button";

import toast from "react-hot-toast";
import { LoginAdmin } from "@/service/auth.api";
import { ACCESS_TOKEN, REFRESH_TOKEN } from "@/models/auth";

const LoginPage = () => {
  const [loading, setIsLoading] = useState(false);
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [formData, setFormData] = useState({
    email: "",
    password: "",
  });

  const navigate = useNavigate();
  const validate = (): boolean => {
    let isValid = true;

    if (formData.email.trim().length === 0) {
      isValid = false;
    }

    if (formData.password.trim().length === 0) {
      isValid = false;
    }

    return isValid;
  };


  const handleOnChange = (key:string, value:string) => {
    setFormData({...formData, [key]:value});
  };

  const handleLogin = async(e: React.FormEvent) => {
    e.preventDefault();
    setIsLoading(true);

    if (!validate()) {
      toast.error("Kindly provide your input to missing fields");
      return;
    }

    async function hitAdminLoginApi() {
      try {
        const payload = {
          email: formData.email,
          password: formData.password,
        };

        const res = await LoginAdmin(payload);
        console.log(res);

        // 15mins
        JsCookies.set(ACCESS_TOKEN, res.access_token, {
          expires: 0.0104,
          secure: true,
          sameSite: "Strict",
        });

        // 1 day
        JsCookies.set(REFRESH_TOKEN, res.refresh_token, {
          expires: 1,
          secure: true,
          sameSite: "Strict",
        });

        console.log("login success");

        const admin_payload = {
          name: res.name,
          email: res.email,
          id: res.id,
          role: res.role,
        };

        toast.success("Welcome")

        setTimeout(()=>navigate("/dashboard-v1/dashboard-v2", { state: { "admin": admin_payload } }), 1500);
        //  navigate("/dashboard-v1", { state: { "admin": admin_payload } })

    } catch (error: any) {
        
        console.log("login failure");
        console.log(error);

      } finally {
        setIsLoading(false);
      }
    }

    await toast.promise(
      hitAdminLoginApi,
      {
        loading:"Loading...",
      }
    )
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-purple-600 via-blue-500 to-indigo-700 p-4">
      <motion.div
        initial={{ opacity: 0, y: 40 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ duration: 0.8, ease: "easeOut" }}
        className="w-full max-w-md"
      >
        <Card className="rounded-2xl shadow-2xl backdrop-blur-lg bg-white/80">
          <CardHeader>
            <CardTitle className="text-center text-2xl font-bold text-gray-800">
              Welcome Back 👋
            </CardTitle>
          </CardHeader>

          <CardContent>
            <form onSubmit={handleLogin} className="space-y-4">
              {/* Email  */}
              <div className="space-y-2">
                <Label htmlFor="email" className="text-gray-700">
                  Email
                </Label>
                <Input
                  id="email"
                  type="email"
                  value={formData.email}
                  onChange={(e)=>handleOnChange("email", e.target.value)}
                  required
                  className="mt-1 focus:ring-2 focus:ring-blue-400"
                  placeholder="johndoe.@example.com"
                />
              </div>

              <div className="space-y-2">
                <Label htmlFor="password" className="text-gray-700">
                  Password
                </Label>
                <Input
                  id="password"
                  type="password"
                  onChange={(e)=>handleOnChange("password", e.target.value)}
                  value={formData.password}
                  required
                  className="mt-1 focus:ring-2 focus:ring-blue-400"
                  placeholder="••••••••"
                />
              </div>

              <motion.div
                whileHover={{ scale: 1.02 }}
                whileTap={{ scale: 0.98 }}
              >
                <Button
                  type="submit"
                  className="w-full bg-blue-600 hover:bg-blue-700 text-white"
                  disabled={loading}
                >
                  {loading ? "Logging in..." : "Login"}
                </Button>
              </motion.div>
            </form>

            <p className="mt-4 text-center text-sm text-gray-600">
              Don’t have an account?{" "}
              <a
                href="/register"
                className="text-blue-600 hover:underline font-medium"
              >
                Register
              </a>
            </p>
          </CardContent>
        </Card>
      </motion.div>
    </div>
  );
};

export default LoginPage;
