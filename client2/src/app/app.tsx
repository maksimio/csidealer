import { FC, useEffect } from 'react'
import { ChakraProvider, Grid, GridItem, useColorMode } from '@chakra-ui/react'
import React from 'react'
import ReactDOM from 'react-dom/client'
import { Sidebar } from './sidebar'

const ColorMode: FC = () => {
  const { colorMode, toggleColorMode } = useColorMode()

  useEffect(() => {
    console.log(colorMode)
    if (colorMode === 'light') {
      toggleColorMode()
    }
  }, [])

  return null
}

const App: FC = () => {



  return (
    <ChakraProvider>
      <ColorMode />
      <Grid gap={2} templateColumns='200px 1fr' h='full'>
        <GridItem h='full'><Sidebar /></GridItem>
        <GridItem>Основное окно</GridItem>
      </Grid>
    </ChakraProvider>
  )
}


ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
)

