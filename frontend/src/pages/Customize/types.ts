export interface ITask {
  task_name: string
  id: number
}
export interface ITrainingJobs {
  id: number
  log: string
  stop: string
  status: string
  project: string
  username: string
  change_time: string
  create_time: string
  pipeline_url: string
  elapsed_time: string
  final_status: string
  display_name: string
}

export interface IJobsParams {
  id: number
  name: string
  status: string
  namespace: string
  spec_html: string
  labels_html: string
  create_time: string
  info_json_html: string
  algorithm_json: string
  annotations_html: string
  status_more_html: string
}

export interface IResultAppss {
  data: ITrainingJobs[]
  count: number
  label_columns: object
}
