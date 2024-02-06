import { Tabs, theme } from 'antd'
import React, { memo } from 'react'
import Wrapper from 'components/ContainerLayout/ListWrapper'
import Topic from 'components/Topic'
import Task from './Task'
import { GlobalToken } from 'antd'
import styled from 'styled-components'

export const PipelineTask: React.FC = memo(props => {
  const { token } = theme.useToken()
  return (
    <Wrapper hasNoFooter>
      <Topic topic="this is a topic" description="this is a topic desc" />

      <CustomTabsStyle token={token}>
        <Tabs
          defaultActiveKey="1"
          onChange={(key: string) => {
            console.log(key)
          }}
          items={[
            {
              label: `任务列表`,
              key: 'Task',
              children: <Task />,
            },
          ]}
        />
      </CustomTabsStyle>
    </Wrapper>
  )
})

interface IToken {
  token: GlobalToken
}

const CustomTabsStyle = styled.div<IToken>`
  .ant-tabs-top > .ant-tabs-nav {
    background-color: ${props => props.token?.colorBgContainer};
    margin-bottom: 0;
  }
  .ant-tabs-nav {
    background-color: ${props => props.token?.colorBgContainer};
    margin-bottom: 0;
  }
  .ant-tabs-tab {
    padding-left: 28px;
    padding-right: 28px;
  }
  .ant-tabs-tab + .ant-tabs-tab {
    margin-left: 0;
  }
`

export default PipelineTask
