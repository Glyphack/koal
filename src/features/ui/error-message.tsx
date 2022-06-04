import { AxiosError } from 'axios'
import { ResponseError } from '../../api'

interface ErrorMessageProps {
	error: unknown
}

export function ErrorMessage({ error }: ErrorMessageProps) {
	if (!error) return null
	let message =
		error instanceof AxiosError ? (error.response?.data as ResponseError)?.message : null
	if (!message && error instanceof Error) message = error.message
	if (!message) return null

	return <p data-testid="error-message">{message}</p>
}
