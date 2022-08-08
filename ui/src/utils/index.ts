import { AxiosError } from 'axios'

export const checkNotFound = (error: unknown) => {
	return error instanceof AxiosError && error.response?.status === 404
}
