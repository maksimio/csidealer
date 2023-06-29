import { Code, HStack, Heading, Text, VStack } from '@chakra-ui/react'
import { FC } from 'react'

export const Settings: FC = () => {
  return (
    <VStack alignItems='flex-start'>
      <Heading>О приложении</Heading>
      <Text>Приложение предназначено для удобной работы с данными CSI (Channel State Information), полученных с маршрутизаторов. Клиент и сервер поставляются совместно.</Text>
      <HStack>
        <Text>Номер сборки</Text>
        <Code>{__COMMIT_NUMBER}</Code>
      </HStack>
      <HStack>
        <Text>Код сборки</Text>
        <Code>{__COMMIT_HASH}</Code>
      </HStack>
      <HStack>
        <Text>Дата</Text>
        <Code>{__COMMIT_DATE}</Code>
      </HStack>
    </VStack>
  )
}