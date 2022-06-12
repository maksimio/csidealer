import { IApi } from 'api/interfaces'
import { AlertService } from 'services/alertService'

export class LogFileController {
  constructor(
    private api: IApi,
    private alertService: AlertService
  ) { }

  async getList() {
    const list = await this.api.getLogState()
    this.alertService.success('Данные приняты')
    this.alertService.warn('Данные приняты')
    this.alertService.error('Данные приняты')
    this.alertService.info('Данные приняты')
    console.log(list)
  }
}