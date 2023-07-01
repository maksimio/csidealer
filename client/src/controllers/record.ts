import { action, runInAction } from 'mobx'
import ApiService from 'services/api'
import { FileType, Store } from 'store'

export class RecordController {
  constructor(private store: Store, private apiService: ApiService) {
    setInterval(() => {
      this.updateLogState()
    }, 1000)
  }

  private updateLogState = async () => {
    const res = await this.apiService.getLogState()
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
      await this.apiService.logStop()
    }
  }

  toggleRecording = async () => {
    await this.updateLogState()
    if (this.store.recording) {
      await this.apiService.logStop()
    } else {
      await this.apiService.logStart(this.store.filename)
    }
    await this.updateLogState()
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
}
