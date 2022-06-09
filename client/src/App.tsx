import { Container } from '@chakra-ui/react'
import { Route, Routes } from 'react-router-dom'
import Devices from './pages/Devices'
import Log from './pages/Log'
import NotFound from './pages/NotFound'
import SimpleSidebar from './pages/Sidebar'

const App = () => {
  return (
    <SimpleSidebar>
      <Routes>
        <Route path='*' element={<NotFound />} />
        <Route path='log' element={<Log />} />
        <Route path='devices' element={<Devices />} />
      </Routes>
    </SimpleSidebar>
  )
}

export default App
