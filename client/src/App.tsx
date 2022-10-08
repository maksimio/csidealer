import { Route, Routes, Navigate } from 'react-router-dom'
import {
  Dashboard,
  Devices,
  FileLog,
  NotFound,
  Settings,
} from 'components/pages'
import { Container, Grid, GridItem } from '@chakra-ui/react'
import { Statusbar, LocationProvider, Sidebar } from 'components'

const App = () => {
  return (
    <>
      <LocationProvider />
      <Grid templateRows={'1fr 20px'} templateColumns={'50px 1fr'} h="100vh">
        <GridItem>
          <Sidebar />
        </GridItem>
        <GridItem>
          <Container maxW="full">
            <Routes>
              <Route path="/" element={<Navigate to="dashboard" />} />
              <Route path="*" element={<NotFound />} />
              <Route path="dashboard" element={<Dashboard />} />
              <Route path="devices" element={<Devices />} />
              <Route path="log" element={<FileLog />} />
              <Route path="settings" element={<Settings />} />
            </Routes>
          </Container>
        </GridItem>
        <GridItem background={'cyan'} colSpan={2}>
          <Statusbar />
        </GridItem>
      </Grid>
    </>
  )
}

export default App
