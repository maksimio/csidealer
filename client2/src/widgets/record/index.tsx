import { Box, Button, HStack } from '@chakra-ui/react'
import { FC } from 'react'
import { Card } from 'shared/card'
import { Stats } from './stats'
import { NameConf } from './nameconf'

export const Record: FC = () => {
  return (
    <Card maxW='700px'>
      <HStack alignItems='flex-start'>
        <NameConf />
        <Stats />
      </HStack>
      <Box mt={2}>
        <Button w='full' p={10} colorScheme='green'>
          НАЧАТЬ ЗАПИСЬ
        </Button>
      </Box>
    </Card>
  )
}
