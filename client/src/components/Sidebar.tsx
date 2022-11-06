import { VStack, IconButton, Flex, Spacer, useColorMode } from '@chakra-ui/react'
import { IconChartLine, IconFiles, IconRouter, IconSettings, TablerIcon } from '@tabler/icons'
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
    <IconButton
      borderRadius={0}
      width="50px"
      height="60px"
      isActive={windowStore.path === link}
      variant="ghost"
      aria-label={label}
      icon={<Icon />}
      onClick={() => navigate(link)}
    />
  )
})

const Sidebar: FC = () => {
  const { colorMode } = useColorMode()
  return (
    <Flex p={2} h="full" direction="column" bg={colorMode === 'light' ? 'blue.300' : 'blue.800'}>
      <VStack spacing={0}>
        <IconLink label="Устройства" link="/devices" Icon={IconRouter} />
        <IconLink label="Графики" link="/charts" Icon={IconChartLine} />
        <IconLink label="Файлы" link="/files" Icon={IconFiles} />
      </VStack>
      <Spacer />
      <VStack>
        <IconLink label="Настройки" link="/settings" Icon={IconSettings} />
      </VStack>
    </Flex>
  )
}

export default Sidebar
