import { Box, Heading } from '@chakra-ui/react'
import { FC } from 'react'
import { Record } from 'widgets/record'
import { FileList } from 'widgets/filelist'

export const Files: FC = () => {
  return (
    <Box>
      <Heading>Запись CSI в файл</Heading>
      <Record />
      <Heading>Список файлов</Heading>
      <FileList />
    </Box>
  )
}
