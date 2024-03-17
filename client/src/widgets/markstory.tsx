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
import { FC } from 'react'
import { Card } from 'shared/card'

interface MarkStepProps {
  text: string
  date: Date
}

const MarkStep: FC<MarkStepProps> = ({ date, text }) => {
  return (
    <Step>
      <StepIndicator>
        <StepStatus complete={<StepIcon />} incomplete={<StepNumber />} active={<StepNumber />} />
      </StepIndicator>

      <Box flexShrink='0'>
        <StepTitle>{text}</StepTitle>
        <StepDescription>{date.toLocaleTimeString()}</StepDescription>
      </Box>

      <StepSeparator />
    </Step>
  )
}

const steps = [
  { title: 'First', description: 'Contact Info' },
  { title: 'Second', description: 'Date & Time' },
  { title: 'Third', description: 'Select Rooms' },
]

export const MarkStory: FC = () => {
  const { activeStep } = useSteps({
    index: 1,
    count: steps.length,
  })

  return (
    <Card>
      <Stepper index={activeStep} orientation='vertical' gap='0'>
        {steps.map((step, index) => (
          <Step key={index + 8}>
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
    </Card>
  )
}
