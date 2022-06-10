import Cookies from 'js-cookie'
import create from 'zustand'

interface AuthState {
	isAuthenticated: boolean
	triedToVisitPage: string
	authenticate: (token: string) => void
	signOut: () => void
	setTriedToVisitPage: (page: string) => void
}

export const useAuth = create<AuthState>((set) => ({
	isAuthenticated: !!Cookies.get('token'),
	triedToVisitPage: '',
	authenticate: (token) => {
		Cookies.set('token', token)
		set({ isAuthenticated: true })
	},
	signOut: () => {
		Cookies.remove('token')
		set({ isAuthenticated: false })
	},
	setTriedToVisitPage: (page) => {
		set({ triedToVisitPage: page })
	},
}))
