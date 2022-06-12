import { FC, ReactNode } from 'react'
import {
  IconButton,
  Box,
  CloseButton,
  Flex,
  Icon,
  useColorModeValue,
  Link,
  Drawer,
  DrawerContent,
  Text,
  useDisclosure,
  BoxProps,
  FlexProps,
} from '@chakra-ui/react'
import { VscMenu } from 'react-icons/vsc'
import { IconType } from 'react-icons'
import { Link as RouterLink } from 'react-router-dom'
import { useApplication } from 'hooks'
import { observer } from 'mobx-react-lite'

const SimpleSidebar: FC<{ children: ReactNode }> = ({ children }) => {
  const { isOpen, onOpen, onClose } = useDisclosure()
  return (
    <Box minH='100vh' bg={useColorModeValue('gray.100', 'gray.900')}>
      <SidebarContent
        onClose={() => onClose}
        display={{ base: 'none', md: 'block' }}
      />
      <Drawer
        autoFocus={false}
        isOpen={isOpen}
        placement='left'
        onClose={onClose}
        returnFocusOnClose={false}
        onOverlayClick={onClose}
        size='full'>
        <DrawerContent>
          <SidebarContent onClose={onClose} />
        </DrawerContent>
      </Drawer>
      {/* mobilenav */}
      <MobileNav display={{ base: 'flex', md: 'none' }} onOpen={onOpen} />
      <Box ml={{ base: 0, md: 60 }} p='4'>
        {children}
      </Box>
    </Box>
  )
}

export default SimpleSidebar

interface SidebarProps extends BoxProps {
  onClose: () => void
}

const SidebarContent = ({ onClose, ...rest }: SidebarProps) => {
  const { layoutStore } = useApplication()
  const links = layoutStore.linkItems.map((link) => (
    <NavItem key={link.name} icon={link.icon} path={link.path}>
      {link.name}
    </NavItem>
  ))

  return (
    <Box
      bg={useColorModeValue('white', 'gray.900')}
      borderRight='1px'
      borderRightColor={useColorModeValue('gray.200', 'gray.700')}
      w={{ base: 'full', md: 60 }}
      pos='fixed'
      h='full'
      {...rest}>
      <Flex h='20' alignItems='center' mx='8' justifyContent='space-between'>
        <Text fontSize='2xl' fontFamily='monospace' fontWeight='bold'>
          CSI Dealer
        </Text>
        <CloseButton display={{ base: 'flex', md: 'none' }} onClick={onClose} />
      </Flex>
      {links}
    </Box>
  )
}

interface NavItemProps extends FlexProps {
  icon: IconType
  path: string
  children: string | number
}
const NavItem = observer(({ icon, path, children, ...rest }: NavItemProps) => {
  const { layoutStore } = useApplication()
  const isActive = layoutStore.path === path

  return (
    <Link as={RouterLink} style={{ textDecoration: 'none' }} _focus={{ boxShadow: 'none' }} to={path}>
      <Flex
        align='center'
        p='4'
        mx='4'
        borderRadius='lg'
        role='group'
        cursor='pointer'
        bg={isActive ? 'cyan.400' : ''}
        color={isActive ? 'white' : ''}
        _hover={{ bg: 'cyan.400', color: 'white' }}
        {...rest}>
        {icon && (
          <Icon mr='4' fontSize='22' _groupHover={{ color: 'white' }} as={icon} />
        )}
        {children}
      </Flex>
    </Link>
  )
})

interface MobileProps extends FlexProps {
  onOpen: () => void
}

const MobileNav = ({ onOpen, ...rest }: MobileProps) => {
  return (
    <Flex
      ml={{ base: 0, md: 60 }}
      px={{ base: 4, md: 24 }}
      height='20'
      alignItems='center'
      bg={useColorModeValue('white', 'gray.900')}
      borderBottomWidth='1px'
      borderBottomColor={useColorModeValue('gray.200', 'gray.700')}
      justifyContent='flex-start'
      {...rest}>
      <IconButton
        variant='outline'
        onClick={onOpen}
        aria-label='open menu'
        icon={<VscMenu />}
      />

      <Text fontSize='2xl' ml='8' fontFamily='monospace' fontWeight='bold'>
        CSI Dealer
      </Text>
    </Flex>
  )
}