import axios, { AxiosInstance } from 'axios'
import EventEmitter from 'events'
import { CsiPackage, IApiService, LogState, StatusResponse, TcpClientIp } from './interfaces'

const EVENT_WS_DATA = 'ws.data'

export default class ApiService implements IApiService {
  private readonly baseUrl: string
  private instance: AxiosInstance
  private ws: WebSocket
  private eventEmitter: EventEmitter = new EventEmitter()

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
