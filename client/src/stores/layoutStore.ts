import { makeAutoObservable } from 'mobx'
import { IconType } from 'react-icons'
import { Location } from 'react-router-dom'
import { VscSettings, VscFolderLibrary, VscRadioTower, VscDashboard } from 'react-icons/vsc'

interface LinkItem {
  name: string
  icon: IconType
  path: string
}

export class LayoutStore {
  private location: Location | undefined
  linkItems: LinkItem[] = [
    { name: 'Главная', icon: VscDashboard, path: '/dashboard' },
    { name: 'Устройства', icon: VscRadioTower, path: '/devices' },
    { name: 'Логгирование', icon: VscFolderLibrary, path: '/log' },
    { name: 'Параметры', icon: VscSettings, path: '/settings' },
  ]

  constructor() {
    makeAutoObservable(this)
  }

  setLocation(location: Location): void {
    this.location = location
  }

  get path(): string {
    if (this.location === undefined) {
      return '/'
    }

    return this.location.pathname
  }
}