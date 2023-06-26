import { makeAutoObservable } from 'mobx'
import { CsiPackage } from 'services/api'

export const MAX_SERIES_LENGTH = 1000

export class Store {
  path: string = ''

  package?: CsiPackage
  readonly seriesX = Array(MAX_SERIES_LENGTH).fill(0).map((_, i) => i)
  seriesY: number[][] = [
    Array(MAX_SERIES_LENGTH).fill(1),
    Array(MAX_SERIES_LENGTH).fill(0),
    Array(MAX_SERIES_LENGTH).fill(0),
    Array(MAX_SERIES_LENGTH).fill(0),
  ]

  constructor() {
    makeAutoObservable(this)
  }
}