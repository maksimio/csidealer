import { makeAutoObservable } from 'mobx'

class TooltipStore {
  use: boolean = true

  constructor() {
    makeAutoObservable(this)
  }

  setUse(use: boolean) {
    this.use = use
  }
}

export const tooltipStore = new TooltipStore()