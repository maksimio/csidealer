import { Location } from 'react-router-dom'
import { WindowStore } from 'stores'

export class LayoutController {
  constructor(
    private windowStore: WindowStore
  ) { }

  public setLocation = (location: Location) => {
    this.windowStore.setLocation(location)
  }
}