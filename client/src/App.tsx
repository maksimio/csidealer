import { Route, Routes, Navigate } from 'react-router-dom'
import {
  Dashboard,
  Devices,
  FileLog,
  NotFound,
  Settings,
} from 'components/pages'
import { Grid, GridItem } from '@chakra-ui/react'
import { Statusbar, LocationProvider, Sidebar } from 'components'

const App = () => {
  return (
    <>
      <LocationProvider />
      <Grid templateRows={'1fr 20px'} templateColumns={'50px 1fr'} h="100vh">
        <GridItem background={'red'}>
          <Sidebar />
        </GridItem>
        <GridItem background={'green'}>
          <Routes>
            <Route path="/" element={<Navigate to="dashboard" />} />
            <Route path="*" element={<NotFound />} />
            <Route path="log" element={<FileLog />} />
            <Route path="devices" element={<Devices />} />
            <Route path="settings" element={<Settings />} />
            <Route path="dashboard" element={<Dashboard />} />
          </Routes>
        </GridItem>
        <GridItem background={'blue'} colSpan={2}>
          <Statusbar />
        </GridItem>
      </Grid>
    </>
  )
}

export default App
