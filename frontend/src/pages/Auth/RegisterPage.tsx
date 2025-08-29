
import { Card, CardHeader, CardTitle, CardContent, CardFooter } from "@/components/ui/card";
import { Select, SelectTrigger, SelectContent, SelectItem, SelectValue } from "@/components/ui/select";
import { ACCESS_TOKEN, REFRESH_TOKEN, type AdminRegisterRequest } from "@/models/auth";
import { RegisterAdmin } from "@/service/auth.api";

import JsCookies from "js-cookie" 
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Label } from "@/components/ui/label";
import toast from "react-hot-toast";
import { useState } from "react";
import { motion } from "framer-motion";
import {useNavigate} from "react-router-dom";


const RegisterPage = () => {
  const [isLoading, setIsLoading] = useState(false);

  const [formData, setFormData] = useState({
    name :"",
    email: "",
    password : "",
    role : "",
  });

  const navigate = useNavigate();


  const handleOnChange = (key:string, value:string) => {
    setFormData({...formData, [key]:value});
  };

  const validate=() => {
      let isAllValid = true;


      if (formData.email.trim().length==0)isAllValid=false;
      if (formData.name.trim().length==0)isAllValid=false;
      if(formData.name.trim().length==0)isAllValid==false;

      return isAllValid;
  }

  const handleSubmit = (e : React.FormEvent) => {
    setIsLoading(true);
    e.preventDefault()
    console.log("submitted : ", formData);


    if (!validate()) {
      toast.error("Please provide missing credentials.");
      return;
    }

    const payload : AdminRegisterRequest = {
      email:formData.email,
      name: formData.name,
      password: formData.password,
      role: formData.role as 'ADMIN | SUPERVISOR',
    };


    async function hitAdminRegisterApi() {

      try {
        const res = await RegisterAdmin(payload);
        console.log(res);

        // 15mins 
        JsCookies.set(ACCESS_TOKEN, res.access_token, {expires: 0.0104,httpOnly:true, secure: true, sameSite: "Strict" })
        // 7 days 
        JsCookies.set(REFRESH_TOKEN, res.refresh_token, {expires: 7,httpOnly:true, secure: true, sameSite: "Strict" })
        console.log("success");


        navigate("/dashboard", {state : {"admin":res.admin}});
      
      } catch (error : any) {

        console.log("failure")
        console.log(error);
        
      } finally {
        setIsLoading(false);
      }
    }

    toast.promise(
      hitAdminRegisterApi,
      {
        loading:"Loading...",
        success:`Admin - ${formData.name} registered successfully`,
        error:`Something went wrong while registering admin.`
      }
    )
  };





  return (
    <div 
      className="flex items-center justify-center min-h-screen bg-gradient-to-br from-indigo-500 via-purple-500 to-pink-500 p-4"
    >
      <motion.div
        initial={{opacity:0, y:40}}
        animate={{opacity:1, y:0}}
        transition={{duration:0.6, ease:"easeOut"}}
        className="w-full max-w-md"
      > 

        <Card
          className="shadow-2xl rounded-2xl backdrop-blur-lg bg-white/90 border border-white/40"
        >
            <CardHeader className="text-center">
              <CardTitle className="text-2xl font-bold text-gray-800">Create Account</CardTitle>
              <p className="text-sm text-gray-500 mt-1">Register to get started with our CRM</p>
            </CardHeader>


            <form onSubmit={handleSubmit}>
              <CardContent className="space-y-4">

                {/* name  */}
                <div className="space-y-2">
                  <Label htmlFor="name"> Name </Label>
                  <Input
                    id="name"
                    type="text"
                    placeholder="John Doe"
                    value={formData.name}
                    onChange={(e)=>handleOnChange("name", e.target.value)}
                    required
                  />
                </div>


                {/* Email  */}
                <div  className="space-y-2">
                  <Label htmlFor="email">Email</Label>
                  <Input
                    id="email"
                    placeholder="johndoe@example.com"
                    type="email"
                    value={formData.email}
                    onChange={(e) => handleOnChange("email", e.target.value)}
                    required
                  />
                </div>

                <div className="space-y-2">
                  <Label htmlFor="password">Password</Label>
                  <Input
                    id="password"
                    type="password"
                    placeholder="••••••••"
                    value={formData.password}
                    onChange={(e)=>handleOnChange("password", e.target.value)}
                    required
                  />
                </div>


                <div className="space-y-2">
                  <Label htmlFor="role">Role</Label>
                  <Select onValueChange={(value)=>handleOnChange("role", value)}>
                    <SelectTrigger id="role">
                      <SelectValue placeholder="Select your role"/>
                    </SelectTrigger>

                    <SelectContent>
                      <SelectItem value="ADMIN">Admin</SelectItem>
                      {/* <SelectItem value="SUPERVISOR">Supervisor</SelectItem> */}
                    </SelectContent>
                  </Select>
                </div>

              </CardContent>


              <CardFooter className="flex flex-col gap-3 mt-8">
                <Button
                  type="submit"
                  className="w-full bg-indigo-600 hover:bg-indigo-700 text-white hover:rounded-xl shadow-md transition-all"
                >
                  { isLoading ? "Registering..." : "Register"}
                </Button>

                <p className="text-sm text-gray-500 text-center">
                  Already have an account?{" "}

                  {/* will add the right href link later  */}
                  <a href="/login" className="text-indigo-800 hover:underline font-medium  ">Log in</a>
                </p>
              </CardFooter>
            </form>
        </Card>
      </motion.div>
    </div>
  );
}

export default RegisterPage