import { HStack, Heading } from '@chakra-ui/react'
import { FC } from 'react'
import { Card } from 'shared/card'
import { ChartSubcarriers } from 'widgets/chartsubcarriers'
import { ChartTimeline } from 'widgets/charttimeline'

export const Data: FC = () => {
  return (
    <>
      <Heading>Амплитудные значения</Heading>
      <HStack>
        <Card h='400px' width="75%">
          <Heading size='md'>График поднесущих во времени</Heading>
          <ChartTimeline />
        </Card>
        <Card h='400px' width="25%">
          <Heading size='md'>Пакет CSI</Heading>
          <ChartSubcarriers />
        </Card>
      </HStack>
      <Heading>Фазовые значения</Heading>
    </>
  )
}