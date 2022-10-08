import axios, { AxiosInstance } from 'axios'
import { IApiService, LogState, StatusResponse } from './interfaces'

export default class ApiService implements IApiService {
  private readonly baseUrl: string
  private instance: AxiosInstance
  private ws: WebSocket

  constructor(domen: string, port: number, address: string, wsPort: number) {
    this.baseUrl = `http://${domen}:${port}/${address}`
    this.instance = axios.create({ baseURL: this.baseUrl })

    this.ws = new WebSocket(`ws://${domen}:${wsPort}`)
    this.ws.onmessage = function (event) {
      console.log('Данные получены:', JSON.parse(event.data))
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
