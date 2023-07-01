import { FC, PropsWithChildren } from 'react'
import {
  Browser,
  BrowserControllersContext,
  BrowserStoreContext,
  getControllers,
  getStore,
} from 'browser'

const browser = new Browser()

export const WithBrowser: FC<PropsWithChildren> = ({ children }) => {
  return (
    <BrowserStoreContext.Provider value={getStore(browser)}>
      <BrowserControllersContext.Provider value={getControllers(browser)}>{children}</BrowserControllersContext.Provider>
    </BrowserStoreContext.Provider>
  )
}
