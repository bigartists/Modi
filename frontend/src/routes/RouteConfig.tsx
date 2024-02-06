import { createBrowserRouter, createHashRouter } from 'react-router-dom'
import LayoutLazyPage from 'pages/Layout/Loadable'
import CustomizeLazyPage from 'pages/Customize/Loadable'
import NotFoundPage from 'pages/NotFoundPage'
import SignInLazyPage from 'pages/Login/Signin/Loadable'
import { SignInloader } from 'pages/Login/Signin'
export const routerList = createHashRouter([
  {
    id: 'root',
    path: '/',
    async lazy() {
      return {
        Component: LayoutLazyPage,
      }
    },
    children: [
      {
        index: true,
        async lazy() {
          return {
            Component: CustomizeLazyPage,
          }
        },
      },
      {
        path: 'pipeline',
        async lazy() {
          return {
            Component: CustomizeLazyPage,
          }
        },
      },
    ],
  },
  {
    path: 'login',
    async lazy() {
      return {
        Component: SignInLazyPage,
      }
    },
    loader: SignInloader,
  },
  {
    path: '*',
    Component: NotFoundPage,
  },
])
