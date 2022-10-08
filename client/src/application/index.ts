import { DeviceController } from 'controllers/deviceController'
import { LayoutController } from 'controllers/layoutController'
import { LogFileController } from 'controllers/logFileController'
import { AlertService, ApiService } from 'services'
import { DeviceStore } from 'stores/deviceStore'
import { LayoutStore } from 'stores/layoutStore'
import { NotificationStore } from 'stores/notificationStore'

export default class Application {
  alertService = new AlertService()
  apiService = new ApiService('localhost', 80, 'api/v1')

  layoutStore = new LayoutStore()
  deviceStore = new DeviceStore()
  notificationStore = new NotificationStore()

  deviceController = new DeviceController(this.apiService, this.deviceStore)
  logFileController = new LogFileController(this.apiService, this.alertService)
  layoutController = new LayoutController(this.layoutStore)
}
