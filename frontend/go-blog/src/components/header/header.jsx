import {HStack, IconButton, Menu, MenuButton, MenuItem, MenuList} from "@chakra-ui/react";
import {HamburgerIcon} from '@chakra-ui/icons'


export const Header = () => {
  return (
      <HStack justifyItems='end' width='100%' style={{position: 'relative'}}>
        <img src="/header-image.jpg" alt="Logo banner" height={250}/>
        <Menu>
          <MenuButton
              as={IconButton}
              aria-label='Options'
              icon={<HamburgerIcon/>}
              variant='outline'
              backgroundColor='#ffffff'
              style={{position: 'absolute', right: '10px', top: '10px', zIndex: '10'}}
          />
          <MenuList>
            <MenuItem>
              Blog
            </MenuItem>
            <MenuItem>
              About Me
            </MenuItem>
            <MenuItem>
              Unity / AI
            </MenuItem>
          </MenuList>
        </Menu>
      </HStack>
  );
}
