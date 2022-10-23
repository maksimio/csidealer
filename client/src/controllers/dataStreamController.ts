import { IApiService } from 'services'
import { CsiPackage } from 'services/api/interfaces'
import CsiStore from 'stores/csiStore'

export default class DataStreamController {
  constructor(private apiService: IApiService, private csiStore: CsiStore) {
    this.apiService.onWsData(this.onWsDataHandler)
  }

  private onWsDataHandler = (csiPackage: CsiPackage) => {
    this.csiStore.push(csiPackage)
  }
}
