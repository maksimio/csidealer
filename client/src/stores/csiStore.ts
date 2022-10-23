import { makeAutoObservable } from 'mobx'
import { CsiPackage } from 'services/api/interfaces'

export default class CsiStore {
  packages: CsiPackage[] = []
  updFlag: boolean = false
  x = Array(56).fill(0).map((_, i) => i)

  constructor() {
    makeAutoObservable(this)
  }

  push(csiPackage: CsiPackage) {
    this.packages.push(csiPackage)
    this.updFlag = !this.updFlag
  }
}
