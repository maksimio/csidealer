import axios, { AxiosInstance } from 'axios'
import EventEmitter from 'events'
import { CsiPackage, IApiService, LogState, StatusResponse } from './interfaces'

const EVENT_WS_DATA = 'ws.data'

export default class ApiService implements IApiService {
  private readonly baseUrl: string
  private instance: AxiosInstance
  private ws: WebSocket
  private eventEmitter: EventEmitter = new EventEmitter()

  constructor(domen: string, port: number, address: string, wsPort: number) {
    this.baseUrl = `http://${domen}:${port}/${address}`
    this.instance = axios.create({ baseURL: this.baseUrl })

    this.ws = new WebSocket(`ws://${domen}:${wsPort}`)
    this.ws.onmessage = (event: MessageEvent<string>) => {
      const data: CsiPackage = JSON.parse(event.data)
      this.eventEmitter.emit(EVENT_WS_DATA, data)
    }
  }

  onWsData(cl: (data: CsiPackage) => void) {
    this.eventEmitter.on(EVENT_WS_DATA, cl)
  }

  async logStart<T = StatusResponse>(filename: string): Promise<T> {
    const response = await this.instance.get<T>('/log/start', {
      params: { filepath: filename },
    })
    return response.data
  }

  async logStop<T = StatusResponse>(): Promise<T> {
    const response = await this.instance.get<T>('/log/stop')
    return response.data
  }

  async getLogState<T = LogState>(): Promise<T> {
    const response = await this.instance.get<T>('/log/state')
    return response.data
  }
}
