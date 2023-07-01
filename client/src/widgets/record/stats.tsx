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
  Box,
} from '@chakra-ui/react'
import { useControllers, useStore } from 'browser'
import { observer } from 'mobx-react-lite'
import { FC } from 'react'

interface LimitInputProps {
  isDisabled: boolean
  value: number
  units: string
  onChange: (limitation: string) => void
  step: number
}

const LimitInput: FC<LimitInputProps> = ({ isDisabled, value, units, onChange, step }) => {
  return (
    <HStack>
      <Text color={isDisabled ? 'gray.600' : undefined}>Не более</Text>
      <NumberInput step={step} onChange={onChange} value={value} isDisabled={isDisabled} size='sm' maxW='100px'>
        <NumberInputField />
        <NumberInputStepper>
          <NumberIncrementStepper />
          <NumberDecrementStepper />
        </NumberInputStepper>
      </NumberInput>
      <Text color={isDisabled ? 'gray.600' : undefined}>{units}</Text>
    </HStack>
  )
}

export const Stats: FC = observer(() => {
  const store = useStore()
  const { recordController } = useControllers()

  const duration = Math.round((new Date().getTime() - store.recordStartTimestamp) / 1000)
  const seconds = Math.round(duration % 60)
  const minutes = (duration - seconds) / 60
  const durationStr = `${String(minutes).padStart(2, '0')}:${String(seconds).padStart(2, '0')}`

  return (
    <VStack gap={5} alignItems='flex-start'>
      <Box>
        <Checkbox isChecked={store.limitSize} onChange={recordController.toggleLimitSize}>
          Ограничить размер файла
        </Checkbox>
        <LimitInput
          step={0.1}
          onChange={recordController.setSizeLimitation}
          units='МБайт'
          value={store.sizeLimitation}
          isDisabled={!store.limitSize}
        />
      </Box>
      <Box>
        <Checkbox isChecked={store.limitCount} onChange={recordController.toggleLimitCount}>
          Ограничить число пакетов
        </Checkbox>
        <LimitInput
          step={1}
          onChange={recordController.setCountLimitation}
          units='пакетов'
          value={store.countLimitation}
          isDisabled={!store.limitCount}
        />
      </Box>
      <HStack>
        <Text color='gray.400'>Размер файла (МБайт):</Text>
        <Code>{store.recording ? store.recordSize.toFixed(2) : '-'}</Code>
      </HStack>
      <HStack>
        <Text color='gray.400'>Число пакетов:</Text>
        <Code>{store.recording ? store.recordCount : '-'}</Code>
      </HStack>
      <HStack w='250px'>
        <Text color='gray.400'>Длительность (мм:сс):</Text>
        <Code>{store.recording ? durationStr : '-'}</Code>
      </HStack>
    </VStack>
  )
})
