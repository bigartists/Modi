import { RouterProvider } from 'react-router-dom'
import {
  ConfigProvider,
  GlobalToken,
  Layout,
  Switch,
  theme,
  App as AntdApp,
} from 'antd'
import zhCN from 'antd/es/locale/zh_CN'
import { routerList } from 'routes/RouteConfig'
import { useSelector, useDispatch } from 'react-redux'
import { selectTheme } from 'pages/Layout/slice/selector'
import { MessageInstance } from 'antd/es/message/interface'
import { NotificationInstance } from 'antd/es/notification/interface'
import { ModalStaticFunctions } from 'antd/es/modal/confirm'

let message: MessageInstance
let notification: NotificationInstance
let modal: Omit<ModalStaticFunctions, 'warn'>

export default function App() {
  const rootTheme = useSelector(selectTheme)
  const staticFunction = AntdApp.useApp()
  message = staticFunction.message
  modal = staticFunction.modal
  notification = staticFunction.notification

  return (
    <ConfigProvider
      locale={zhCN}
      theme={{
        token: {
          colorPrimary: '#1946b9',
          fontSize: 14,
          borderRadius: 4,
        },
        algorithm:
          rootTheme === 'light' ? theme.defaultAlgorithm : theme.darkAlgorithm,
      }}
    >
      <RouterProvider
        router={routerList}
        fallbackElement={<p>页面加载中...</p>}
      />
    </ConfigProvider>
  )
}

export { modal, message, notification }
