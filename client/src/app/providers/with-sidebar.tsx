import { FC, PropsWithChildren } from 'react'
import { Grid, GridItem, Button, VStack, Flex, Heading } from '@chakra-ui/react'
import { AccessPoint, ChartLine, FilePencil, FocusCentered, Help, Icon, Settings } from 'tabler-icons-react'
import { useNavigate } from 'react-router-dom'
import { Card } from 'shared/card'
import { observer } from 'mobx-react-lite'
import { useControllers, useStore } from 'browser'

const Logo: FC = () => {
  return (
    <Flex h='full' justifyContent='center' alignItems='center'>
      <Heading size='md' color='blue.400'>
        Smart Wi-Fi
      </Heading>
    </Flex>
  )
}

interface MenuItemProps {
  LeftIcon: Icon
  text: string
  path: string
  isDisabled?: boolean
}

const MenuItem: FC<MenuItemProps> = observer(({ LeftIcon, text, path, isDisabled }) => {
  const store = useStore()
  const { navController } = useControllers()
  const navigate = useNavigate()

  function handleNavigate() {
    navController.setPath(path)
    navigate(path)
  }

  const active = path === store.path

  return (
    <Button
      isDisabled={isDisabled}
      onClick={handleNavigate}
      justifyContent='flex-end'
      size='sm'
      variant={active ? 'solid' : 'ghost'}
      leftIcon={<LeftIcon size='18' />}
    >
      {text}
    </Button>
  )
})

const Menu: FC = () => {
  return (
    <VStack gap={2} alignItems='left'>
      <MenuItem path='dashboard' text='Главная' LeftIcon={AccessPoint} />
      <MenuItem path='data' text='Данные' LeftIcon={ChartLine} />
      <MenuItem path='files' text='Запись' LeftIcon={FilePencil} />
      <MenuItem isDisabled path='recognition' text='Распознавание' LeftIcon={FocusCentered} />
      <MenuItem path='params' text='Параметры' LeftIcon={Settings} />
      <MenuItem isDisabled path='help' text='Помощь' LeftIcon={Help} />
    </VStack>
  )
}

const Sidebar: FC = () => {
  return (
    <Grid pl={1} pb={1} templateRows='100px 1fr' h='100vh'>
      <GridItem>
        <Logo />
      </GridItem>
      <GridItem>
        <Menu />
      </GridItem>
    </Grid>
  )
}

export const WithSidebar: FC<PropsWithChildren> = ({ children }) => {
  return (
    <Grid gap={2} templateColumns='200px 1fr' h='full'>
      <GridItem h='full'>
        <Sidebar />
      </GridItem>
      <GridItem pt={3} pr={3}>
        {children}
      </GridItem>
    </Grid>
  )
}
