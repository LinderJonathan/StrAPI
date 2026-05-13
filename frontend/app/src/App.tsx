import './App.css'
import { BrowserRouter, Routes, Route } from "react-router-dom"
import Navbar from "./components/Navbar"
import Statistics from "./pages/Statistics"
import Home from "./pages/Home"
import MyActivities from './pages/MyActivities'

function App() {
  return (
    <BrowserRouter>
      <Navbar />
      <Routes>
        <Route path="/Statistics" element={<Statistics />} />
        <Route path="/MyActivities" element={<MyActivities />} />
        <Route path="/" element={<Home />} />
      </Routes>
    </BrowserRouter>
  )
}

export default App