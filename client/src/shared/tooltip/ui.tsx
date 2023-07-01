import { Tooltip, TooltipProps } from '@chakra-ui/react'
import { observer } from 'mobx-react-lite'
import { FC, PropsWithChildren } from 'react'
import { tooltipStore } from '.'

export const OwnTooltip: FC<PropsWithChildren & TooltipProps> = observer(({ children, ...props }) => {
  if (!tooltipStore.use) {
    return <>{children}</>
  }
  
  props.hasArrow = true
  props.openDelay = 300
  return <Tooltip {...props}>{children}</Tooltip>
})
