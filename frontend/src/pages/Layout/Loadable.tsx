import { defaultLazyLoad } from 'routes/loadable'

export default defaultLazyLoad(
  () => import(/* webpackChunkName: "mainPage" */ './index'),
  module => module.LayoutPage,
)
