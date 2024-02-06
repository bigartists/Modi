import { Dropdown, GlobalToken, Space, theme } from 'antd'
import styled from 'styled-components'
import { useSelector } from 'react-redux'
import { DownOutlined } from '@ant-design/icons'
import { Link, useNavigate } from 'react-router-dom'
import { SPACE_TIMES } from 'assets/styles/styledcom/StyleConstants'
import { selectTheme, selectUsername } from 'pages/Layout/slice/selector'
import { getLocal } from 'utils/storageManage'
import { USERNAME } from 'pages/Layout/slice/thunk'
const docLink = `https://github.com/mozhehanghui/react-antd-scaffold`

export default () => {
  const navigate = useNavigate()
  // const username = useSelector(selectUsername)
  const username = getLocal(USERNAME) as string
  console.log('🚀 ~ file: index.tsx:16 ~ username:', username)
  const { token } = theme.useToken()
  const rootTheme = useSelector(selectTheme)

  const logout = () => {
    navigate('/login')
  }

  const items = [
    {
      label: <Link to="/userCenter">用户中心</Link>,
      key: 'doc',
    },
    {
      label: <a onClick={logout}>退出登录</a>,
      key: 'signin',
    },
  ]

  return (
    <LayoutHeader token={token} rootTheme={rootTheme}>
      <div className="nav">
        <div className="logo">
          <Link to="/">
            <div className="normal">
              <span className="title">logo</span>
            </div>
          </Link>
        </div>
      </div>
      <div className="userinfo">
        <Space size="large">
          <a href={docLink} target="_blank" rel="noreferrer" title="用户手册">
            <Space>
              <span>操作手册</span>
              <span
                style={{ fontSize: '12px' }}
                className="iconfont tc-KingBI-share"
              ></span>
            </Space>
          </a>
          <span></span>
          <Dropdown menu={{ items }} trigger={['click']}>
            <Space onClick={e => e.preventDefault()}>
              <span className="username">
                <span className="iconfont tc-user_circle"></span>
                <span>{username || '未登录'}</span>
                <DownOutlined />
              </span>
            </Space>
          </Dropdown>
        </Space>
      </div>
    </LayoutHeader>
  )
}

interface ILayoutHeader {
  token: GlobalToken
  rootTheme: 'light' | 'dark'
}

const LayoutHeader = styled.div<ILayoutHeader>`
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  width: 100%;
  flex: 0 0 auto;
  z-index: 3;
  display: flex;
  height: ${SPACE_TIMES(13.5)};
  line-height: ${SPACE_TIMES(13.5)};
  backdrop-filter: blur(20px);
  box-shadow: 1px 12px 80px rgba(0, 82, 217, 0.3);
  transition: left 0.3s cubic-bezier(0.25, 0.1, 0.25, 1) 0s;
  background: linear-gradient(
    to right,
    rgba(255, 255, 255, 0.65),
    rgb(242, 243, 255)
  );

  background: ${props =>
    props.rootTheme === 'light'
      ? 'linear-gradient(to right, rgba(255, 255, 255, 0.65), rgb(242, 243, 255))'
      : 'linear-gradient(to right, rgba(0, 0, 0, 0.65), rgb(0, 0, 0))'};

  background-color: ${props =>
    props.rootTheme === 'dark' ? props.token?.colorBgContainer : ''};

  .nav {
    flex: 1;
    margin-left: ${SPACE_TIMES(2)};
    .logo {
      flex: 0 0 auto;
      display: flex;
      height: ${SPACE_TIMES(12.5)};
      line-height: ${SPACE_TIMES(11)};
      padding-left: ${SPACE_TIMES(4)};
      margin-bottom: ${SPACE_TIMES(2)};

      font-size: ${SPACE_TIMES(5)};
      cursor: pointer;
      // justify-content: space-between;
      align-items: center;
      user-select: none;
      a {
        text-decoration: none;
      }
      .title {
        color: ${props => props.token?.colorText};
        font-size: ${SPACE_TIMES(4.5)};
        margin-left: ${SPACE_TIMES(3)};
        text-decoration: none;

        .dot {
          margin: 0 2px;
          font-size: ${SPACE_TIMES(6)};
        }
      }
    }
  }
  .userinfo {
    padding: 0 ${SPACE_TIMES(4)} 0 0;
    flex: 0 0 auto;
    display: flex;
    color: ${props => props.token?.colorTextSecondary};
    justify-content: space-around;
    align-items: center;
    a {
      color: ${props => props.token?.colorTextSecondary};
    }
    .username {
      color: ${props => props.token?.colorTextSecondary};
      font-weight: 400;
      user-select: none;
      cursor: pointer;
      span {
        margin-left: ${SPACE_TIMES(2)};
      }
    }
  }
`
