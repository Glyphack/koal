import { ReactNode } from 'react'
import { Navigate } from 'react-router-dom'
import { useAuth } from './use-auth'

interface RequireUnAuth {
	children: ReactNode
}

export function RequireUnAuth({ children }: RequireUnAuth) {
	const isAuthenticated = useAuth((state) => state.isAuthenticated)
	if (isAuthenticated) return <Navigate to="/project" replace />
	return <>{children}</>
}
