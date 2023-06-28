import { useControllers } from 'browser'
import { FC, Suspense, lazy, useEffect } from 'react'
import { Routes, Route, Navigate, useLocation } from 'react-router-dom'

const Dashboard = lazy(async () => ({ default: (await import('pages/dashboard')).Dashboard }))
const Recognition = lazy(async () => ({ default: (await import('pages/recognition')).Recognition }))
const Charts = lazy(async () => ({ default: (await import('pages/charts')).Charts }))
const Record = lazy(async () => ({ default: (await import('pages/record')).Record }))
const Devices = lazy(async () => ({ default: (await import('pages/devices')).Devices }))
const Help = lazy(async () => ({ default: (await import('pages/help')).Help }))
const Settings = lazy(async () => ({ default: (await import('pages/settings')).Settings }))
const NotFound = lazy(async () => ({ default: (await import('pages/notfound')).NotFound }))

export const WithRouter: FC = () => {
  const location = useLocation()
  const { navController } = useControllers()

  useEffect(() => {
    navController.setPath(location.pathname.slice(1))
  }, [])
  // TODO: реализовать красивый fallback
  return (
    <Suspense fallback={<div>Загрузка...</div>}>
      <Routes>
        <Route path='/' element={<Navigate to='dashboard' />} />
        <Route path='dashboard' Component={Dashboard} />
        <Route path='recognition' Component={Recognition} />
        <Route path='charts' Component={Charts} />
        <Route path='record' Component={Record} />
        <Route path='devices' Component={Devices} />
        <Route path='help' Component={Help} />
        <Route path='params' Component={Settings} />
        <Route path='*' Component={NotFound} />
        <Route />
      </Routes>
    </Suspense>
  )
}