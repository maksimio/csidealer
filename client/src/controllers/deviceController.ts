import { IApiService } from 'services'
import { DeviceStore } from 'stores'

export class DeviceController {
  constructor(
    private api: IApiService,
    private deviceStore: DeviceStore
  ) { }
}