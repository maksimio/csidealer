import { IApiService } from 'services'
import { DeviceStore } from 'stores'

export default class DeviceController {
  constructor(
    private api: IApiService,
    private deviceStore: DeviceStore
  ) { }
}