import { Button } from '@/components/ui/button'
import { motion } from 'framer-motion'
import { Home, Sparkles } from 'lucide-react'
import { useNavigate } from 'react-router-dom'

const NotFound = () => {
  const navigate = useNavigate();
  
  return (
    <div
      className='relative flex h-screen w-full flex-col items-center justify-center bg-gradient-to-br from-gray-950 via-gray-900 to-gray-800 text-center overflow-hidden'
    >

      <motion.div
        initial={{opacity:0}}
        animate={{opacity:[0.2, 0.5, 0.2]}}
        transition={{repeat:Infinity, duration: 6}}
        className='absolute -top-20 -left-20 h-96 w-96 rounded-full bg-purple-700 blur-3xl opacity-30'
      />

      <motion.div
        initial={{opacity:0}}
        animate={{opacity:[0.2, 0.5, 0.2]}}
        transition={{repeat: Infinity, duration:6}}
        className='absolute bottom-[-100px] right-[-100px] h-[400px] w-[400px]  rounded-full bg-pink-600 blur-3xl opacity-30'
      />



      <motion.h1

        className='text-[8rem] font-extrabold tracking-widest text-transparent bg-clip-text bg-gradient-to-r from-purple-400 to-pink-600 drop-shadow-lg'
      >
        404 
      </motion.h1>

      {/* Subtitle */}
      <motion.p
        initial={{ y: 20, opacity: 0 }}
        animate={{ y: 0, opacity: 1 }}
        transition={{ delay: 0.3, duration: 0.6 }}
        className="mt-4 text-xl text-gray-300 flex items-center gap-2"
      >
        <Sparkles className="h-5 w-5 text-purple-400 animate-pulse" />
        Oops! Page not found.
      </motion.p>


       <motion.div
        initial={{ y: 30, opacity: 0 }}
        animate={{ y: 0, opacity: 1 }}
        transition={{ delay: 0.6, duration: 0.6 }}
        className="mt-8"
      >
        <Button
          onClick={() => navigate("/atheena/dashboard-v2")}
          className="flex items-center gap-2 rounded-2xl px-6 py-5 text-lg shadow-lg shadow-purple-500/30 hover:scale-105 transition-transform"
        >
          <Home className="h-5 w-5" />
          Go Home
        </Button>
      </motion.div>
    </div>
  )
}

export default NotFound