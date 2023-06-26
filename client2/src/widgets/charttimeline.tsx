import { useStore } from 'browser'
import { observer } from 'mobx-react-lite'
import { FC, useEffect, useRef } from 'react'
import { colors } from 'shared/chart'
import { MAX_SERIES_LENGTH } from 'store'
// @ts-ignore
import { WebglPlot, WebglLine } from 'webgl-plot'

let webglp: WebglPlot
let lines: WebglLine[] = []

export const ChartTimeline: FC = observer(() => {
  const canvasRef = useRef<HTMLCanvasElement>(null)
  const store = useStore()

  useEffect(() => {
    if (canvasRef.current) {
      const devicePixelRatio = window.devicePixelRatio || 1
      canvasRef.current.width = canvasRef.current.clientWidth * devicePixelRatio
      canvasRef.current.height = canvasRef.current.clientHeight * devicePixelRatio

      webglp = new WebglPlot(canvasRef.current)
      lines = []

      store.seriesY.forEach((_, i) => {
        const line = new WebglLine(colors[i], MAX_SERIES_LENGTH)
        line.arrangeX()
        webglp.addLine(line)
        lines.push(line)
      })
    }
  }, [store.seriesY])

  useEffect(() => {
    let id = 0
    let renderPlot = () => {
      for (let k = 0; k < lines.length; k++) {
        for (let i = 0; i < MAX_SERIES_LENGTH; i++) {
          lines[k].setY(i, store.seriesY[k][i] / 150 - 1)
        }
      }
      id = requestAnimationFrame(renderPlot)
      webglp.update()
    }
    id = requestAnimationFrame(renderPlot)
    return () => {
      renderPlot = () => { }
      cancelAnimationFrame(id)
    }
  }, [store.seriesY])

  return <canvas style={{ width: '100%', height: '100%' }} ref={canvasRef} />
})
