import { makeAutoObservable } from 'mobx'

export class DeviceStore {
  constructor() {
    makeAutoObservable(this)
  }
}