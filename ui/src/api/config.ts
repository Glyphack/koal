import axios, { AxiosError } from 'axios'
import Cookies from 'js-cookie'
import { queryClient } from '../app'
import { useAuth } from '../features/auth'

const API_URL = import.meta.env.VITE_API_URL
export const publicRequest = axios.create({ baseURL: API_URL })
export const privateRequest = axios.create({ baseURL: API_URL })
privateRequest.interceptors.request.use(
	(config) => {
		const token = Cookies.get('token')
		const bearer = `Bearer ${token}`
		if (config.headers) config.headers.Authorization = bearer
		else config.headers = { Authorization: bearer }
		return config
	},
	(error) => Promise.reject(error)
)

privateRequest.interceptors.response.use(
	(value) => value,
	(error: AxiosError) => {
		const isUnAuthorized = error.response?.status === 401
		if (isUnAuthorized) {
			useAuth.getState().signOut()
			queryClient.clear()
		}
		return error
	}
)
