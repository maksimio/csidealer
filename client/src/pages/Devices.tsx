import { FC } from 'react'
import Device from 'components/Device'
import { Container } from '@chakra-ui/react'

const Devices: FC = () => {
  return (
    <Container>
      <Device />
      <Device />
      <Device />
    </Container>
  )
}

export default Devices