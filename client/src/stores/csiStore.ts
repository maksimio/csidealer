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
    // const lastPackages = this.packages.slice(-10)
    // csiPackage.data.forEach((h, i) => {
    //   for (let j = 0; j < this.size; j++) {
    //     let sum = 0
    //     for (let k = 0; k < lastPackages.length; k++) {
    //       sum += lastPackages[k].data[i][j]
    //     }
    //     sum += csiPackage.data[i][j]
    //     h[j] = sum / 11
    //   }
    // })

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

  get frequency(): number {
    const n = 500
    const tail = this.packages.slice(-n)
    const first = tail.at(0)
    const last = tail.at(-1)
    if (first === undefined || last === undefined) {
      return 0
    }
    // return (last.info.ts - first.info.ts) / n
    return n * 1000 / (last.ts - first.ts)
  }
}
