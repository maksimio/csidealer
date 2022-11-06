import { FC, ReactNode } from 'react'
import { useColorMode, Stack } from '@chakra-ui/react'

interface CardProps {
  children: ReactNode
}

const Card: FC<CardProps> = ({ children }) => {
  const { colorMode } = useColorMode()

  return (
    <Stack p="4" boxShadow="lg" borderRadius="md" bg={colorMode === 'light' ? 'blue.100' : 'blue.900'}>
      {children}
    </Stack>
  )
}

export default Card
