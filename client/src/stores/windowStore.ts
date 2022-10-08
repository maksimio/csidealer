import { makeAutoObservable } from 'mobx'
import { Location } from 'react-router-dom'


export default class WindowStore {
  private location: Location | undefined

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