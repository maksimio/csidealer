import { FC } from 'react'
import { Box, Grid, GridItem, Button, VStack, Flex, Heading } from '@chakra-ui/react'
import { ChartLine, FilePencil, FocusCentered,  Icon, LayoutDashboard, Settings } from 'tabler-icons-react'

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
}

const MenuItem: FC<MenuItemProps> = ({ LeftIcon, text }) => {
  return (
    <Button justifyContent='flex-end' size='sm' variant='ghost' leftIcon={<LeftIcon size='14' />}>{text}</Button>
  )
}

const Menu: FC = () => {
  return (
    <VStack gap={1} alignItems='left' p={1}>
      <MenuItem text='Главная' LeftIcon={LayoutDashboard} />
      <MenuItem text='Графики' LeftIcon={ChartLine} />
      <MenuItem text='Запись' LeftIcon={FilePencil} />
      <MenuItem text='Распознавание' LeftIcon={FocusCentered} />
      <MenuItem text='Параметры' LeftIcon={Settings} />
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