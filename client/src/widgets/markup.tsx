import { Button, HStack, Box, Tag, TagLabel, TagCloseButton, VStack, Text, Heading, Code } from '@chakra-ui/react'
import { useControllers, useStore } from 'browser'
import { observer } from 'mobx-react-lite'
import { FC } from 'react'
import { Card } from 'shared/card'
import { Mark } from 'store'

interface ChipProps {
  mark: Mark
}

const Chip: FC<ChipProps> = observer(({ mark }) => {
  const { recordController } = useControllers()

  return (
    <Tag
      onClick={() => recordController.toggleActiveMark(mark.id)}
      cursor='pointer'
      mr={3}
      mb={3}
      size='lg'
      borderRadius='full'
      variant='solid'
      colorScheme={mark.isActive ? 'green' : 'gray'}
    >
      <TagLabel>{mark.text}</TagLabel>
      <TagCloseButton onClick={() => recordController.deleteMark(mark.id)} />
    </Tag>
  )
})

const Marks: FC = observer(() => {
  const store = useStore()

  const marks = [...store.marks.values()].map((v) => <Chip key={v.id} mark={v} />)

  const emptyText = (
    <VStack justifyContent='center' h='full'>
      <Text color='gray.400'>Список объектов для разметки пуст</Text>
      <Text color='gray.400'>Нажмите Добавить, чтобы создать объект</Text>
    </VStack>
  )

  return (
    <Box h='full' w='full'>
      {marks.length ? marks : emptyText}
    </Box>
  )
})

export const Markup: FC = observer(() => {
  const store = useStore()
  const { recordController } = useControllers()

  return (
    <Card maxW='650px' h='450px'>
      <VStack justifyContent='space-between' h='full'>
        <Heading size='md' w='full'>
          Ваши метки
        </Heading>
        <Text color='gray.400'>
          Добавьте собственные метки и активируйте / выключайте их в нужный момент записи. Метки свяжутся с файлом данных CSI
        </Text>
        <Marks />
        <Text>
          Всего меток: <Code>{store.marks.size}</Code>
        </Text>
        <HStack w='full' mt={10}>
          <Button onClick={recordController.unactiveMarks} w='200px'>
            Очистить выбор
          </Button>
          <HStack w='full' justifyContent='flex-end'>
            <Button onClick={recordController.addMark} colorScheme='green'>
              Добавить
            </Button>
            <Button onClick={recordController.clearMarks} colorScheme='red'>
              Очистить список
            </Button>
          </HStack>
        </HStack>
      </VStack>
    </Card>
  )
})
