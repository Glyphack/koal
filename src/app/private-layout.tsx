import { ReactNode } from 'react'
import { useQueryClient } from 'react-query'
import { useAuth } from '../features/auth'
import { Button } from '../features/ui'

interface PrivateLayoutProps {
	children: ReactNode
}

export function PrivateLayout({ children }: PrivateLayoutProps) {
	const signOut = useAuth((state) => state.signOut)
	const queryClient = useQueryClient()
	const handleSignOut = () => {
		signOut()
		queryClient.clear()
	}

	return (
		<div>
			<header>
				<Button onClick={handleSignOut}>Sign Out</Button>
			</header>
			<div>{children}</div>
		</div>
	)
}
