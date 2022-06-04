import { Navigate, Route, Routes } from 'react-router-dom'
import { SignInPage } from '../pages/sign-in'
import { SignUpPage } from '../pages/sign-up'

export function Router() {
	return (
		<Routes>
			<Route path="/sign-up" element={<SignUpPage />} />
			<Route path="/sign-in" element={<SignInPage />} />
			<Route index element={<Navigate to="/sign-in" />} />
		</Routes>
	)
}
