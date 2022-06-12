import { IApi } from 'api/interfaces'

export class LogFileController {
  constructor(
    private api: IApi
  ) { }

  async getList() {
    const list = await this.api.getLogState()
    console.log(list)
  }
}