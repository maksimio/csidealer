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
  }

  async start() {}

  async stop() {}

  async getStatus() {
    try {
      const status = await this.apiService.getRoutersStatus()
    } catch (error) {}
  }
}
