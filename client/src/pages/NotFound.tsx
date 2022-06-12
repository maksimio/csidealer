import { Button, Center, Container, Text } from '@chakra-ui/react'
import { FC } from 'react'
import { useNavigate } from 'react-router-dom'


const NotFound: FC = () => {
  const navigate = useNavigate()

  return (
    <Container>
      <Text align='center' color='blackAlpha.300' fontSize='10rem'>404</Text>
      <Text align='center' fontSize='xl'>Страница не найдена</Text>
      <Button mt='10' width='full' onClick={() => navigate('dashboard')}>На главную</Button>
    </Container>
  )
}

export default NotFound