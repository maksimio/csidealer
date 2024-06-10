import { action, makeAutoObservable } from 'mobx'
import { CsiPackage, IRouterStatus } from 'services/api'

const MAX_SERIES_LENGTH = 1000

export interface Mark {
  id: string
  text: string
  isActive: boolean
}

export interface MarkHistory {
  mark: Mark
  time: Date
}

export enum FileType {
  Train = 'train',
  Test = 'test',
  Validate = 'validate',
}

export class Store {
  path: string = ''

  // данные csi
  package?: CsiPackage
  abs: number[][] = [
    Array(MAX_SERIES_LENGTH).fill(0),
    Array(MAX_SERIES_LENGTH).fill(0),
    Array(MAX_SERIES_LENGTH).fill(0),
    Array(MAX_SERIES_LENGTH).fill(0),
  ]

  phase: number[][] = [
    Array(MAX_SERIES_LENGTH).fill(0),
    Array(MAX_SERIES_LENGTH).fill(0),
    Array(MAX_SERIES_LENGTH).fill(0),
    Array(MAX_SERIES_LENGTH).fill(0),
  ]

  // запись в файл
  useFileType = true
  fileType: FileType = FileType.Train
  useDate = true
  private date = new Date()
  useLabel = false
  label = ''
  name = ''
  names = new Set<string>()

  limitSize = false
  sizeLimitation = 1
  limitCount = false
  countLimitation = 1000

  recording = false
  recordSize = 0 // в МБайтах
  recordCount = 0
  recordStartTimestamp = 0

  // роутеры
  routerStatus: IRouterStatus = {
    rx: { addr: '?', id: '?', is_clientmain_active: false, is_connected: false, is_sendData_active: false },
    tx: { addr: '?', id: '?', is_clientmain_active: false, is_connected: false, is_sendData_active: false },
    sendData: { dstMacAddr: '?', ifName: '?', numOfPacketToSend: 0, pktIntervalUs: 0, pktLen: 0 },
    serverIp: '?',
    serverPort: 0,
  }

  // разметка данных
  marks = new Map<string, Mark>()
  marksHistory: MarkHistory[] = []

  get filename() {
    const d = this.date
    const date = this.useDate
      ? `${d.getFullYear()}.${String(d.getMonth()).padStart(2, '0')}.${String(d.getDate()).padStart(2, '0')}-${String(
          d.getHours()
        ).padStart(2, '0')}.${String(d.getMinutes()).padStart(2, '0')}.${String(d.getSeconds()).padStart(2, '0')}_`
      : ''
    const fileType = this.useFileType ? `${this.fileType}_` : ''
    const label = this.useLabel ? `_(${this.label})` : ''

    return `${date}${fileType}${this.name}${label}.dat`
  }

  constructor() {
    makeAutoObservable(this)

    setInterval(
      action(() => {
        this.date = new Date()
      }),
      1000
    )
  }
}
