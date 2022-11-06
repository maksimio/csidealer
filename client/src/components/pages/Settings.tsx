import { Container, FormControl, FormLabel, Switch, useColorMode } from '@chakra-ui/react'
import { FC } from 'react'

const Settings: FC = () => {
  const { colorMode, toggleColorMode } = useColorMode()

  return (
    <Container pt={10}>
      <FormControl display="flex" alignItems="center">
        <FormLabel htmlFor="theme-mode-switch" mb="0">
          Тема: {colorMode === 'light' ? 'светлая' : 'темная'}
        </FormLabel>
        <Switch isChecked={colorMode === 'light'} id="theme-mode-switch" onChange={toggleColorMode} />
      </FormControl>
    </Container>
  )
}

export default Settings
