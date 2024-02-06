import { message } from 'pages/App'
import axios, { AxiosRequestConfig, AxiosInstance } from 'axios'
import { APIResponse } from 'types'
const baseURL = '/api/v1'
export const request = axios.create({
  baseURL,
  validateStatus(status) {
    return status < 400 // 约束http status<400 的进入resolved
  },
})

request.interceptors.request.use(
  function (config) {
    const { userInfo } = {
      userInfo: {
        authCode: 'test',
      },
    }
    if (userInfo.authCode) {
      config.headers = Object.assign({}, config.headers, {
        authcode: userInfo.authCode,
      })
    }
    return config
  },
  function (error) {
    console.log('request error =>', error)
    return Promise.reject(error)
  },
)

// 添加响应拦截器
request.interceptors.response.use(
  function (response) {
    console.log('response =>', response)
    if (response?.data?.status !== 1) {
      if (response.data?.message === '登录过期') {
        // 调用logout方法
        //  logout()
        Promise.reject(response.data)
      }
      const errMessage = response?.data?.message || '请求失败'
      message.error(errMessage)
      return Promise.reject(errMessage)
    }

    return response?.data
  },
  function (error) {
    console.log('response error =>', error)

    if (error.response?.status === 401) {
      // logout()
      return
    }

    message.error(error.message)

    return Promise.reject(error)
  },
)
