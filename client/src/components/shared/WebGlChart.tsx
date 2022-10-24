import { useApplication } from 'hooks'
import { observer } from 'mobx-react-lite'
import { FC, useEffect, useRef } from 'react'
import { WebglPlot, WebglLine, ColorRGBA } from 'webgl-plot'

let webglp: WebglPlot
let line: WebglLine

interface WebGlChartProps {
  freq: number
  amp: number
  noise?: number
}

const WebGlChart: FC<WebGlChartProps> = observer(({ freq, amp, noise }) => {
  const canvasRef = useRef<HTMLCanvasElement>(null)
  const { csiStore } = useApplication()

  useEffect(() => {
    if (canvasRef.current) {
      const devicePixelRatio = window.devicePixelRatio || 1
      canvasRef.current.width = canvasRef.current.clientWidth * devicePixelRatio
      canvasRef.current.height = canvasRef.current.clientHeight * devicePixelRatio

      webglp = new WebglPlot(canvasRef.current)
      const numX = 1000

      line = new WebglLine(new ColorRGBA(1, 0, 0, 1), numX)
      console.log(line)
      webglp.addLine(line)
      line.arrangeX()
    }
  }, [])

  useEffect(() => {
    let id = 0
    let renderPlot = () => {
      for (let i = 0; i < line.numPoints; i++) {
        line.setY(i, csiStore.timeseries[0][i] / 200)
      }
      id = requestAnimationFrame(renderPlot)
      webglp.update()
    }
    id = requestAnimationFrame(renderPlot)
    return () => {
      renderPlot = () => {}
      cancelAnimationFrame(id)
    }
  }, [freq, amp, noise, csiStore.timeseries])

  const canvasStyle = {
    width: '100%',
    height: '70vh',
  }

  return <canvas style={canvasStyle} ref={canvasRef} />
})

export default WebGlChart
