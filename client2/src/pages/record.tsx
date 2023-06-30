import {
  Box,
  Button,
  Checkbox,
  Code,
  HStack,
  Heading,
  IconButton,
  Input,
  NumberDecrementStepper,
  NumberIncrementStepper,
  NumberInput,
  NumberInputField,
  NumberInputStepper,
  Radio,
  RadioGroup,
  Text,
  VStack,
} from '@chakra-ui/react'
import { FC } from 'react'
import { Card } from 'shared/card'
import { Plus, X } from 'tabler-icons-react'

export const Record: FC = () => {
  return (
    <Box>
      <Heading>Запись CSI в файл</Heading>
      <Card maxW='700px'>
        <HStack alignItems='flex-start'>
          <VStack alignItems='flex-start'>
            <HStack gap={5}>
              <Checkbox defaultChecked>Тип</Checkbox>
              <Checkbox defaultChecked>Дата</Checkbox>
            </HStack>
            <RadioGroup flexDirection='column'>
              <HStack alignItems='flex-start'>
                <Radio value='train'>train</Radio>
                <Radio value='test'>test</Radio>
                <Radio value='validate'>validate</Radio>
              </HStack>
            </RadioGroup>
            <HStack>
              <Checkbox defaultChecked>Метка</Checkbox>
              <Input size='sm' placeholder='Введите метку' />
            </HStack>
            <HStack>
              <Input size='sm' placeholder='Введите название' />
              <IconButton aria-label='Добавить' size='sm' icon={<Plus size='20px' />} />
            </HStack>
            <Card w='300px' bg='gray.900'>
              <HStack justifyContent='space-between'>
                <Text>sdfsdf</Text>
                <IconButton size='xs' variant='ghost' aria-label='Удалить' icon={<X size='16px' />} />
              </HStack>
            </Card>
            <Text>Название файла:</Text>
            <Code>2023.08.13-17.46.51_train_bottle_(metka).dat</Code>
          </VStack>
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
              <Text color='gray.400'>Размер файла (МБайт):</Text><Code>100.2</Code>
            </HStack>
            <HStack>
              <Text color='gray.400'>Число пакетов:</Text><Code>103</Code>
            </HStack>
            <HStack>
              <Text color='gray.400'>Длительность записи (мм:сс):</Text><Code>03:11</Code>
            </HStack>
          </VStack>
        </HStack>
        <Box mt={2}>
          <Button w='full' p={10} colorScheme='green'>
            НАЧАТЬ ЗАПИСЬ
          </Button>
        </Box>
      </Card>
      <Heading>Список файлов</Heading>
    </Box>
  )
}
