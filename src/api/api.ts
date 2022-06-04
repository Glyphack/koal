import axios from 'axios'
import Cookies from 'js-cookie'
import { SignInRequest, SignUpRequest, SignUpResponse } from './types'

const API_URL = import.meta.env.VITE_API_URL
const publicRequest = axios.create({ baseURL: API_URL })
export const privateRequest = axios.create({ baseURL: API_URL })
privateRequest.interceptors.request.use(
	(config) => {
		const token = Cookies.get('token')
		const bearer = `Bearer ${token}`
		if (config.headers) config.headers.Authorization = bearer
		else config.headers = { Authorization: bearer }
		return config
	},
	(error) => {
		return Promise.reject(error)
	}
)

const signUp = async (payload: SignUpRequest) =>
	publicRequest.post<SignUpResponse>('/auth/register', payload)

const signIn = async (payload: SignInRequest) => publicRequest.post('/auth/login', payload)

export const api = {
	signUp,
	signIn,
}
