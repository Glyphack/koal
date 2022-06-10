import { Navigate, Route, Routes } from 'react-router-dom'
import { RequireAuth } from '../features/auth'
import { NotFoundPage } from '../pages/404'
import { ProjectPage } from '../pages/project'
import { ProjectsPage } from '../pages/projects'
import { SignInPage } from '../pages/sign-in'
import { SignUpPage } from '../pages/sign-up'
import { PrivateLayout } from './private-layout'

const publicRoutes = [
	{ path: '/sign-up', element: <SignUpPage /> },
	{ path: '/sign-in', element: <SignInPage /> },
]
const privateRoutes = [
	{ path: '/project', element: <ProjectsPage /> },
	{ path: '/project/:id', element: <ProjectPage /> },
]

const publicRouter = publicRoutes.map((route) => (
	<Route key={route.path} path={route.path} element={route.element} />
))
const privateRouter = privateRoutes.map((route) => (
	<Route
		key={route.path}
		path={route.path}
		element={
			<RequireAuth>
				<PrivateLayout>{route.element}</PrivateLayout>
			</RequireAuth>
		}
	/>
))

export function Router() {
	return (
		<Routes>
			{publicRouter}
			{privateRouter}
			<Route index element={<Navigate to="/project" />} />
			<Route path="*" element={<NotFoundPage />} />
		</Routes>
	)
}
