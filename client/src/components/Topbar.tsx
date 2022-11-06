import { Code, Flex, Heading, HStack, Spacer, Text, useColorMode } from '@chakra-ui/react'
import { useApplication } from 'hooks'
import { observer } from 'mobx-react-lite'
import { FC } from 'react'

const Topbar: FC = observer(() => {
  const { colorMode } = useColorMode()
  const { csiStore, deviceStore } = useApplication()

  return (
    <Flex h="full" direction="row" bg={colorMode === 'light' ? 'blue.300' : 'blue.800'}>
      <HStack>
        <Heading size="md" ml={2} letterSpacing={2}>
          <Text as="span" color={colorMode === 'light' ? 'blue.800' : 'blue.300'}>
            CSI
          </Text>{' '}
          <Text as="span" color={colorMode === 'light' ? 'blue.800' : 'blue.300'}>
            Dealer
          </Text>
        </Heading>
      </HStack>
      <Spacer />
      <HStack pr={2}>
        <Code colorScheme={deviceStore.isClientConnect ? 'green' : 'red'}>{deviceStore.isClientConnect ? 'Подключено' : 'Отключено'}</Code>
        <Code>Частота: {csiStore.frequency.toFixed(1)} Гц</Code>
      </HStack>
    </Flex>
  )
})

export default Topbar
