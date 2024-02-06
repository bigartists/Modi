import React from 'react'
import ReactDOM from 'react-dom/client'
import { ConfigProvider, theme } from 'antd'
import App from 'pages/App'
import appCombineReducers from 'redux/reducer'

import { Inspector } from 'react-dev-inspector'
import { HelmetProvider } from 'react-helmet-async'
import { Provider } from 'react-redux'

import configureStore from 'redux/store'
import dayjs from 'dayjs'
import 'dayjs/locale/zh-cn'
import 'react-app-polyfill/ie11'
import 'react-app-polyfill/stable'
import 'assets/styles/css/tw.css'
import 'assets/styles/css/index.css'
import 'antd/dist/reset.css'
dayjs.locale('zh-cn')
const root = ReactDOM.createRoot(document.getElementById('root') as HTMLElement)
const IS_DEVELOPMENT = process.env.NODE_ENV === 'development'
export const store = configureStore(appCombineReducers)
const InspectorWrapper = IS_DEVELOPMENT ? Inspector : React.Fragment
//选中要审查的元素， 按住ctrl+command+shfit+c
root.render(
  <InspectorWrapper>
    <Provider store={store}>
      <HelmetProvider>
        <React.StrictMode>
          <App />
        </React.StrictMode>
      </HelmetProvider>
    </Provider>
  </InspectorWrapper>,
)
