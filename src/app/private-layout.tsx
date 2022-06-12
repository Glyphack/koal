import { ReactNode } from 'react'
import { HiArrowSmLeft } from 'react-icons/hi'
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
		<div className="space-y-10">
			<header className="flex justify-end">
				<Button variant="text" onClick={handleSignOut}>
					<HiArrowSmLeft /> Sign Out
				</Button>
			</header>
			<div>{children}</div>
		</div>
	)
}
