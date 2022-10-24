import {
  VStack,
  IconButton,
  Tooltip,
  Flex,
  Spacer,
  useColorMode,
} from '@chakra-ui/react'
import { Icon3dRotate, IconSettings, IconWaveSawTool, TablerIcon } from '@tabler/icons'
import { observer } from 'mobx-react-lite'
import { FC } from 'react'
import { useNavigate } from 'react-router-dom'
import { useApplication } from 'hooks'

interface IconLinkProps {
  label: string
  link: string
  Icon: TablerIcon
}

const IconLink: FC<IconLinkProps> = observer(({ label, link, Icon }) => {
  const { windowStore } = useApplication()
  const navigate = useNavigate()

  return (
    <Tooltip
      defaultIsOpen
      label={label}
      hasArrow
      placement="right"
      openDelay={500}
    >
      <IconButton
        isActive={windowStore.path === link}
        variant="ghost"
        aria-label={label}
        icon={<Icon />}
        onClick={() => navigate(link)}
      />
    </Tooltip>
  )
})

const Sidebar: FC = () => {
  const { colorMode } = useColorMode()
  return (
    <Flex
      p={3}
      h="full"
      direction="column"
      background={colorMode === 'light' ? 'gray.50' : 'gray.700'}
    >
      <VStack>
        <IconLink label="Дашбоард" link="/dashboard" Icon={IconWaveSawTool} />
        <IconLink label="Дашбоард" link="/devices" Icon={Icon3dRotate} />
      </VStack>
      <Spacer />
      <VStack>
        <IconLink label="Настройки" link="/settings" Icon={IconSettings} />
      </VStack>
    </Flex>
  )
}

export default Sidebar
