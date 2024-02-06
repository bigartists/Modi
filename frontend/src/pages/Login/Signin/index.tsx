import { Button, Form, Input, Typography, type FormProps, Space } from 'antd'
import { LockOutlined, UserOutlined } from '@ant-design/icons'
import { redirect, useNavigate } from 'react-router-dom'
import { useCallback, useState, memo } from 'react'
import styles from './index.module.less'
import bg from 'assets/images/bg_light.png'
import { useDispatch } from 'react-redux'
import { USERNAME, login } from 'pages/Layout/slice/thunk'
import { getLocal } from 'utils/storageManage'
export interface ISigninProps {}

export const SignIn: React.FC<ISigninProps> = memo(() => {
  const navigate = useNavigate()
  const [form] = Form.useForm()
  const [loading, setLoading] = useState(false)
  const dispatch = useDispatch()

  const onFinished: FormProps['onFinish'] = async (formData: any) => {
    setLoading(true)
    try {
      await dispatch(
        login({
          username: formData.username,
          password: formData.password,
        }),
      )
      console.log('login success =>')
      //   return redirect('/pipeline')
      navigate('/pipeline')
    } catch (error: any) {
      console.log('login error =>', error)
      if (typeof error === 'string') {
        form.setFields([
          {
            name: 'password',
            errors: [error],
          },
        ])
      }
    }
    setLoading(false)
  }

  const toSignup = useCallback(() => {
    navigate('/signup')
  }, [navigate])

  return (
    <div
      className={`${styles.container} flex-1 flex`}
      style={{
        backgroundImage: `url(${bg})`,
      }}
    >
      <div className={styles.minwidth}>
        <div className={styles.layout}>{/* <Background></Background> */}</div>
        <div className={styles.content}>
          <div className="text-center">
            <div></div>
            <div className={styles.title}>产品名称</div>
            <div className={styles.model}>产品一句话描述</div>
          </div>
          <Form name="login" className="form" onFinish={onFinished} form={form}>
            <Form.Item
              name="username"
              rules={[
                {
                  required: true,
                  message: '请输入用户名/手机号',
                },
              ]}
            >
              <Input
                prefix={<UserOutlined style={{ color: '#BFBFBF' }} />}
                placeholder="请输入用户名/手机号"
              ></Input>
            </Form.Item>
            <Form.Item
              name="password"
              style={{ marginBottom: '0' }}
              rules={[
                {
                  required: true,
                  message: '请输入6-18位密码',
                },
              ]}
            >
              <Input.Password
                placeholder="请输入6-18位密码"
                prefix={<LockOutlined style={{ color: '#BFBFBF' }} />}
              ></Input.Password>
            </Form.Item>
            <Space className="flex justify-between mb-10">
              <span style={{ fontSize: '12px' }}>
                没有账号？
                <Typography.Link
                  onClick={toSignup}
                  style={{ color: '#0047BA', fontSize: '12px' }}
                >
                  点击注册
                </Typography.Link>
              </span>
              <Typography.Link style={{ color: '#0047BA', fontSize: '12px' }}>
                忘记密码
              </Typography.Link>
            </Space>

            <Form.Item>
              <Button
                block
                htmlType={'submit'}
                type={'primary'}
                size={'large'}
                loading={loading}
              >
                登录
              </Button>
            </Form.Item>
          </Form>
        </div>
      </div>
    </div>
  )
})

export function SignInloader() {
  const username = getLocal(USERNAME)
  console.log('🚀 ~ file: index.tsx:128 ~ SignInloader ~ username:', username)

  if (username) {
    return redirect('/pipeline')
  }
  return null
}

export default SignIn
