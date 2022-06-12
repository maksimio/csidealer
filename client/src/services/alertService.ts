import { toast, ToastOptions } from 'react-toastify'

export class AlertService {
  private config: ToastOptions = {
    position: 'bottom-center',
    autoClose: 5000,
    hideProgressBar: false,
    closeOnClick: true,
    pauseOnHover: true,
    draggable: true,
    progress: undefined,
  }

  success(message: string) {
    toast.success(message, this.config)
  }

  warn(message: string) {
    toast.warn(message, this.config)
  }

  error(message: string) {
    toast.error(message, this.config)
  }

  info(message: string) {
    toast.info(message, this.config)
  }

  // TODO написать emit и on, чтобы связать со стором сохранения
} 