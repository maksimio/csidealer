import { makeAutoObservable } from 'mobx'
import { CsiPackage } from 'services/api/interfaces'

export default class CsiStore {
  packages: CsiPackage[] = []
  updFlag: boolean = false
  x = Array(56).fill(0).map((_, i) => i)

  length: number = 1000
  xtime = Array(this.length).fill(0).map((_, i) => i)
  timeseries: number[][] = [
    Array(this.length).fill(0),
    Array(this.length).fill(0),
    Array(this.length).fill(0),
    Array(this.length).fill(0),
  ]

  constructor() {
    makeAutoObservable(this)
  }

  push(csiPackage: CsiPackage) {
    this.packages.push(csiPackage)

    this.timeseries[0].shift()
    this.timeseries[0].push(csiPackage.data[0][33])

    this.updFlag = !this.updFlag
  }
}
