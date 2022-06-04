import Cookies from 'js-cookie'
import create from 'zustand'

interface AuthState {
	isAuthenticated: boolean
	authenticate: (token: string) => void
	logout: () => void
}

export const useAuth = create<AuthState>((set) => ({
	isAuthenticated: !!Cookies.get('token'),
	authenticate: (token) => {
		Cookies.set('token', token)
		set({ isAuthenticated: true })
	},
	logout: () => {
		Cookies.remove('token')
		set({ isAuthenticated: false })
	},
}))
