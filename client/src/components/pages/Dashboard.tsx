import { Button } from '@chakra-ui/react'
import { useApplication } from 'hooks'
import { FC } from 'react'


const Dashboard: FC = () => {
  const { logFileController } = useApplication()
  return (
    <> Главная
      <Button onClick={() => logFileController.getList()}>Тест</Button>
    </>
  )
}

export default Dashboard