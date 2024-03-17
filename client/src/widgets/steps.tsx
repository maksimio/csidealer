import { FC } from 'react'
import {
  Box,
  Step,
  StepDescription,
  StepIcon,
  StepIndicator,
  StepNumber,
  StepSeparator,
  StepStatus,
  StepTitle,
  Stepper,
  useSteps,
} from '@chakra-ui/react'

const steps = [
  { title: 'Подключите', description: 'оба маршрутизатора' },
  { title: 'Добавьте', description: 'объекты для разметки' },
  { title: 'Начните', description: 'запись данных CSI в файл' },
  { title: 'Переключайте', description: 'объекты во время записи' },
  { title: 'Остановите', description: 'запись данных' },
]

// https://chakra-ui.com/docs/components/stepper/usage

export const Steps: FC = () => {
  const { activeStep } = useSteps({ index: 1, count: steps.length })

  return (
    <Stepper index={activeStep}>
      {steps.map((step, index) => (
        <Step key={index}>
          <StepIndicator>
            <StepStatus complete={<StepIcon />} incomplete={<StepNumber />} active={<StepNumber />} />
          </StepIndicator>

          <Box flexShrink='0'>
            <StepTitle>{step.title}</StepTitle>
            <StepDescription>{step.description}</StepDescription>
          </Box>

          <StepSeparator />
        </Step>
      ))}
    </Stepper>
  )
}
