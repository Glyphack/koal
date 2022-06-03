import { Navigate, Route, Routes } from 'react-router-dom'
import { SignUpPage } from '../pages/sign-up'

export function Router() {
	return (
		<Routes>
			<Route path="/sign-up" element={<SignUpPage />} />
			<Route index element={<Navigate to="/sign-up" />} />
		</Routes>
	)
}
