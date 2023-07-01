import { action } from 'mobx'
import { FileType, Store } from 'store'

export class RecordController {
  constructor(private store: Store) {}

  toggleUseFileType = action(() => {
    this.store.useFileType = !this.store.useFileType
  })

  toggleUseDate = action(() => {
    this.store.useDate = !this.store.useDate
  })

  toggleUseLabel = action(() => {
    this.store.useLabel = !this.store.useLabel
  })

  setFileType = action((fileType: FileType) => {
    this.store.fileType = fileType
  })

  setLabel = action((label: string) => {
    this.store.label = label
  })

  setName = action((name: string) => {
    this.store.name = name
  })

  removeName = action((name: string) => {
    this.store.names.delete(name)
  })

  addName = action(() => {
    if (this.store.name.length) {
      this.store.names.add(this.store.name)
    }
  })
}
