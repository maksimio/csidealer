export class Device {
  constructor(
    readonly id: string,
    readonly addr: string,
    public connected: boolean,
    public clientmain: boolean,
    public sendData: boolean,
    public available: boolean
  ) { }
}