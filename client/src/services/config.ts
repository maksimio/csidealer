export class ConfigService {
  set(key: string, value: any) {
    window.localStorage.setItem(key, value)
  }

  get(key: string) {
    return window.localStorage.getItem(key)
  }
}