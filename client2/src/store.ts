import { makeAutoObservable } from 'mobx'

export class Store {
  path: string = ''


  constructor() {
    makeAutoObservable(this)
  }
}