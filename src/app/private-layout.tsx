import Cookies from 'js-cookie'
import jwtDecode from 'jwt-decode'
import { ReactNode } from 'react'
import { useMatch } from 'react-router-dom'
import { SignOutButton } from '../features/auth'
import { Link } from '../features/ui'

interface PrivateLayoutProps {
	children: ReactNode
}

export function PrivateLayout({ children }: PrivateLayoutProps) {
	return (
		<div className="space-y-10">
			<header className="flex items-center justify-between gap-4">
				<ProjectsLink />
				<div className="flex flex-col items-end gap-1 text-right">
					<UserEmail />
					<SignOutButton />
				</div>
			</header>
			<div>{children}</div>
		</div>
	)
}

function ProjectsLink() {
	const isInProjectPage = useMatch('/project/:id')
	const isInInboxPage = useMatch('/inbox')
	const showProjectsLink = isInProjectPage || isInInboxPage

	return (
		<div>
			{showProjectsLink && (
				<Link to="/project" withLeftIcon>
					Projects
				</Link>
			)}
		</div>
	)
}

function UserEmail() {
	const token = Cookies.get('token')
	if (!token) return null
	const user = jwtDecode<{ sub: string }>(token)
	const email = user.sub
	return <h6 className="text-xs text-gray-500">{email}</h6>
}
