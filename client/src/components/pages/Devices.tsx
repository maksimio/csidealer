import { FC } from 'react'
import { Box, Container } from '@chakra-ui/react'
import WebGlChart from 'components/shared/WebGlChart'
import WebGlChart2 from 'components/shared/WebGlChart2'
import WebGlChart3 from 'components/shared/WebGlChart3'
import WebGlChart4 from 'components/shared/WebGlChart4'

const Devices: FC = () => {
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
      <Box height="27vh" display="flex">
        <Box width="75%">
          <WebGlChart4 />
        </Box>
        <Box width="25%">
          <WebGlChart3 />
        </Box>
      </Box>
    </Container>
  )
}

export default Devices
