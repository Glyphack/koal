import { AxiosError } from 'axios'
import { checkNotFound } from '.'

const mockBadRequestError = new AxiosError(undefined, undefined, undefined, undefined, {
	data: { message: 'response test error message' },
	status: 400,
	statusText: 'Bad Request',
	config: {},
	headers: {},
})

const mockNotFoundError = new AxiosError(undefined, undefined, undefined, undefined, {
	data: { message: 'response test error message' },
	status: 404,
	statusText: 'Not Found',
	config: {},
	headers: {},
})

describe('Utils', () => {
	test('checkNotFound should work', () => {
		expect(checkNotFound(new Error())).toBeFalsy()
		expect(checkNotFound(mockBadRequestError)).toBeFalsy()
		expect(checkNotFound(mockNotFoundError)).toBeTruthy()
	})
})
