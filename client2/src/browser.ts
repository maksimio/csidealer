import { ChartController } from 'controllers/chart'
import { NavController } from 'controllers/nav'
import { createContext, useContext } from 'react'
import ApiService from 'services/api'
import { Store } from 'store'

export class Browser {
  apiService = new ApiService('127.0.0.1', 80, 'api/v1', 8082)

  store = new Store()

  navController = new NavController(this.store)
  chartController = new ChartController(this.store, this.apiService)
}


export function getStore(b: Browser) {
  return b.store
}

export function getControllers(b: Browser) {
  return {
    navController: b.navController
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

