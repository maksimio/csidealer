import { Container } from '@chakra-ui/react'
import LineChart from 'components/shared/LineChart'
import { useApplication } from 'hooks'
import { observer } from 'mobx-react-lite'
import { FC } from 'react'

const Dashboard: FC = observer(() => {
  const { csiStore } = useApplication()
  const csi = csiStore.packages.at(-1)?.data
  if (csi === undefined) {
    return null
  }

  return (
    <Container>
      <LineChart
        options={{
          width: 840,
          height: 520,
          series: [
            {},
            { label: 'h11', stroke: 'black' },
            { label: 'h12', stroke: 'red' },
            { label: 'h21', stroke: 'green' },
            { label: 'h22', stroke: 'blue' },
          ],
          scales: { x: { time: false }, y: { min: 40, max: 300 } },
        }}
        data={[csiStore.x, ...csi]}
        updateDataFlag={csiStore.updFlag}
        updateSettingsFlag={false}
      />
    </Container>
  )
})

export default Dashboard
