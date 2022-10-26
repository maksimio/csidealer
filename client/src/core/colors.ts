import { ColorRGBA } from 'webgl-plot'

export const colors = [
  {r: 192, g: 57, b: 43},
  {r: 82, g: 190, b: 128},
  {r: 84, g: 153, b: 199},
  {r: 245, g: 176, b: 65},
].map(c => new ColorRGBA(c.r / 255, c.g / 255, c.b / 255, 1))

