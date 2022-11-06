import { Route, Routes, Navigate } from 'react-router-dom'
import { Charts, Devices, Files, NotFound, Settings } from 'components/pages'
import { Container, Grid, GridItem } from '@chakra-ui/react'
import { Topbar, LocationProvider, Sidebar } from 'components'

const App = () => {
  return (
    <>
      <LocationProvider />
      <Grid templateRows={'36px 1fr'} templateColumns={'50px 1fr'} h="100vh">
        <GridItem colSpan={2}>
          <Topbar />
        </GridItem>
        <GridItem>
          <Sidebar />
        </GridItem>
        <GridItem>
          <Container maxW="full">
            <Routes>
              <Route path="/" element={<Navigate to="devices" />} />
              <Route path="*" element={<NotFound />} />
              <Route path="devices" element={<Devices />} />
              <Route path="charts" element={<Charts />} />
              <Route path="files" element={<Files />} />
              <Route path="settings" element={<Settings />} />
            </Routes>
          </Container>
        </GridItem>
      </Grid>
    </>
  )
}

export default App
