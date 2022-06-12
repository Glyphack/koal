import { ReactNode } from 'react'
import { HiArrowSmLeft } from 'react-icons/hi'
import { useQueryClient } from 'react-query'
import { useMatch } from 'react-router-dom'
import { useAuth } from '../features/auth'
import { Button, Link } from '../features/ui'

interface PrivateLayoutProps {
	children: ReactNode
}

export function PrivateLayout({ children }: PrivateLayoutProps) {
	const isInProjectPage = useMatch('/project/:id')
	const signOut = useAuth((state) => state.signOut)
	const queryClient = useQueryClient()
	const handleSignOut = () => {
		signOut()
		queryClient.clear()
	}

	return (
		<div className="space-y-10">
			<header className="flex items-center justify-between gap-4">
				<div>
					{isInProjectPage && (
						<Link to="/project" withLeftIcon>
							Projects
						</Link>
					)}
				</div>
				<Button variant="text" onClick={handleSignOut}>
					<HiArrowSmLeft /> Sign Out
				</Button>
			</header>
			<div>{children}</div>
		</div>
	)
}
