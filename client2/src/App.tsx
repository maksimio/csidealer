import { FC } from 'react'
import { Button, ChakraProvider } from '@chakra-ui/react'
import React from 'react'
import ReactDOM from 'react-dom/client'

const App: FC = () => {

  return (
    <ChakraProvider>
    <Button >hello world</Button>
  </ChakraProvider>
  )
}


ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
)

