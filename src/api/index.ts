import axios from 'axios'
import { SignUpRequest } from './types'
export * from './types'

const API_URL = import.meta.env.VITE_API_URL
const request = axios.create({ baseURL: API_URL })

const signUp = async (payload: SignUpRequest) => request.post('/auth/register', payload)

export const api = {
	signUp,
}
