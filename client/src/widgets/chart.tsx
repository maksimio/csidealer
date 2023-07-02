import { FC, useEffect, useRef } from 'react'
import { colors } from 'shared/chart'
import { WebglPlot, WebglLine } from 'webgl-plot'

const DELIMETER = 150
const SHIFT = -1

interface ChartProps {
  data: number[][]
}

export const Chart: FC<ChartProps> = ({ data }) => {
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
          lines[k].setY(i, data[k][i] / DELIMETER + SHIFT)
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
