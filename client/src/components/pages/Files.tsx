import { Box, Button, Code, Container, Heading, Input, Stack, Text } from '@chakra-ui/react'
import Card from 'components/shared/Card'
import { useApplication } from 'hooks'
import { observer } from 'mobx-react-lite'
import { ChangeEvent, FC } from 'react'

const LogStatus: FC = observer(() => {
  const { fileStore, fileController } = useApplication()

  const deltaMs = new Date().getTime() - fileStore.startDate.getTime()

  return (
    <Card>
      <Heading size="md" mb={5}>
        Статус: <Code as="span">{fileStore.isLogging ? 'активна' : 'остановлена'}</Code>
      </Heading>
      {fileStore.isLogging ? (
        <>
          <Text>
            Начало: <Code as="span">{fileStore.startDate.toLocaleString()}</Code>
          </Text>
          <Text>
            Длительность: <Code as="span">{new Date(deltaMs).toISOString().split('T')[1].split('.')[0]}</Code>
          </Text>
          <Text>
            МБайт: <Code as="span">{(fileStore.byteCount / 1024 / 1024).toFixed(2)}</Code>
          </Text>
          <Text>
            Пакетов: <Code as="span">{fileStore.packageCount}</Code>
          </Text>
          <Stack pt={5} direction={{ base: 'column', md: 'row' }} alignItems="center">
            <Box>
              <Button onClick={fileController.toggleLog}>Остановить</Button>
            </Box>
            <Text fontSize="xs" as="i">
              Файл будет сохранен в папку logs в месте запуска сервера
            </Text>
          </Stack>
        </>
      ) : (
        <></>
      )}
    </Card>
  )
})

const Files: FC = observer(() => {
  const { fileController, fileStore } = useApplication()

  const handleInputChange = (e: ChangeEvent<HTMLInputElement>) => {
    fileController.setFilename(e.target.value)
  }

  return (
    <Container maxW="2xl" width="full" mt={5}>
      <Heading mb={5}>Запись в файл</Heading>
      <Stack mb={5} direction={{ base: 'column', md: 'row' }}>
        <Input placeholder="Имя файла" value={fileStore.filename} onChange={handleInputChange} />
        <Button disabled={!fileStore.filename.length} variant="outline" onClick={fileController.toggleLog}>
          Запись
        </Button>
      </Stack>
      <LogStatus />
    </Container>
  )
})

export default Files
