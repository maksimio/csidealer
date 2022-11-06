import { makeAutoObservable } from 'mobx'

export default class FileStore {
  isLogging: boolean = false
  packageCount: number = 0
  startDate:  Date = new Date(0)
  byteCount: number = 0

  filename: string = ''

  constructor() {
    makeAutoObservable(this)
  }

  setLogStates(isLogging: boolean, packageCount: number, startTs: number, byteCount: number) {
    this.isLogging = isLogging
    this.packageCount = packageCount
    this.startDate = new Date(startTs)
    this.byteCount = byteCount
  }

  setFilename(filename: string) {
    this.filename = filename
  }
}
