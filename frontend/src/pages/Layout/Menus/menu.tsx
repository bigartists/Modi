import { AppstoreOutlined, MailOutlined } from '@ant-design/icons'
import { MenuProps } from 'antd'

export const items: MenuProps['items'] = [
  {
    label: 'Pipeline',
    key: 'pipeline',
    icon: <MailOutlined />,
  },
  {
    label: 'Navigation Other',
    key: 'other',
    icon: <AppstoreOutlined />,
  },
]
