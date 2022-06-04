import { ReactNode } from 'react'
import { Navigate } from 'react-router-dom'
import { useAuth } from './use-auth'

interface RequireAuthProps {
	children: ReactNode
}

export function RequireAuth({ children }: RequireAuthProps) {
	const isAuthenticated = useAuth((state) => state.isAuthenticated)
	if (!isAuthenticated) return <Navigate to="/sign-in" replace />
	return <>{children}</>
}
