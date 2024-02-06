export interface UserState {
  username: string
  isLogin?: boolean
  loginLoading: boolean
  collapsed: boolean
  theme: 'light' | 'dark'
}

export interface User {
  username: string
  storageMap?: Object
}

export interface LoginParams {
  username: string
  password: string
}
