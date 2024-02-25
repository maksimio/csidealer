import { Button, Heading, Divider, HStack, Box, VStack, AbsoluteCenter, Text } from '@chakra-ui/react'
import { useControllers } from 'browser'
import { FC } from 'react'
import { Send, Box as BoxIcon } from 'tabler-icons-react'

const Router: FC = () => {
  return (
    <HStack minH={250} minW={250} bg='cyan.700' borderRadius={'50%'} borderWidth={5} borderColor='gray.400'>
      <VStack w='full' h='full' gap={3}>
        <Heading>Tx</Heading>
        <Text>192.168.1.1</Text>
        <HStack gap={6}>
          <BoxIcon />
          <Send />
        </HStack>
      </VStack>
    </HStack>
  )
}

const Connection: FC = () => {
  const { routerController } = useControllers()

  return (
    <VStack w='full' gap={7}>
      <Button onClick={routerController.reconnect}>Установить соединение</Button>
      <Box position='relative' padding='10' w='full'>
        <Divider borderWidth={3} borderRadius={5} />
        <AbsoluteCenter px='4' bg='gray.800'>
          <Text>соединение установлено</Text>
        </AbsoluteCenter>
      </Box>
      <Button colorScheme='green'>Начать прием CSI</Button>
    </VStack>
  )
}

const TransmitInfo: FC = () => {
  return (
    <HStack w='full' justifyContent='space-around' pt={10} pr={10} pl={10}>
      <Router />
      <Connection />
      <Router />
    </HStack>
  )
}

export const Dashboard: FC = () => {
  return (
    <VStack minW='full'>
      <TransmitInfo />
    </VStack>
  )
}
