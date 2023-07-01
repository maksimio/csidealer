import { makeAutoObservable } from 'mobx'
import { CsiPackage } from 'services/api'

export const MAX_SERIES_LENGTH = 1000

export enum FileType {
  Train = 'train',
  Test = 'test',
  Validate = 'validate',
}

export class Store {
  path: string = ''

  // данные csi
  package?: CsiPackage
  readonly seriesX = Array(MAX_SERIES_LENGTH)
    .fill(0)
    .map((_, i) => i)
  seriesY: number[][] = [
    Array(MAX_SERIES_LENGTH).fill(1),
    Array(MAX_SERIES_LENGTH).fill(0),
    Array(MAX_SERIES_LENGTH).fill(0),
    Array(MAX_SERIES_LENGTH).fill(0),
  ]

  // запись в файл
  useFileType = true
  fileType: FileType = FileType.Train
  useDate = true
  useLabel = false
  label = ''
  name = ''
  names = new Set<string>()

  limitSize = false
  sizeLimitation = 1
  limitCount = false
  countLimitation = 1000

  isRecord = false
  recordSize = 0
  recordCount = 0
  recordDuration = 0

  get filename() {
    const date = this.useDate ? `2023.08.13-17.46.51_` : ''
    const fileType = this.useFileType ? `${this.fileType}_` : ''
    const label = this.useLabel ? `_(${this.label})` : ''

    return `${date}${fileType}${this.name}${label}.dat`
  }

  constructor() {
    makeAutoObservable(this)
  }
}
