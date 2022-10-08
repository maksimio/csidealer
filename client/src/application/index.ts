import { DeviceController } from 'controllers/deviceController'
import { LayoutController } from 'controllers/layoutController'
import { LogFileController } from 'controllers/logFileController'
import { AlertService, ApiService } from 'services'
import { DeviceStore, WindowStore, NotificationStore } from 'stores'

export default class Application {
  private alertService = new AlertService()
  private apiService = new ApiService('localhost', 80, 'api/v1')

  windowStore = new WindowStore()
  deviceStore = new DeviceStore()
  notificationStore = new NotificationStore()

  deviceController = new DeviceController(this.apiService, this.deviceStore)
  logFileController = new LogFileController(this.apiService, this.alertService)
  layoutController = new LayoutController(this.windowStore)
}
