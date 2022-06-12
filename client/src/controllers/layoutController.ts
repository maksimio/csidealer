import { Location } from 'react-router-dom'
import { LayoutStore } from 'stores/layoutStore'

export class LayoutController {
  constructor(
    private LayoutStore: LayoutStore
  ) { }

  public setLocation = (location: Location) => {
    this.LayoutStore.setLocation(location)
  }
}