import { action } from 'mobx'
import { Store } from 'store'

export class NavController {
  constructor(private store: Store) { }

  setPath = action((path: string) => {
    this.store.path = path
  })
}