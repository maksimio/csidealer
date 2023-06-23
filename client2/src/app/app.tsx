import { FC, useEffect } from 'react'
import { ChakraProvider, Grid, GridItem, useColorMode } from '@chakra-ui/react'
import React from 'react'
import ReactDOM from 'react-dom/client'
import { Sidebar } from './sidebar'
import { BrowserRouter, Route, Navigate, Routes } from 'react-router-dom'
import { Dashboard } from 'pages/dashboard'
import { Settings } from 'pages/settings'
import { Recognition } from 'pages/recognition'
import { Charts } from 'pages/charts'
import { Record } from 'pages/record'
import { Devices } from 'tabler-icons-react'
import { Help } from 'pages/help'

const ColorMode: FC = () => {
  const { colorMode, toggleColorMode } = useColorMode()

  useEffect(() => {
    if (colorMode === 'light') {
      toggleColorMode()
    }
  }, [])

  return null
}

const App: FC = () => {
  return (
    <ChakraProvider>
      <BrowserRouter>
        <ColorMode />
        <Grid gap={2} templateColumns='200px 1fr' h='full'>
          <GridItem h='full'>
            <Sidebar />
          </GridItem>
          <GridItem>
            <Routes>
              <Route path='/' element={<Navigate to='dashboard' />} />
              <Route path='dashboard' Component={Dashboard} />
              <Route path='recognition' Component={Recognition} />
              <Route path='charts' Component={Charts} />
              <Route path='record' Component={Record} />
              <Route path='devices' Component={Devices} />
              <Route path='help' Component={Help} />
              <Route path='params' Component={Settings} />
              <Route />
            </Routes>
          </GridItem>
        </Grid>
      </BrowserRouter>
    </ChakraProvider>
  )
}

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
)
