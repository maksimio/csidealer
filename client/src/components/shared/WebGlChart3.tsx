import { colors } from 'core/colors'
import { useApplication } from 'hooks'
import { observer } from 'mobx-react-lite'
import { FC, useEffect, useRef } from 'react'
import { WebglPlot, WebglLine } from 'webgl-plot'

let webglp: WebglPlot
let lines: WebglLine[] = []

const WebGlChart3: FC = observer(() => {
  const canvasRef = useRef<HTMLCanvasElement>(null)
  const { csiStore } = useApplication()

  useEffect(() => {
    if (canvasRef.current) {
      const devicePixelRatio = window.devicePixelRatio || 1
      canvasRef.current.width = canvasRef.current.clientWidth * devicePixelRatio
      canvasRef.current.height = canvasRef.current.clientHeight * devicePixelRatio

      webglp = new WebglPlot(canvasRef.current)
      lines = []
      const lastDiff = csiStore.diffs.at(-1)
      if (lastDiff === undefined) {
        return
      }

      lastDiff.forEach((_, i) => {
        const line = new WebglLine(colors[i], csiStore.size)
        line.arrangeX()
        webglp.addLine(line)
        lines.push(line)
      })
    }
  }, [csiStore.size, csiStore.diffs])

  useEffect(() => {
    let id = 0 + 1000000000
    let renderPlot = () => {
      const lastDiff = csiStore.diffs.at(-1) 
      if (lastDiff === undefined) {
        return
      }

      for (let k = 0; k < lines.length; k++) {
        for (let i = 0; i < csiStore.size; i++) {
          lines[k].setY(i, lastDiff[k][i] / 15)
        }
      }
      id = requestAnimationFrame(renderPlot) + 1000000000
      webglp.update()
    }
    id = requestAnimationFrame(renderPlot) + 1000000000
    return () => {
      renderPlot = () => {}
      cancelAnimationFrame(id)
    }
  }, [csiStore.size, csiStore.diffs])

  return <canvas style={{ width: '100%', height: '100%' }} ref={canvasRef} />
})

export default WebGlChart3
