import { makeAutoObservable } from 'mobx'
import { CsiPackage } from 'services/api/interfaces'

export default class CsiStore {
  packages: CsiPackage[] = []
  diffs: number[][][] = []
  
  updFlag: boolean = false
  size: number = 56
  x = Array(this.size).fill(0).map((_, i) => i)

  length: number = 1000
  xtime = Array(this.length).fill(0).map((_, i) => i)
  timeseries: number[][] = [
    Array(this.length).fill(0),
    Array(this.length).fill(0),
    Array(this.length).fill(0),
    Array(this.length).fill(0),
  ]

  diffTimeseries: number[][] = [
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

    if(this.packages.length > 3000) {
      this.packages.shift()
    }

    const diff = csiPackage.data.map(arr => arr.map((v, i) => {
      if (i === arr.length - 1) {
        return 0
      }
      return arr[i + 1] - v
    }))

    this.diffs.push(diff)

    this.timeseries.forEach((ts, i) => {
      ts.shift()
      ts.push(csiPackage.data[i][0])
    })

    this.diffTimeseries.forEach((ts, i) => {
      ts.shift()
      ts.push(diff[i][0])
    })

    this.updFlag = !this.updFlag
  }
}
