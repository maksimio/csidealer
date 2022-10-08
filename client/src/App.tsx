import { useApplication } from 'hooks/ApplicationContext'
import { useEffect } from 'react'
import { Route, Routes, useLocation, Navigate } from 'react-router-dom'
import {
  Dashboard,
  Devices,
  FileLog,
  NotFound,
  Settings,
  Sidebar,
} from 'components/pages'

const App = () => {
  const location = useLocation()
  const { layoutController } = useApplication()

  useEffect(() => {
    layoutController.setLocation(location)
  }, [layoutController, location])

  return (
    <Sidebar>
      <Routes>
        <Route path='/' element={<Navigate to='dashboard' />} />
        <Route path='*' element={<NotFound />} />
        <Route path='log' element={<FileLog />} />
        <Route path='devices' element={<Devices />} />
        <Route path='settings' element={<Settings />} />
        <Route path='dashboard' element={<Dashboard />} />
      </Routes>
    </Sidebar>
  )
}

export default App
