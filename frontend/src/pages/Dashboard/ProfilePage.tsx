import React from 'react'
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { motion } from "framer-motion";
import { Mail, User, Shield } from "lucide-react";
import { ADMIN_EMAIL, ADMIN_NAME, ADMIN_ROLE } from '@/models/auth';



const ProfilePage = () => {
  const role = localStorage.getItem(ADMIN_ROLE);
  const email = localStorage.getItem(ADMIN_EMAIL);
  const name = localStorage.getItem(ADMIN_NAME);

  return (
    <motion.div
      initial={{ opacity: 0, scale: 0.8, y: 40 }}
      animate={{ opacity: 1, scale: 1, y: 0 }}
      transition={{ duration: 0.6, ease: "easeOut" }}
      whileHover={{ scale: 1.02 }}
      className="max-w-sm mx-auto "
    >
      <Card className="relative overflow-hidden rounded-2xl shadow-xl border bg-slate-950">
        <motion.div
          initial={{ opacity: 0 }}
          animate={{ opacity: 0.15 }}
          transition={{ duration: 1, repeat: Infinity, repeatType: "mirror" }}
          className="absolute inset-0 bg-gradient-to-r from-purple-500 via-pink-500 to-yellow-500"
        />
        <CardHeader className="relative z-10 text-center">
          <CardTitle className="text-2xl font-bold text-white drop-shadow-lg">
            {name}
          </CardTitle>
        </CardHeader>
        <CardContent className="relative z-10 space-y-4 text-white">
          <motion.div
            initial={{ opacity: 0, x: -30 }}
            animate={{ opacity: 1, x: 0 }}
            transition={{ delay: 0.3 }}
            className="flex items-center space-x-3"
          >
            <Mail className="h-5 w-5 text-pink-300" />
            <span className="text-base">{email}</span>
          </motion.div>

          <motion.div
            initial={{ opacity: 0, x: 30 }}
            animate={{ opacity: 1, x: 0 }}
            transition={{ delay: 0.5 }}
            className="flex items-center space-x-3"
          >
            <Shield className="h-5 w-5 text-yellow-300" />
            <span className="text-base">{role}</span>
          </motion.div>
        </CardContent>
      </Card>
    </motion.div>
  )
}

export default ProfilePage