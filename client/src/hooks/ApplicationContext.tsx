import { createContext, FC, ReactNode } from 'react'
import Application from 'application'
import { useContext } from 'react'

interface ApplicationProviderProps {
  children: ReactNode
}

const ApplicationContext = createContext<null | Application>(null)

export const ApplicationProvider: FC<ApplicationProviderProps> = ({ children }) => {
  const application = new Application()
  return (
    <ApplicationContext.Provider value={application}>
      {children}
    </ApplicationContext.Provider>
  )
}

export const useApplication = (): Application => {
  const application = useContext(ApplicationContext)
  if (!application) {
    throw new Error('useApplication должен быть использован в ApplicationProvider')
  }
  return application
}