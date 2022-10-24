import { FC } from 'react'
import { Container } from '@chakra-ui/react'
import WebGlChart from 'components/shared/WebGlChart'

const Devices: FC = () => {
  return (
    <Container maxW={'full'}>
      устройства
      <WebGlChart freq={0.001} amp={1} />
    </Container>
  )
}

export default Devices