import {
  VStack,
  IconButton,
  Tooltip,
  Flex,
  Spacer,
  useColorMode,
} from '@chakra-ui/react'
import {
  IconReload,
  IconSettings,
  IconSquareNumber1,
  IconSquareNumber2,
  IconSquareNumber4,
  IconSquareNumber5,
  TablerIcon,
} from '@tabler/icons'
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
        <IconLink
          label="Лабораторная работа №1"
          link="/lab1"
          Icon={IconSquareNumber1}
        />
        <IconLink
          label="Лабораторная работа №2"
          link="/lab2"
          Icon={IconSquareNumber2}
        />
        <IconLink
          label="Лабораторная работа №4"
          link="/lab4"
          Icon={IconSquareNumber4}
        />
        <IconLink
          label="Лабораторная работа №5"
          link="/lab5"
          Icon={IconSquareNumber5}
        />
      </VStack>
      <Spacer />
      <VStack>
        <Tooltip
          defaultIsOpen
          label="Обновить"
          hasArrow
          placement="right"
          openDelay={500}
        >
          <IconButton
            variant="solid"
            isRound
            colorScheme="teal"
            aria-label="Обновить"
            icon={<IconReload />}
          />
        </Tooltip>
        <IconLink label="Настройки" link="/settings" Icon={IconSettings} />
      </VStack>
    </Flex>
  )
}

export default Sidebar
