import { IApiService } from 'services'
import { DeviceStore } from 'stores'

export default class DeviceController {
  constructor(
    private apiService: IApiService,
    private deviceStore: DeviceStore
  ) {
    setInterval(() => {
      this.updateTcpClientIp()
    }, 1000)
   }

  private async updateTcpClientIp() {
    const res = await this.apiService.getTcpClientIp()
    if (res.success) {
      this.deviceStore.setTcpClientIp(res.result)
    }
  }

}