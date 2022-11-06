import { makeAutoObservable } from 'mobx'

export default class DeviceStore {
  tcpClientIp: string = ''

  constructor() {
    makeAutoObservable(this)
  }

  setTcpClientIp(ip: string) {
    this.tcpClientIp = ip
  }

  get isClientConnect() {
    return this.tcpClientIp !== ''
  }
}
