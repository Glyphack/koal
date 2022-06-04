import { render, screen } from '../../test/test-utils'
import { SignInPage } from '../sign-in'

describe('SignInPage', () => {
	test('should have sign up link', async () => {
		render(<SignInPage />)
		const signInLink = screen.getByRole('link', { name: /sign up/i })
		expect(signInLink).toBeInTheDocument()
	})
})
