import { action } from 'mobx'
import ApiService, { CsiPackage } from 'services/api'
import { Store } from 'store'

export class ChartController {
  constructor(private store: Store, private apiService: ApiService) {
    this.apiService.onWsData(this.handleCsiData)
  }

  private handleCsiData = action((data: CsiPackage) => {
    this.store.package = data

    this.store.seriesY.forEach((s, i) => {
      s.shift()
      s.push(data.data[i][0])
    })
  })
}