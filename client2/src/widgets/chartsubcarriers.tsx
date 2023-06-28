import { useStore } from 'browser'
import { observer } from 'mobx-react-lite'
import { FC, useEffect, useRef } from 'react'
import { colors } from 'shared/chart'
import { WebglPlot, WebglLine } from 'webgl-plot'

let webglp: WebglPlot
let lines: WebglLine[] = []

export const ChartSubcarriers: FC = observer(() => {
  const canvasRef = useRef<HTMLCanvasElement>(null)
  const store = useStore()

  useEffect(() => {
    if (canvasRef.current) {
      const devicePixelRatio = window.devicePixelRatio || 1
      canvasRef.current.width = canvasRef.current.clientWidth * devicePixelRatio
      canvasRef.current.height = canvasRef.current.clientHeight * devicePixelRatio

      webglp = new WebglPlot(canvasRef.current)
      lines = []
      ;[0, 1, 2, 3].forEach((_, i) => {
        const line = new WebglLine(colors[i], 56)
        line.arrangeX()
        webglp.addLine(line)
        lines.push(line)
      })
    }
  }, [store.package])

  useEffect(() => {
    let id = 0 + 1000000
    let renderPlot = () => {
      if (store.package === undefined) {
        id = requestAnimationFrame(renderPlot) + 1000000
        webglp.update()
        return
      }
      for (let k = 0; k < lines.length; k++) {
        for (let i = 0; i < 56; i++) {
          lines[k].setY(i, store.package.data[k][i] / 150 - 1)
        }
      }
      id = requestAnimationFrame(renderPlot) + 1000000
      webglp.update()
    }
    id = requestAnimationFrame(renderPlot) + 1000000
    return () => {
      renderPlot = () => {}
      cancelAnimationFrame(id)
    }
  }, [store.package])

  return <canvas style={{ width: '100%', height: '100%' }} ref={canvasRef} />
})
