import { runInAction } from 'mobx'
import { ApiService, NoticeService } from 'services'
import { Store } from 'store'

export class RouterController {
  constructor(private store: Store, private apiService: ApiService, private noticeService: NoticeService) {}

  reconnect = async () => {
    const id = this.noticeService.loading('Установка соединения')
    try {
      await this.apiService.reconnectRouters()
      this.noticeService.finish(id, 'Соединение с роутерами установлено', 'success')
    } catch (error) {
      this.noticeService.finish(id, 'Не удалось подключиться к роутерам', 'error')
    }

    await this.getStatus()
  }

  toggleTransmitCSI = async () => {
    const isTransmit = this.store.routerStatus.rx.is_clientmain_active && this.store.routerStatus.tx.is_sendData_active
    const id = this.noticeService.loading('Управление приемом CSI')
    try {
      if (isTransmit) {
        await this.apiService.stopCsiTransmit()
        this.noticeService.finish(id, 'Прием CSI остановлен', 'success')
      } else {
        await this.apiService.startCsiTransmit()
        this.noticeService.finish(id, 'Начат прием CSI', 'success')
      }
    } catch {
      this.noticeService.finish(id, 'Не удалось отправить команды передачи', 'error')
    }

    await this.getStatus()
  }

  getStatus = async () => {
    try {
      const status = await this.apiService.getRoutersStatus()
      if (status.success) {
        runInAction(() => (this.store.routerStatus = status.result))
      } else {
        this.noticeService.error('Что-то пошло не так на сервере при запросе информации о роутерах')
      }
    } catch (error) {
      this.noticeService.error('Нет соединения с сервером')
    }
  }
}
