import { useApplication } from 'hooks'
import { observer } from 'mobx-react-lite'
import { FC, useEffect, useRef } from 'react'
import { WebglPlot, WebglLine, ColorRGBA } from 'webgl-plot'

let webglp: WebglPlot
let lines: WebglLine[] = []

const WebGlChart4: FC = observer(() => {
  const canvasRef = useRef<HTMLCanvasElement>(null)
  const { csiStore } = useApplication()

  useEffect(() => {
    if (canvasRef.current) {
      const devicePixelRatio = window.devicePixelRatio || 1
      canvasRef.current.width = canvasRef.current.clientWidth * devicePixelRatio
      canvasRef.current.height = canvasRef.current.clientHeight * devicePixelRatio

      webglp = new WebglPlot(canvasRef.current)
      lines = []

      const colors = [
        new ColorRGBA(1, 0, 0, 1),
        new ColorRGBA(0, 1, 0, 1),
        new ColorRGBA(0, 0, 1, 1),
        new ColorRGBA(0, 0, 0, 1),
      ]

      csiStore.diffTimeseries.forEach((_, i) => {
        const line = new WebglLine(colors[i], csiStore.length)
        line.arrangeX()
        webglp.addLine(line)
        lines.push(line)
      })
    }
  }, [csiStore.length, csiStore.diffTimeseries])

  useEffect(() => {
    let id = 1000
    let renderPlot = () => {
      for (let k = 0; k < lines.length; k++) {
        for (let i = 0; i < csiStore.length; i++) {
          lines[k].setY(i, csiStore.diffTimeseries[k][i] / 10)
        }
      }
      id = requestAnimationFrame(renderPlot) + 1000
      webglp.update()
    }
    id = requestAnimationFrame(renderPlot) + 1000
    return () => {
      renderPlot = () => {}
      cancelAnimationFrame(id)
    }
  }, [csiStore.diffTimeseries, csiStore.length])

  return <canvas style={{ width: '100%', height: '100%' }} ref={canvasRef} />
})

export default WebGlChart4
