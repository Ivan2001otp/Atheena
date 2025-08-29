import { useState } from "react";
import { motion } from "framer-motion";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Button } from "@/components/ui/button";


const LoginPage = () => {
  const [loading, setLoading] = useState(false)
  const [formData, setFormData] = useState({
    email:"",
    password:""
  });

  const handleLogin = (e:React.FormEvent) => {
    e.preventDefault()
    setLoading(true);

  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-purple-600 via-blue-500 to-indigo-700 p-4">
        <motion.div
            initial={{opacity:0, y:40}}
            animate={{opacity:1, y:0}}
            transition={{duration:0.8, ease:"easeOut"}}
            className="w-full max-w-md"
        >
            <Card 
                className="rounded-2xl shadow-2xl backdrop-blur-lg bg-white/80"
            >
                <CardHeader>
                    <CardTitle className="text-center text-2xl font-bold text-gray-800">Welcome Back 👋</CardTitle>
                </CardHeader>


                <CardContent>
                    <form onSubmit={handleLogin} className="space-y-4">
                        {/* Email  */}
                        <div className="space-y-2">
                            <Label htmlFor="email" className="text-gray-700">Email</Label>
                            <Input
                                id="email"
                                type="email"
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
                            required
                            className="mt-1 focus:ring-2 focus:ring-blue-400"
                            placeholder="••••••••"
                            />
                        </div>


                        <motion.div whileHover={{ scale: 1.02 }} whileTap={{ scale: 0.98 }}>
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
  )
}

export default LoginPage