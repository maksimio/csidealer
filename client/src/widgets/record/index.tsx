import { Box, Button, HStack, Text, Code, VStack } from '@chakra-ui/react'
import { FC } from 'react'
import { Card } from 'shared/card'
import { Stats } from './stats'
import { NameConf } from './nameconf'
import { observer } from 'mobx-react-lite'
import { useControllers, useStore } from 'browser'

const FileName: FC = observer(() => {
  const store = useStore()

  return (
    <VStack gap={1} mt={5} w='full'>
      <Text fontSize='sm' color='gray.500'>
        имя файла:
      </Text>
      <Code fontSize='md'>{store.filename}</Code>
    </VStack>
  )
})

const RecordToggler: FC = observer(() => {
  const store = useStore()
  const { recordController } = useControllers()

  return (
    <Box mt={5}>
      <Button onClick={recordController.toggleRecording} w='full' p={10} colorScheme={store.recording ? 'red' : 'green'}>
        {store.recording ? 'ОСТАНОВИТЬ ЗАПИСЬ' : 'НАЧАТЬ ЗАПИСЬ'}
      </Button>
    </Box>
  )
})

export const Record: FC = () => {
  return (
    <Card maxW='650px'>
      <HStack alignItems='flex-start' justifyContent='space-between'>
        <NameConf />
        <Stats />
      </HStack>
      <FileName />
      <RecordToggler />
    </Card>
  )
}
