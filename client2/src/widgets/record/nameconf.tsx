import {
  VStack,
  Text,
  HStack,
  RadioGroup,
  Input,
  IconButton,
  TagLabel,
  TagCloseButton,
  Checkbox,
  Radio,
  Box,
  Tag,
  Code,
} from '@chakra-ui/react'
import { useControllers, useStore } from 'browser'
import { observer } from 'mobx-react-lite'
import { FC } from 'react'
import { FileType } from 'store'
import { Plus } from 'tabler-icons-react'

interface NameTemplateProps {
  name: string
}

const NameTemplate: FC<NameTemplateProps> = ({ name }) => {
  const { recordController } = useControllers()

  return (
    <Tag
      onClick={() => recordController.setName(name)}
      mr={2}
      mb={2}
      cursor='pointer'
      _hover={{ bg: 'gray.500' }}
      transition='all linear 0.1s'
    >
      <TagLabel>{name}</TagLabel>
      <TagCloseButton onClick={() => recordController.removeName(name)} />
    </Tag>
  )
}

export const NameConf: FC = observer(() => {
  const store = useStore()
  const { recordController } = useControllers()
  const names = [...store.names.values()].map((v) => <NameTemplate name={v} />)

  return (
    <VStack alignItems='flex-start'>
      <HStack gap={5}>
        <Checkbox onChange={recordController.toggleUseFileType} isChecked={store.useFileType}>
          Тип
        </Checkbox>
        <Checkbox onChange={recordController.toggleUseDate} isChecked={store.useDate}>
          Дата
        </Checkbox>
      </HStack>
      <RadioGroup onChange={recordController.setFileType} value={store.fileType}>
        <HStack alignItems='flex-start'>
          <Radio value={FileType.Train}>{FileType.Train}</Radio>
          <Radio value={FileType.Test}>{FileType.Test}</Radio>
          <Radio value={FileType.Validate}>{FileType.Validate}</Radio>
        </HStack>
      </RadioGroup>
      <HStack>
        <Checkbox onChange={recordController.toggleUseLabel} isChecked={store.useLabel}>
          Метка
        </Checkbox>
        <Input
          onChange={(e) => recordController.setLabel(e.target.value)}
          value={store.label}
          isDisabled={!store.useLabel}
          size='sm'
          placeholder='Введите метку'
          maxLength={10}
        />
      </HStack>
      <HStack>
        <Input
          value={store.name}
          onChange={(e) => recordController.setName(e.target.value)}
          size='sm'
          placeholder='Введите название'
          onKeyDown={(e) => {
            if (e.key === 'Enter' && !store.names.has(store.name) && store.name.length) {
              recordController.addName()
            }
          }}
        />
        <IconButton
          isDisabled={store.names.has(store.name) || !store.name.length}
          onClick={recordController.addName}
          aria-label='Добавить'
          size='sm'
          icon={<Plus size='20px' />}
        />
      </HStack>
      <Box maxW='350px'>{names}</Box>
      <Text>Название файла:</Text>
      <Code>{store.filename}</Code>
    </VStack>
  )
})
