import { ReactNode } from 'react'
import { useAuth } from '../features/auth'
import { Button } from '../features/ui'

interface PrivateLayoutProps {
	children: ReactNode
}

export function PrivateLayout({ children }: PrivateLayoutProps) {
	const signOut = useAuth((state) => state.signOut)

	return (
		<div>
			<header>
				<Button onClick={signOut}>Sign Out</Button>
			</header>
			<div>{children}</div>
		</div>
	)
}
