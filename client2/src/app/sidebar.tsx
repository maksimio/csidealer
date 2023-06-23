import { FC } from 'react'
import { Grid, GridItem, Button, VStack, Flex, Heading } from '@chakra-ui/react'
import { ChartLine, FilePencil, FocusCentered,  Icon, LayoutDashboard, Settings } from 'tabler-icons-react'
import { useNavigate } from 'react-router-dom'
import { Card } from 'shared/card'

const Logo: FC = () => {
  return (
    <Flex h='full' justifyContent='center' alignItems='center'>
      <Heading size='md' color='blue.400'>Smart Wi-Fi</Heading>
    </Flex>
  )
}

interface MenuItemProps {
  LeftIcon: Icon
  text: string
  path: string
}

const MenuItem: FC<MenuItemProps> = ({ LeftIcon, text, path }) => {
  const navigate = useNavigate()

  function handleNavigate() {
    navigate(path)
  }

  return (
    <Button onClick={handleNavigate} justifyContent='flex-end' size='sm' variant='ghost' leftIcon={<LeftIcon size='14' />}>{text}</Button>
  )
}

const Menu: FC = () => {
  return (
    <VStack gap={1} alignItems='left'>
      <MenuItem path='dashboard' text='Главная' LeftIcon={LayoutDashboard} />
      <MenuItem path='charts' text='Графики' LeftIcon={ChartLine} />
      <MenuItem path='record' text='Запись' LeftIcon={FilePencil} />
      <MenuItem path='recognition' text='Распознавание' LeftIcon={FocusCentered} />
      <MenuItem path='params' text='Параметры' LeftIcon={Settings} />
    </VStack>
  )
}

const ShortInfo: FC = () => {
  return (
    <Card h='full'>Краткая информация</Card>
  )
}

export const Sidebar: FC = () => {
  return (
    <Grid pl={1} pb={1} templateRows='100px 1fr 250px' h='100vh'>
      <GridItem><Logo/></GridItem>
      <GridItem><Menu/></GridItem>
      <GridItem><ShortInfo /></GridItem>
    </Grid>
  )
}