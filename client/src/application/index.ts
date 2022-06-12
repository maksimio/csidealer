import { DeviceController } from 'controllers/deviceController'
import { LayoutController } from 'controllers/layoutController'
import { LogFileController } from 'controllers/logFileController'
import { DeviceStore } from 'stores/deviceStore'
import { LayoutStore } from 'stores/layoutStore'
import { RestApiService } from '../api'

export default class Application {
    apiService = new RestApiService('localhost', 80, 'api/v1')

    layoutStore = new LayoutStore()
    deviceStore = new DeviceStore()

    deviceController = new DeviceController(this.apiService, this.deviceStore)
    logFileController = new LogFileController(this.apiService)
    layoutController = new LayoutController(this.layoutStore)
}