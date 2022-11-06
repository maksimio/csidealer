import { Container, Box } from '@chakra-ui/react'
import WebGlChart from 'components/shared/WebGlChart'
import WebGlChart2 from 'components/shared/WebGlChart2'
import { observer } from 'mobx-react-lite'
import { FC } from 'react'

const Charts: FC = observer(() => {
  return (
    <Container maxW={'full'}>
      <Box height="67vh" display="flex" width="100%">
        <Box width="75%">
          <WebGlChart />
        </Box>
        <Box width="25%">
          <WebGlChart2 />
        </Box>
      </Box>
    </Container>
  )
})

export default Charts
