import { action, runInAction } from 'mobx'
import { ApiService } from 'services'
import { FileType, Store } from 'store'
import { v4 } from 'uuid'

export class RecordController {
  constructor(private store: Store, private apiService: ApiService) {}

  updateWriteStatus = async () => {
    const res = await this.apiService.getWriteStatus()
    if (!res.success) {
      return
    }

    runInAction(() => {
      this.store.recording = res.result.is_open
      this.store.recordSize = res.result.write_byte_count / 1048576
      this.store.recordCount = res.result.package_count
      this.store.recordStartTimestamp = res.result.start_ts
    })

    if (!this.store.recording) {
      return
    }

    if (
      (this.store.limitCount && this.store.recordCount > this.store.countLimitation) ||
      (this.store.limitSize && this.store.recordSize > this.store.sizeLimitation)
    ) {
      await this.apiService.writeStop()
    }
  }

  toggleRecording = async () => {
    await this.updateWriteStatus()
    if (this.store.recording) {
      await this.apiService.writeStop()
    } else {
      await this.apiService.writeStart(this.store.filename)
    }
    await this.updateWriteStatus()
  }

  toggleUseFileType = action(() => {
    this.store.useFileType = !this.store.useFileType
  })

  toggleUseDate = action(() => {
    this.store.useDate = !this.store.useDate
  })

  toggleUseLabel = action(() => {
    this.store.useLabel = !this.store.useLabel
  })

  toggleLimitCount = action(() => {
    this.store.limitCount = !this.store.limitCount
  })

  toggleLimitSize = action(() => {
    this.store.limitSize = !this.store.limitSize
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

  setSizeLimitation = action((limitation: string) => {
    const l = Number(limitation)
    if (!isNaN(l)) {
      this.store.sizeLimitation = Number(limitation)
    }
  })

  setCountLimitation = action((limitation: string) => {
    const l = Number(limitation)
    if (!isNaN(l)) {
      this.store.countLimitation = l
    }
  })

  removeName = action((name: string) => {
    this.store.names.delete(name)
  })

  addName = action(() => {
    if (this.store.name.length) {
      this.store.names.add(this.store.name)
    }
  })

  // разметка
  toggleActiveMark = action((id: string) => {
    const mark = this.store.marks.get(id)
    if (!mark) {
      return
    }

    mark.isActive = !mark.isActive
  })

  unactiveMarks = action(() => {
    for (let v of this.store.marks.values()) {
      v.isActive = false
    }
  })

  addMark = action(() => {
    const id = v4()
    this.store.marks.set(id, { id, isActive: false, text: new Date().toLocaleTimeString() })
  })

  clearMarks = action(() => {
    this.store.marks.clear()
  })

  deleteMark = action((id: string) => {
    this.store.marks.delete(id)
  })
}
