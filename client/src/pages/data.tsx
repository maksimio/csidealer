import { HStack, Heading, Box } from '@chakra-ui/react'
import { useStore } from 'browser'
import { observer } from 'mobx-react-lite'
import { FC } from 'react'
import { Card } from 'shared/card'
import { Chart } from 'widgets/chart'

export const Data: FC = observer(() => {
  const store = useStore()

  return (
    <>
      <Heading>Амплитудные значения</Heading>
      <HStack>
        <Card width='75%'>
          <Heading size='md'>График поднесущих во времени</Heading>
          <Box h='300px'>
            <Chart delimeter={150} shift={-1} data={store.abs} />
          </Box>
        </Card>
        <Card width='25%'>
          <Heading size='md'>Пакет CSI</Heading>
          <Box h='300px'>
            {store.package && <Chart delimeter={150} shift={-1} data={store.package.abs} />}
          </Box>
        </Card>
      </HStack>
      <Heading mt={5}>Фазовые значения</Heading>
      <HStack>
        <Card width='75%'>
          <Heading size='md'>График поднесущих во времени</Heading>
          <Box h='300px'>
            <Chart delimeter={5} shift={0}  data={store.phase} />
          </Box>
        </Card>
        <Card width='25%'>
          <Heading size='md'>Пакет CSI</Heading>
          <Box h='300px'>
            {store.package && <Chart delimeter={5} shift={0} data={store.package.phase} />}
          </Box>
        </Card>
      </HStack>
    </>
  )
})
