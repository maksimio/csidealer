import { Box } from '@chakra-ui/react'
import { useApplication } from 'hooks'
import { observer } from 'mobx-react-lite'
import { FC } from 'react'

const Statusbar: FC = observer(() => {
  const { csiStore } = useApplication()
  return (
    <Box mr={10} fontSize={13} textAlign="right">
      Частота: {csiStore.frequency.toFixed(1)} Гц
    </Box>
  )
})

export default Statusbar
