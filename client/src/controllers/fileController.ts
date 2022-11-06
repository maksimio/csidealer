import { IApiService, AlertService } from 'services'
import { FileStore } from 'stores'

export default class FileController {
  constructor(private apiService: IApiService, private alertService: AlertService, private fileStore: FileStore) {
    setInterval(() => {
      this.updateLogState()
    }, 1000)
  }

  private async updateLogState() {
    const res = await this.apiService.getLogState()
    if (res.success) {
      this.fileStore.setLogStates(res.result.is_open, res.result.package_count, res.result.start_ts, res.result.write_byte_count)
    }
  }

  private async startLog() {
    const res = await this.apiService.logStart(this.fileStore.filename)
    if (res.success) {
      this.alertService.success('Запись в файл начата')
    } else {
      this.alertService.error(`Ошибка начала записи: ${res.message}`)
    }
  }

  private async stopLog() {
    const res = await this.apiService.logStop()
    if (res.success) {
      this.alertService.success('Запись в файл остановлена')
    } else {
      this.alertService.error(`Ошибка остановки записи: ${res.message}`)
    }
  }

  toggleLog = async () => {
    if (this.fileStore.isLogging) {
      this.stopLog()
    } else {
      this.startLog()
    }
  }

  setFilename(filename: string) {
    this.fileStore.setFilename(filename)
  }
}
