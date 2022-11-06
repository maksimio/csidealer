import {
  DeviceController,
  WindowController,
  fileController,
} from 'controllers'
import DataStreamController from 'controllers/dataStreamController'
import { AlertService, ApiService } from 'services'
import { DeviceStore, WindowStore, NotificationStore, CsiStore, FileStore } from 'stores'

export default class Application {
  windowStore = new WindowStore()
  deviceStore = new DeviceStore()
  notificationStore = new NotificationStore()
  csiStore = new CsiStore()
  fileStore = new FileStore()

  private alertService = new AlertService()
  private apiService = new ApiService('localhost', 80, 'api/v1', 8082)

  deviceController = new DeviceController(this.apiService, this.deviceStore)
  fileController = new fileController(this.apiService, this.alertService, this.fileStore)
  windowController = new WindowController(this.windowStore)
  dataStreamController = new DataStreamController(this.apiService, this.csiStore)
}
