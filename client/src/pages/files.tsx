import { Heading } from '@chakra-ui/react'
import { FC } from 'react'
import { Record } from 'widgets/record'
import { Grid, GridItem } from '@chakra-ui/react'
import { Markup } from 'widgets/markup'
import { Steps } from 'widgets/steps'
import { MarkStory } from 'widgets/markstory'

export const Files: FC = () => {
  return (
    <Grid h='200px' templateRows='repeat(2, 1fr)' templateColumns='1fr 1fr 0.5fr' gap={8}>
      <GridItem colSpan={3}>
        <Steps />
      </GridItem>
      <GridItem colSpan={1}>
        <Heading>Запись CSI в файл</Heading>
        <Record />
      </GridItem>
      <GridItem colSpan={1}>
        <Heading>Разметка данных</Heading>
        <Markup />
      </GridItem>
      <GridItem colSpan={1}>
        <Heading>История меток</Heading>
        <MarkStory />
      </GridItem>
    </Grid>
  )
}
