import AppRoutes from './Routes/AppRouter'
import { BrowserRouter } from "react-router-dom";
import {  Toaster } from "react-hot-toast";

function App() {
 
 return (
  <>
    <BrowserRouter>
      <AppRoutes/>
    </BrowserRouter>

    <Toaster position='top-center' reverseOrder={false}/>
  </>
  
 );
}

export default App
