import {
  CheckCircleOutlined,
  ClockCircleOutlined,
  ExclamationCircleOutlined,
  SyncOutlined,
} from '@ant-design/icons'
import {
  Button,
  PaginationProps,
  Popconfirm,
  Space,
  Table,
  Tag,
  Modal,
} from 'antd'
import { ColumnsType } from 'antd/lib/table'
import React, { memo, useCallback, useEffect, useRef, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import Wrapper from 'components/ContainerLayout/ListWrapper'
import { ITask } from './types'
import { v4 as uuid } from 'uuid'
import { datefmtFn, durationTs } from 'utils/format'

const CustomizeTask: React.FC = memo(props => {
  const navigate = useNavigate()

  const log = useCallback(
    (logId: number) => () => {
      navigate(`log/${logId}`)
    },
    [navigate],
  )

  const [open, setOpen] = useState<boolean>(false)

  const columns: ColumnsType<ITask> = [
    {
      title: '任务名称',
      dataIndex: 'task_name',
      key: 'task_name',
      render: (text, record) => <a onClick={log(record.id)}>{text}</a>,
    },
    {
      title: '操作',
      width: '15%',
      key: 'action',
      render: (_: any, record: ITask) => {
        const displayStatus = record.task_name
        return (
          <Space size="middle">
            <a key={uuid()} onClick={log(record.id)}>
              编辑
            </a>

            <Popconfirm title="确定重跑该任务吗">
              <a key={uuid()}>查看</a>
            </Popconfirm>
          </Space>
        )
      },
    },
  ]

  const create = useCallback(() => {
    // navigate('create')
    setOpen(!open)
  }, [open])

  const operations = (
    <Space>
      {/* <Button>任务名称</Button>
      <Button onClick={detail}>任务状态</Button> */}
      <Button type="primary" onClick={create}>
        创建任务
      </Button>
      <Button icon={<SyncOutlined />}></Button>
    </Space>
  )

  const onChange: PaginationProps['onChange'] = (
    pageNumber: number,
    pageSize: number,
  ) => {
    console.log('Page: ', pageNumber, 'pagesize: ', pageSize)
  }

  const data = React.useMemo(() => {
    return new Array(50)
      .fill({
        id: uuid(),
        task_name: '任务名称',
      })
      .map(task => {
        return {
          ...task,
          task_name: `${task.task_name}-${task.id}`,
        }
      })
  }, [])

  return (
    <Wrapper
      pagination={{
        showQuickJumper: true,
        defaultCurrent: 1,
        total: 100,
        onChange: onChange,
      }}
    >
      <div className="tool">{operations}</div>
      <article className="w-full h-full">
        <Table
          rowKey={record => uuid()}
          dataSource={data}
          columns={columns}
          pagination={false}
        />
      </article>

      <Modal title="Basic Modal" open={open} onOk={create} onCancel={create}>
        <p>Some contents...</p>
        <p>Some contents...</p>
        <p>Some contents...</p>
      </Modal>
    </Wrapper>
  )
})

export default CustomizeTask
