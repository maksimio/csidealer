import { ChakraProvider, extendTheme, ThemeConfig, useColorMode } from '@chakra-ui/react'
import { FC, PropsWithChildren, useEffect } from 'react'

const configChakra: ThemeConfig = {
  useSystemColorMode: false,
  initialColorMode: 'dark',
}

const theme = extendTheme(configChakra)

const WithDarkMode: FC = () => {
  const { colorMode, toggleColorMode } = useColorMode()

  useEffect(() => {
    if (colorMode === 'light') {
      toggleColorMode()
    }
  }, [])

  return null
}

export const WithChakra: FC<PropsWithChildren> = ({ children }) => {
  return (
    <ChakraProvider theme={theme}>
      <WithDarkMode />
      {children}
    </ChakraProvider>
  )
}
