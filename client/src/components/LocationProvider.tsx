import { useApplication } from 'hooks'
import { FC, useEffect } from 'react'
import { useLocation } from 'react-router-dom'

const LocationProvider: FC = () => {
  const location = useLocation()
  const { windowController: layoutController } = useApplication()

  useEffect(() => {
    layoutController.setLocation(location)
  }, [layoutController, location])

  return null
}

export default LocationProvider
