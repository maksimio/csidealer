import { FC, useEffect, useRef } from 'react'
import { WebglPlot, WebglLine } from 'webgl-plot'
import { ColorRGBA } from 'webgl-plot'

const colors = [
  {r: 192, g: 57, b: 43},
  {r: 82, g: 190, b: 128},
  {r: 84, g: 153, b: 199},
  {r: 245, g: 176, b: 65},
].map(c => new ColorRGBA(c.r / 255, c.g / 255, c.b / 255, 1))

const DELIMETER = 150
const SHIFT = -1

interface ChartProps {
  data: number[][]
  delimeter: number
  shift: number
}

export const Chart: FC<ChartProps> = ({ data, delimeter, shift }) => {
  const ref = useRef<HTMLCanvasElement>(null)

  useEffect(() => {
    if (!ref.current) {
      return
    }

    const devicePixelRatio = window.devicePixelRatio || 1
    ref.current.width = ref.current.clientWidth * devicePixelRatio
    ref.current.height = ref.current.clientHeight * devicePixelRatio

    const webglp = new WebglPlot(ref.current)
    const lines: WebglLine[] = []

    data.forEach((d, i) => {
      const line = new WebglLine(colors[i], d.length)
      line.arrangeX()
      webglp.addLine(line)
      lines.push(line)
    })

    let id = 0
    let renderPlot = () => {
      for (let k = 0; k < lines.length; k++) {
        for (let i = 0; i < data[k].length; i++) {
          lines[k].setY(i, data[k][i] / delimeter + shift)
        }
      }
      id = requestAnimationFrame(renderPlot)
      webglp.update()
    }
    id = requestAnimationFrame(renderPlot)

    return () => {
      renderPlot = () => {}
      cancelAnimationFrame(id)
    }
  }, [data.length])

  return <canvas style={{ width: '100%', height: '100%' }} ref={ref} />
}
