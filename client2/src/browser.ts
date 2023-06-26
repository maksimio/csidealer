import { createContext, useContext } from 'react'
import { Store } from 'store'

export class Browser {
  store = new Store()
}


export function getStore(b: Browser) {
  return b.store
}

export function getControllers(b: Browser) {
  return {

  }
}

export const BrowserStoreContext = createContext<ReturnType<typeof getStore> | null>(null)
export const BrowserControllersContext = createContext<ReturnType<typeof getControllers> | null>(null)

export const useStore = () => {
  const stores = useContext(BrowserStoreContext)
  if (!stores) {
    throw new Error('useStore должен использоваться внутри BrowserStoresContext')
  }
  return stores
}

export const useControllers = () => {
  const controllers = useContext(BrowserControllersContext)
  if (!controllers) {
    throw new Error('useControllers должен использоваться внутри BrowserSControllersContext')
  }
  return controllers
}

