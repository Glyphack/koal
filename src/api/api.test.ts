import Cookies from 'js-cookie'
import { privateRequest } from './api'

describe('API', () => {
	test('should add authorization header for private requests', async () => {
		Cookies.set('token', 'access-token')
		const response = await privateRequest.get('/mock')
		const authorizationHeader = response.config.headers?.Authorization as string
		expect(authorizationHeader).toBe('Bearer access-token')
	})
})
