import { makeAutoObservable } from 'mobx'

export default class DeviceStore {
  constructor() {
    makeAutoObservable(this)
  }
}