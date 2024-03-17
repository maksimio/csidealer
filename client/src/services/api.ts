import axios, { AxiosInstance } from 'axios'
import { EventEmitter } from 'events'
import { Mark } from 'store'

export interface ErrorResponse {
  success: false
  message: string
}

interface SuccessResponse {
  success: true
}

interface SuccessResponseWithResult<T> {
  success: true
  result: T
}

type ResponseWithResult<T> = SuccessResponseWithResult<T> | ErrorResponse
export type StatusResponse = SuccessResponse | ErrorResponse

// ---------  Логирование
interface ILogState {
  start_ts: number
  is_open: boolean
  write_byte_count: number
  package_count: number
}

export type WriteStatus = ResponseWithResult<ILogState>

interface IMarkState {
  timestamp: number
  packet_num: number
}

export type MarkStatus = ResponseWithResult<IMarkState>

// -------- Устройства
export interface RouterInfo {
  id: string
  addr: string
  is_connected: boolean
  is_clientmain_active: boolean
  is_sendData_active: boolean
}

export interface IRouterStatus {
  rx: RouterInfo
  tx: RouterInfo
  sendData: {
    ifName: string
    dstMacAddr: string
    numOfPacketToSend: number
    pktIntervalUs: number
    pktLen: number // TODO: переименовать в payloadLen
  }
  serverIp: string
  serverPort: number
}

export type RouterStatus = SuccessResponseWithResult<IRouterStatus>

// -------- WebSocket
interface CsiInfo {
  bwidth: number
  csilen: number
  err: number
  nc: number
  noise: number
  nr: number
  ntones: number
  payloadlen: number
  rate: number
  rssi0: number
  rssi1: number
  rssi2: number
  rssi3: number
  ts: number
  txchan: number
}

export interface CsiPackage {
  abs: number[][]
  phase: number[][]
  id: string
  n: number
  ts: number
  info: CsiInfo
}

export type TcpClientIp = SuccessResponseWithResult<string>

const EVENT_WS_DATA = 'ws.data'

export class ApiService {
  private readonly baseUrl: string
  private instance: AxiosInstance
  private ws: WebSocket
  private eventEmitter = new EventEmitter()

  constructor(domen: string, port: number, address: string, wsPort: number) {
    this.baseUrl = `http://${domen}:${port}/${address}`
    this.instance = axios.create({ baseURL: this.baseUrl })

    axios.interceptors.response.use(
      function (response) {
        return response
      },
      function (error) {
        return Promise.reject(error)
      }
    )

    this.ws = new WebSocket(`ws://${domen}:${wsPort}`)
    this.ws.onmessage = (event: MessageEvent<string>) => {
      const data: CsiPackage = JSON.parse(event.data)
      this.eventEmitter.emit(EVENT_WS_DATA, data)
    }
  }

  onWsData(cl: (data: CsiPackage) => void) {
    this.eventEmitter.on(EVENT_WS_DATA, cl)
  }

  async writeStart(filename: string): Promise<StatusResponse> {
    try {
      const response = await this.instance.get<StatusResponse>('/write/start', { params: { filepath: filename } })
      return response.data
    } catch (e) {
      return { success: false, message: 'неизвестная ошибка' } // TODO научиться работать с AXIOS
    }
  }

  async writeStop(): Promise<StatusResponse> {
    try {
      const response = await this.instance.get<StatusResponse>('/write/stop')
      return response.data
    } catch (e) {
      return { success: false, message: 'неизвестная ошибка' }
    }
  }

  async getWriteStatus() {
    const response = await this.instance.get<WriteStatus>('/write/status')
    return response.data
  }

  async setMark(mark: Mark) {
    const response = await this.instance.get<MarkStatus>('/write/mark', {
      params: { id: mark.id, text: mark.text, is_active: mark.isActive },
    }) // разворачиваем observable-объект
    return response.data
  }

  // ------------------------------------------------- ЗАПУСК И ОСТАНОВКА ПОТОКА ДАННЫХ
  async reconnectRouters() {
    const response = await this.instance.post<StatusResponse>('/routers/reconnect')
    return response.data
  }

  async startCsiTransmit() {
    const response = await this.instance.post<StatusResponse>('/routers/start')
    return response.data
  }

  async stopCsiTransmit() {
    const response = await this.instance.post<StatusResponse>('/routers/stop')
    return response.data
  }

  async getRoutersStatus() {
    const response = await this.instance.get<RouterStatus>('/routers/status')
    return response.data
  }
}
