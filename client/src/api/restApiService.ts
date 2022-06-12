import axios, { AxiosInstance } from 'axios'
import { IApi, LogState, StatusResponse } from './interfaces'

export class RestApiService implements IApi {
  private readonly baseUrl: string
  private axiosInstance: AxiosInstance

  constructor(domen: string, port: number, address: string) {
    this.baseUrl = `http://${domen}:${port}/${address}`
    this.axiosInstance = axios.create({ baseURL: this.baseUrl })
  }

  async logStart<T = StatusResponse>(filename: string): Promise<T> {
    const response = await this.axiosInstance.get<T>('/log/start',
      { params: { filepath: filename } })
    return response.data
  }

  async logStop<T = StatusResponse>(): Promise<T> {
    const response = await this.axiosInstance.get<T>('/log/stop')
    return response.data
  }

  async getLogState<T = LogState>(): Promise<T> {
    const response = await this.axiosInstance.get<T>('/log/state')
    return response.data
  }
}