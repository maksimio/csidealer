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
            <Chart data={store.seriesY} />
          </Box>
        </Card>
        <Card width='25%'>
          <Heading size='md'>Пакет CSI</Heading>
          <Box h='300px'>
            {store.package && <Chart data={store.package.data} />}
          </Box>
        </Card>
      </HStack>
      <Heading>Фазовые значения</Heading>
    </>
  )
})
