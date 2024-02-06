import { ConfigProvider, MenuProps } from 'antd'
import { Menu } from 'antd'
import { items } from './menu'
import React, { useCallback, useEffect, useMemo, useState } from 'react'
import { useLocation, useNavigate } from 'react-router-dom'
import styled from 'styled-components'
import { BLUE } from 'assets/styles/styledcom/StyleConstants'

const Menus = () => {
  const location = useLocation()
  const navigate = useNavigate()

  const [selectedKeys, setSelectedKeys] = useState<string[]>([])
  const [openKeys, setOpenKeys] = useState<string[]>([])

  const handleSelect: MenuProps['onSelect'] = ({ key, selectedKeys }) => {
    navigate(key)
    setSelectedKeys([key])
  }

  function getPath<T>(array: Array<T>) {
    let prev: T
    return array.map((arr: T, index) => {
      if (index === 0) {
        prev = arr
        return arr
      } else if (index > 0) {
        const current: T = `${prev}/${arr}` as T
        prev = current
        return current
      }
      return arr
    })
  }

  const getOpenKeys = useMemo(() => {
    let arr = location.pathname.split('/')
    arr = arr.slice(1, arr.length - 1)
    return getPath<string>(arr)
  }, [location])

  useEffect(() => {
    const selectedKey = location.pathname.slice(1)
    setSelectedKeys([selectedKey])
    setOpenKeys(getOpenKeys as string[])
  }, [getOpenKeys, location])

  const onOpenChange = useCallback((val: string[]) => {
    setOpenKeys(val)
  }, [])

  return (
    <ConfigProvider
      theme={{
        token: {
          borderRadius: 4,
        },
      }}
    >
      <Wrapper>
        <Menu
          mode="inline"
          items={items}
          style={{ height: '100%', userSelect: 'none', borderRight: 0 }}
          openKeys={openKeys}
          onSelect={handleSelect}
          onOpenChange={onOpenChange}
          selectedKeys={selectedKeys}
        />
      </Wrapper>
    </ConfigProvider>
  )
}

const Wrapper = styled.div`
  .ant-menu {
    /* border-right: 1px solid #f0f0f0; */
    .ant-menu-item {
      border-radius: 3px;
      font-size: 14px;
      &:not(.ant-menu-item-selected):hover {
        color: ${BLUE};
      }
      &.ant-menu-item-active {
        color: ${BLUE};
      }

      &.ant-menu-item-selected {
        background-color: ${BLUE};
        color: #fff;
      }
    }
  }
`

export default Menus
