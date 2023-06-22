import { FC } from 'react'
import { Box, Grid, GridItem, Button, VStack, Flex, Heading } from '@chakra-ui/react'
import { ChartLine, FilePencil, FocusCentered,  Icon, LayoutDashboard, Settings } from 'tabler-icons-react'
import { useNavigate } from 'react-router-dom'

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
    <VStack gap={1} alignItems='left' p={1}>
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
    <Box>Краткая информация</Box>
  )
}

export const Sidebar: FC = () => {
  return (
    <Grid templateRows='100px 1fr 100px' h='100vh'>
      <GridItem><Logo /></GridItem>
      <GridItem><Menu /></GridItem>
      <GridItem><ShortInfo /></GridItem>
    </Grid>
  )
}