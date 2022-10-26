import { useApplication } from 'hooks'
import { observer } from 'mobx-react-lite'
import { FC, useEffect, useRef } from 'react'
import { WebglPlot, WebglLine, ColorRGBA } from 'webgl-plot'

let webglp: WebglPlot
let lines: WebglLine[] = []

const WebGlChart2: FC = observer(() => {
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

      const lastPackage = csiStore.packages.at(-1)
      if (lastPackage === undefined) {
        return
      }

      lastPackage.data.forEach((_, i) => {
        const line = new WebglLine(colors[i], csiStore.size)
        line.arrangeX()
        webglp.addLine(line)
        lines.push(line)
      })
    }
  }, [csiStore.packages, csiStore.size])

  useEffect(() => {
    let id = 0 + 1000000
    let renderPlot = () => {
      const lastPackage = csiStore.packages.at(-1)
      if (lastPackage === undefined) {
        return
      }
      for (let k = 0; k < lines.length; k++) {
        for (let i = 0; i < csiStore.size; i++) {
          lines[k].setY(i, lastPackage.data[k][i] / 150 - 1)
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
  }, [csiStore.size, csiStore.packages])

  return <canvas style={{ width: '100%', height: '100%' }} ref={canvasRef} />
})

export default WebGlChart2
