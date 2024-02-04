import { FC } from 'react'
import React from 'react'
import ReactDOM from 'react-dom/client'
import { BrowserRouter } from 'react-router-dom'
import { WithBrowser, WithChakra, WithRouter, WithSidebar } from './providers'
import { WithNotifications } from './providers/with-notifications'

const App: FC = () => {
  return (
    <WithBrowser>
      <WithChakra>
        <WithNotifications />
        <BrowserRouter>
          <WithSidebar>
            <WithRouter />
          </WithSidebar>
        </BrowserRouter>
      </WithChakra>
    </WithBrowser>
  )
}

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
)
