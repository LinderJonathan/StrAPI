import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.tsx'

createRoot
const rootElement = (document.getElementById('root'))
  
if (!rootElement) {
  throw new Error("Root element not found")
}

createRoot(rootElement).render(
  /*TODO: Investigate if BrowserRouter can be used here instead */
  <StrictMode>
    <App />
  </StrictMode>
)