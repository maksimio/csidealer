import { FC } from 'react'
import { Code, Container, Heading, VStack } from '@chakra-ui/react'
import { Stack, Button } from '@chakra-ui/react'
import Card from 'components/shared/Card'

const RouterCard = () => {
  return (
    <Card>
      <Stack direction="row" alignItems="center">
        <Heading size="md">Маршрутизатор</Heading>
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
      <VStack mt={5}>
        <RouterCard />
        <RouterCard />
        <RouterCard />
        <RouterCard />
      </VStack>
    </Container>
  )
}

export default Devices
