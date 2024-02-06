import { defaultLazyLoad } from 'routes/loadable'

export default defaultLazyLoad(
  () => import(/* webpackChunkName: "pipelineTask" */ './index'),
  module => module.PipelineTask,
)
