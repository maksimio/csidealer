import { makeAutoObservable } from 'mobx'
import { TablerIcon, IconBinary } from '@tabler/icons'
import { Location } from 'react-router-dom'

interface LinkItem {
  name: string
  icon: TablerIcon
  path: string
}

export class LayoutStore {
  private location: Location | undefined
  linkItems: LinkItem[] = [
    { name: 'Главная', icon: IconBinary, path: '/dashboard' },
    { name: 'Пакеты', icon: IconBinary, path: '/package' },
    { name: 'Поднесущие', icon: IconBinary, path: '/subcarrier' },
    { name: 'Устройства', icon: IconBinary, path: '/devices' },
    { name: 'Логирование', icon: IconBinary, path: '/log' },
    { name: 'Параметры', icon: IconBinary, path: '/settings' },
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