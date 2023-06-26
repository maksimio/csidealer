import { Charts } from 'pages/charts'
import { Dashboard } from 'pages/dashboard'
import { Devices } from 'pages/devices'
import { Help } from 'pages/help'
import { NotFound } from 'pages/notfound'
import { Recognition } from 'pages/recognition'
import { Record } from 'pages/record'
import { Settings } from 'pages/settings'
import { FC } from 'react'
import { Routes, Route, Navigate } from 'react-router-dom'


export const WithRouter: FC = () => {
  return (
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
  )
}