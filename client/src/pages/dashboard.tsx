import { Button, Heading, Divider, HStack, Box, VStack, AbsoluteCenter, Text, Tooltip, useTheme, Link } from '@chakra-ui/react'
import { useControllers, useStore } from 'browser'
import { observer } from 'mobx-react-lite'
import { FC, useEffect } from 'react'
import { RouterInfo } from 'services'
import { Send, Box as BoxIcon } from 'tabler-icons-react'

interface RouterProps {
  title: string
  info: RouterInfo
}

const Router: FC<RouterProps> = ({ info, title }) => {
  const green = useTheme().colors.green['400']

  return (
    <HStack
      minH={250}
      minW={250}
      bg='cyan.800'
      borderRadius={'50%'}
      borderWidth={5}
      borderColor={info.is_connected ? 'green.400' : 'red.400'}
    >
      <VStack w='full' h='full' gap={3}>
        <Heading>{title}</Heading>
        <Link href={`http://${info.addr}`} isExternal>
          {info.addr}
        </Link>
        <HStack gap={6}>
          <BoxIcon color={info.is_clientmain_active ? green : undefined} />
          <Send color={info.is_sendData_active ? green : undefined} />
        </HStack>
      </VStack>
    </HStack>
  )
}

const Connection: FC = observer(() => {
  const store = useStore()
  const { routerController } = useControllers()

  const isConnect = store.routerStatus.tx.is_connected && store.routerStatus.rx.is_connected
  const isTransmit = store.routerStatus.rx.is_clientmain_active && store.routerStatus.tx.is_sendData_active

  return (
    <VStack w='full' gap={7}>
      <Button onClick={routerController.reconnect}>Установить соединение</Button>
      <Box position='relative' padding='10' w='full'>
        <Divider
          borderWidth={3}
          borderRadius={5}
          variant={isConnect ? 'solid' : 'dashed'}
          borderColor={isTransmit ? 'green.400' : undefined}
        />
        <AbsoluteCenter px='4' bg='gray.800'>
          <Text color={isTransmit ? 'green.400' : 'gray.500'}>
            {isConnect ? (isTransmit ? 'идет прием CSI' : 'соединение установлено') : 'соединение не установлено'}
          </Text>
        </AbsoluteCenter>
      </Box>
      <Button isDisabled={!isConnect} colorScheme={isTransmit ? 'red' : 'green'} onClick={routerController.toggleTransmitCSI}>
        {isTransmit ? 'Остановить прием CSI' : 'Начать прием CSI'}
      </Button>
    </VStack>
  )
})

const TransmitInfo: FC = observer(() => {
  const store = useStore()

  return (
    <HStack w='full' justifyContent='space-around' pt={10} pr={10} pl={10}>
      <Router title='Tx' info={store.routerStatus.tx} />
      <Connection />
      <Router title='Rx' info={store.routerStatus.rx} />
    </HStack>
  )
})

export const Dashboard: FC = () => {
  const { routerController } = useControllers()

  useEffect(() => {
    routerController.getStatus()
  }, [])

  return (
    <VStack minW='full'>
      <TransmitInfo />
    </VStack>
  )
}
