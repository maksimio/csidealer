import { action } from 'mobx'
import { CsiPackage, ApiService } from 'services/api'
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
      for (let i = 0; i < this.store.package.abs.length; i++) {
        for (let k = 0; k < this.store.package.abs[i].length; k++) {
          this.store.package.abs[i][k] = data.abs[i][k]
          this.store.package.phase[i][k] = data.phase[i][k]
        }
      }
    }

    for (let i = 0; i < this.store.abs.length; i++) {
      this.store.abs[i].shift()
      this.store.abs[i].push(data.abs[i][0])
    }

    for (let i = 0; i < this.store.phase.length; i++) {
      this.store.phase[i].shift()
      this.store.phase[i].push(data.phase[i][0])
    }
  })
}
