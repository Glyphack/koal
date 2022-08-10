import { ReactNode, useEffect } from 'react'
import { Navigate, useLocation } from 'react-router-dom'
import { useAuth } from './use-auth'

interface RequireAuthProps {
	children: ReactNode
}

export function RequireAuth({ children }: RequireAuthProps) {
	const location = useLocation()
	const { isAuthenticated, setTriedToVisitedPage } = useAuth((state) => ({
		isAuthenticated: state.isAuthenticated,
		setTriedToVisitedPage: state.setTriedToVisitPage,
	}))

	useEffect(() => {
		if (!isAuthenticated) setTriedToVisitedPage(location.pathname)
	}, [isAuthenticated, location.pathname, setTriedToVisitedPage])

	if (!isAuthenticated) return <Navigate to="/sign-in" replace />
	return <>{children}</>
}
