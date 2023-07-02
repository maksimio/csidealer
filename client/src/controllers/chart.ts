import { action } from 'mobx'
import ApiService, { CsiPackage } from 'services/api'
import { Store } from 'store'

export class ChartController {
  constructor(private store: Store, private apiService: ApiService) {
    this.apiService.onWsData(this.handleCsiData)
  }

  private handleCsiData = action((data: CsiPackage) => {
    if (this.store.package === undefined) {
      this.store.package = data
    } else {
      this.store.package.id = data.id
      this.store.package.info = data.info
      this.store.package.n = data.n
      this.store.package.ts = data.ts
      // Для работы графиков WebGL необходимо изменение существующего массива
      for (let i = 0; i < this.store.package.data.length; i++) {
        for (let k = 0; k < this.store.package.data[i].length; k++) {
          this.store.package.data[i][k] = data.data[i][k]
        }
      }
    }

    // this.store.package.data.forEach((s, i) => {
    //   data
    // })

    for (let i = 0; i < this.store.seriesY.length; i++) {
      this.store.seriesY[i].shift()
      this.store.seriesY[i].push(data.data[i][0])
    }
  })
}
