import {
  DeviceController,
  WindowController,
  LogFileController,
} from 'controllers'
import DataStreamController from 'controllers/dataStreamController'
import { AlertService, ApiService } from 'services'
import { DeviceStore, WindowStore, NotificationStore } from 'stores'
import CsiStore from 'stores/csiStore'

export default class Application {
  windowStore = new WindowStore()
  deviceStore = new DeviceStore()
  notificationStore = new NotificationStore()
  csiStore = new CsiStore()

  private alertService = new AlertService()
  private apiService = new ApiService('localhost', 80, 'api/v1', 8082)

  deviceController = new DeviceController(this.apiService, this.deviceStore)
  logFileController = new LogFileController(this.apiService, this.alertService)
  windowController = new WindowController(this.windowStore)
  dataStreamController = new DataStreamController(this.apiService, this.csiStore)
}
