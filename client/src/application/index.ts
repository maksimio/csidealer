import {
  DeviceController,
  WindowController,
  LogFileController,
} from 'controllers'
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
  windowController = new WindowController(this.windowStore)
}
