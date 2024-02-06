import { Navigate, useLocation } from 'react-router-dom'
import { getLocal } from 'utils/storageManage'
import { USERNAME } from './slice/thunk'

function RequireAuth({ children }: { children: JSX.Element }) {
  const username = getLocal(USERNAME)
  let location = useLocation()
  console.log('🚀 ~ file: index.tsx:128 ~ SignInloader ~ username:', username)

  if (!username) {
    return <Navigate to="/login" state={{ from: location }} replace />
  }

  return children
}

export default RequireAuth
