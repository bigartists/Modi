/**
 * Asynchronously loads the component for NotFoundPage
 */

import { LoadingOutlined } from '@ant-design/icons'
import { lazyLoad } from 'routes/loadable'

export default lazyLoad(
  () => import(/* webpackChunkName: "SignIn" */ './index'),
  module => module.SignIn,
  {
    fallback: <LoadingOutlined />,
  },
)
