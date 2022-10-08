import { FC, useEffect, useRef } from 'react'
import uPlot from 'uplot'

type OptionsUpdateState = 'keep' | 'update' | 'create'

const stringify = (obj: Record<string, unknown>) =>
  JSON.stringify(obj, (key, value) =>
    typeof value === 'function' ? value.toString() : value
  )

export const optionsUpdateState = (
  _lhs: uPlot.Options,
  _rhs: uPlot.Options
): OptionsUpdateState => {
  const { width: lhsWidth, height: lhsHeight, ...lhs } = _lhs
  const { width: rhsWidth, height: rhsHeight, ...rhs } = _rhs

  let state: OptionsUpdateState = 'keep'
  if (lhsHeight !== rhsHeight || lhsWidth !== rhsWidth) {
    state = 'update'
  }
  if (Object.keys(lhs).length !== Object.keys(rhs).length) {
    return 'create'
  }
  for (const k of Object.keys(lhs)) {
    // eslint-disable-next-line @typescript-eslint/ban-ts-comment
    // @ts-ignore
    if (stringify(lhs[k]) !== stringify(rhs[k])) {
      state = 'create'
      break
    }
  }
  return state
}

interface UplotReactProps {
  options: uPlot.Options
  data: uPlot.AlignedData
  updateDataFlag: boolean
  updateSettingsFlag: boolean
}

const LineChart: FC<UplotReactProps> = ({
  options,
  data,
  updateDataFlag,
  updateSettingsFlag,
}) => {
  const chartRef = useRef<uPlot | null>(null)
  const targetRef = useRef<HTMLDivElement>(null)

  function destroy(chart: uPlot | null) {
    if (chart) {
      chart.destroy()
      chartRef.current = null
    }
  }

  // eslint-disable-next-line react-hooks/exhaustive-deps
  function create() {
    const newChart = new uPlot(
      options,
      data,
      targetRef.current as HTMLDivElement
    )
    chartRef.current = newChart
  }

  useEffect(() => {
    create()
    return () => {
      destroy(chartRef.current)
    }
  }, [create])

  const prevProps = useRef({ options, data }).current
  useEffect(() => {
    const chart = chartRef.current
    if (prevProps.options !== options) {
      const optionsState = optionsUpdateState(prevProps.options, options)
      if (!chart || optionsState === 'create') {
        destroy(chart)
      } else if (optionsState === 'update') {
        chart.setSize({ width: options.width, height: options.height })
      }
    }
    if (!chart) {
      create()
    } else {
      chart.setData(data)
    }

    return () => {
      prevProps.options = options
      prevProps.data = data
    }
  }, [options, data, updateDataFlag, prevProps, create])

  useEffect(() => {
    destroy(chartRef?.current)
    create()
  }, [create, updateSettingsFlag])

  return <div ref={targetRef} />
}

export default LineChart
