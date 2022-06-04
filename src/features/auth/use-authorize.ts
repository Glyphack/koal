import Cookies from 'js-cookie'
import { useNavigate } from 'react-router-dom'

export const useAuthorize = () => {
	const navigate = useNavigate()
	const authorize = (token: string) => {
		Cookies.set('token', token)
		navigate('/project')
	}
	return authorize
}
