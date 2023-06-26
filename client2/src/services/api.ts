import axios, { AxiosInstance } from 'axios'
import { EventEmitter } from 'events'

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

export type LogState = ResponseWithResult<ILogState>

// -------- Устройства
interface IDeviceInfo {
  id: string
  addr: string
  is_connected: string
  is_clientmain_active: boolean
  is_sendData_active: boolean
  is_available: boolean
}

export type DeviceInfo = SuccessResponseWithResult<IDeviceInfo>

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
  data: number[][]
  id: string
  n: number
  ts: number
  info: CsiInfo
}

export type TcpClientIp = SuccessResponseWithResult<string>


const EVENT_WS_DATA = 'ws.data'

export default class ApiService {
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

  async logStart(filename: string): Promise<StatusResponse> {
    try {
      const response = await this.instance.get<StatusResponse>('/log/start', { params: { filepath: filename } })
      return response.data
    } catch (e) {
      return { success: false, message: 'неизвестная ошибка' } // TODO научиться работать с AXIOS
    }
  }

  async logStop(): Promise<StatusResponse> {
    try {
      const response = await this.instance.get<StatusResponse>('/log/stop')
      return response.data
    } catch (e) {
      debugger
      return { success: false, message: 'неизвестная ошибка' }
    }
  }

  async getLogState<T = LogState>(): Promise<T> {
    const response = await this.instance.get<T>('/log/state')
    return response.data
  }

  async getTcpClientIp<T = TcpClientIp>(): Promise<T> {
    const response = await this.instance.get<T>('/devices/tcp_client_ip')
    return response.data
  }
}
