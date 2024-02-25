import { Button, Heading, Divider, HStack, Text, VStack } from '@chakra-ui/react'
import { FC } from 'react'
import { Link, Send, Box } from 'tabler-icons-react'

const Router: FC = () => {
  return (
    <VStack gap={5}>
      <Heading>192.168.1.1</Heading>
      <HStack gap={5}>
        <Link />
        <Box />
        <Send />
      </HStack>
    </VStack>
  )
}

const TransmitInfo: FC = () => {
  return (
    <HStack w='full' justifyContent='space-around' pt={10} pr={10} pl={10}>
      <Router />
      <VStack w='full'>
        <Text>передача</Text>
        <Divider size='lg' borderWidth={3} borderRadius={5} width='100%' />
        <Button>Начать</Button>
      </VStack>
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
