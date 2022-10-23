import axios, { AxiosInstance } from 'axios'
import { CsiPackage, IApiService, LogState, StatusResponse } from './interfaces'

export default class ApiService implements IApiService {
  private readonly baseUrl: string
  private instance: AxiosInstance
  private ws: WebSocket

  constructor(domen: string, port: number, address: string) {
    this.baseUrl = `http://${domen}:${port}/${address}`
    this.instance = axios.create({ baseURL: this.baseUrl })

    this.ws = new WebSocket(`ws://${domen}:8082`)
    this.ws.onmessage = (event: MessageEvent<CsiPackage>) => {
      console.log('Данные получены:', event.data)
    }
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
