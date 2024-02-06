import { message } from 'pages/App'
import { AxiosError } from 'axios'
// export { default as uuidv4 } from 'uuid/dist/umd/uuidv4.min'

export function errorHandle(error: any) {
  if (error?.response) {
    // AxiosError
    const { response } = error as AxiosError<{
      message: string
    }>
    switch (response?.status) {
      case 401:
        message.error({ key: '401', content: '401' })

        break
      case 400:
        message.error(response?.data?.message || response?.statusText)
        break
      default:
        message.error(response?.statusText || error.message)
        break
    }
  } else if (error?.message) {
    // Error
    message.error({
      content: error.message,
      duration: 2,
    })
  } else {
    message.error({
      content: error,
      duration: 2,
    })
  }
  return error
}

export const isNumberFunc = (arg: any): boolean => {
  return typeof arg === 'number'
}

export const getDataType = (data: any): string => {
  return (
    Object.prototype.toString.call(data).match(/\s(\w+)\]/) as string[]
  )[1]
}

export const overrideRequire = (path: string) => {
  const env = process.env.IS_VITE
  debugger
  const prefix = '/src/'
  if (!env) {
    // return require(path)
  } else {
    return `${prefix}${path}`
  }
}
