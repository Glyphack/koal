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
	const email = getEmailFromToken()
	if (email.error) return null
	return <h6 className="text-xs text-gray-500">{email.data}</h6>
}

function getEmailFromToken() {
	const token = Cookies.get('token')
	if (!token) return { error: true, data: undefined }
	try {
		const user = jwtDecode<{ sub: string }>(token)
		const email = user.sub
		return { error: false, data: email }
	} catch (error) {
		return { error: true, data: undefined }
	}
}
