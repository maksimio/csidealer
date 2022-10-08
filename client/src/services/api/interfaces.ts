export interface IApiService {
  logStart: (filename: string) => Promise<StatusResponse>
  logStop: () => Promise<StatusResponse>
  getLogState: () => Promise<LogState>
}

interface ErrorResponse {
  success: false
  message: string
}

interface SuccessResponse {
  success: true
}

interface SuccessResponseWithResult<T> {
  success: true
  result: T
}

type ResponseWithResult<T> = SuccessResponseWithResult<T> | ErrorResponse
export type StatusResponse = SuccessResponse | ErrorResponse

// ---------  Логирование
interface ILogState {
  start_ts: number
  is_open: boolean
  write_byte_count: number
  package_count: number
}

export type LogState = ResponseWithResult<ILogState>

// -------- Устройства
interface IDeviceInfo {
  id: string,
  addr: string,
  is_connected: string
  is_clientmain_active: boolean
  is_sendData_active: boolean
  is_available: boolean
}

export type DeviceInfo = SuccessResponseWithResult<IDeviceInfo>