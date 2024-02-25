import { toast, Id, TypeOptions } from 'react-toastify'

export class NoticeService {
  loading(msg: string) {
    return toast.loading(msg)
  }

  finish(id: Id, msg: string, type: TypeOptions) {
    toast.update(id, { render: msg, isLoading: false, type, autoClose: 5000, closeOnClick: true, draggable: true })
  }

  info(msg: string) {
    toast.info(msg)
  }
  success(msg: string) {
    toast.success(msg)
  }
  warn(msg: string) {
    toast.warn(msg)
  }
  error(msg: string) {
    toast.error(msg)
  }
}
