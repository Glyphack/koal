import { AxiosError } from 'axios'
import { render, screen } from '../../../test/test-utils'
import { ErrorMessage } from '../error-message'

const mockResponseError = new AxiosError(undefined, undefined, undefined, undefined, {
	data: { message: 'response test error message' },
	status: 400,
	statusText: 'Bad Request',
	config: {},
	headers: {},
})

describe('ErrorMessage', () => {
	test('renders null if no error', () => {
		render(<ErrorMessage error={null} />)
		const errorMessage = screen.queryByTestId('error-message')
		expect(errorMessage).not.toBeInTheDocument()
	})

	test('renders error message', () => {
		render(<ErrorMessage error={new Error('test error message')} />)
		const errorMessage = screen.getByText(/test error message/i)
		expect(errorMessage).toBeInTheDocument()
	})

	test('renders response error message', () => {
		render(<ErrorMessage error={mockResponseError} />)
		const errorMessage = screen.getByText(/response test error message/i)
		expect(errorMessage).toBeInTheDocument()
	})
})
