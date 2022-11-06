import { FC } from 'react'
import { Code, Container } from '@chakra-ui/react'
import { Stack, Text, Button } from '@chakra-ui/react'
import Card from 'components/shared/Card'

const RouterCard = () => {
  return (
    <Card>
      <Stack direction="row" alignItems="center">
        <Text fontWeight="semibold">Your Privacy</Text>
      </Stack>

      <Stack direction={{ base: 'column', md: 'row' }} justifyContent="space-between">
        <Code fontSize={{ base: 'sm' }} textAlign={'left'} maxW={'4xl'}>
          uuid маршрутизатора
        </Code>
        <Stack direction={{ base: 'column', md: 'row' }}>
          <Button variant="outline" colorScheme="green">
            Cookie Preferences
          </Button>
          <Button colorScheme="green">OK</Button>
        </Stack>
      </Stack>
    </Card>
  )
}

const Devices: FC = () => {
  return (
    <Container maxW="2xl" width="full">
      <RouterCard />
      <RouterCard />
      <RouterCard />
      <RouterCard />
    </Container>
  )
}

export default Devices
