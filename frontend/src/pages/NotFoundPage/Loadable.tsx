/**
 * Asynchronously loads the component for NotFoundPage
 */

import { LoadingOutlined } from '@ant-design/icons'
import React from 'react'
import { lazyLoad } from 'routes/loadable'

export const NotFoundPage = lazyLoad(
  () => import(/* webpackChunkName: "NotFoundPage" */ './index'),
  module => module.NotFoundPage,
  {
    fallback: <LoadingOutlined />,
  },
)
