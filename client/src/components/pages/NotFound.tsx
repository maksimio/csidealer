import { Button, Container, Text } from '@chakra-ui/react'
import { FC } from 'react'
import { useNavigate } from 'react-router-dom'

const NotFound: FC = () => {
  const navigate = useNavigate()

  return (
    <Container>
      <Text align="center" fontSize="10rem" colorScheme="blue">
        404
      </Text>
      <Text align="center" fontSize="xl">
        Страница не найдена
      </Text>
      <Button mt="10" width="full" onClick={() => navigate('/')}>
        На главную
      </Button>
    </Container>
  )
}

export default NotFound
