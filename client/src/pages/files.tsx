import { Box, Heading } from '@chakra-ui/react'
import { FC } from 'react'
import { Record } from 'widgets/record'
import { Grid, GridItem } from '@chakra-ui/react'
import { Markup } from 'widgets/markup'
import { Steps } from 'widgets/steps'

export const Files: FC = () => {
  return (
    <Box>
      <Grid h='200px' templateRows='repeat(2, 1fr)' templateColumns='repeat(2, 1fr)' gap={8}>
        <GridItem colSpan={2}>
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
      </Grid>
    </Box>
  )
}
