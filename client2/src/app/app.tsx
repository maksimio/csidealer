import { FC } from 'react'
import React from 'react'
import ReactDOM from 'react-dom/client'
import { BrowserRouter } from 'react-router-dom'
import { WithBrowser, WithChakra, WithRouter, WithSidebar } from './providers'

const App: FC = () => {
  return (
    <WithBrowser>
      <WithChakra>
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
