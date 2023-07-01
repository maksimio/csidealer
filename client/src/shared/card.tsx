import { Box, ChakraProps } from '@chakra-ui/react'
import { FC, PropsWithChildren } from 'react'

export const Card: FC<PropsWithChildren & ChakraProps> = ({ children, ...props }) => {
  return <Box p={2} borderRadius={5} bg='gray.700' {...props}>{children}</Box>
}
