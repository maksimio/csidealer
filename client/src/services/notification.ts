import { toast } from 'react-toastify'

export class NotificationService {
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
