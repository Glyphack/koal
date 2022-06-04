import Cookies from 'js-cookie'
import create from 'zustand'

interface AuthState {
	isAuthenticated: boolean
	authenticate: (token: string) => void
	signOut: () => void
}

export const useAuth = create<AuthState>((set) => ({
	isAuthenticated: !!Cookies.get('token'),
	authenticate: (token) => {
		Cookies.set('token', token)
		set({ isAuthenticated: true })
	},
	signOut: () => {
		Cookies.remove('token')
		set({ isAuthenticated: false })
	},
}))
