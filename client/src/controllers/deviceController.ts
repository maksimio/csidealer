import { IApi } from 'api/interfaces'
import { DeviceStore } from 'stores/deviceStore'

export class DeviceController {
  constructor(
    private api: IApi,
    private deviceStore: DeviceStore
  ) { }
}