import {
  VStack,
  HStack,
  Text,
  Code,
  NumberInput,
  NumberInputField,
  NumberInputStepper,
  NumberIncrementStepper,
  NumberDecrementStepper,
  Checkbox,
} from '@chakra-ui/react'
import { FC } from 'react'

export const Stats: FC = () => {
  return (
    <VStack alignItems='flex-start'>
      <Checkbox defaultChecked>Ограничить размер файла</Checkbox>
      <HStack>
        <Text>Не более</Text>
        <NumberInput size='sm' maxW='100px'>
          <NumberInputField />
          <NumberInputStepper>
            <NumberIncrementStepper />
            <NumberDecrementStepper />
          </NumberInputStepper>
        </NumberInput>
        <Text>МБайт</Text>
      </HStack>
      <Checkbox defaultChecked>Ограничить число пакетов</Checkbox>
      <HStack>
        <Text>Не более</Text>
        <NumberInput size='sm' maxW='100px'>
          <NumberInputField />
          <NumberInputStepper>
            <NumberIncrementStepper />
            <NumberDecrementStepper />
          </NumberInputStepper>
        </NumberInput>
        <Text>пакетов</Text>
      </HStack>
      <HStack>
        <Text color='gray.400'>Размер файла (МБайт):</Text>
        <Code>100.2</Code>
      </HStack>
      <HStack>
        <Text color='gray.400'>Число пакетов:</Text>
        <Code>103</Code>
      </HStack>
      <HStack>
        <Text color='gray.400'>Длительность записи (мм:сс):</Text>
        <Code>03:11</Code>
      </HStack>
    </VStack>
  )
}
