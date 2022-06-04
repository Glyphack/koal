import { render, screen } from '../../test/test-utils'
import { SignInPage } from '../sign-in'

describe('SignInPage', () => {
	test('should be able to sign in', async () => {
		const { user } = render(<SignInPage />)
		const emailField = screen.getByPlaceholderText(/email/i)
		const passwordField = screen.getByPlaceholderText(/password/i)
		const signInButton = screen.getByRole('button', { name: /sign in/i })
		await user.type(emailField, 'test@email.com')
		await user.type(passwordField, 'password')
		await user.click(signInButton)
		expect(signInButton).toBeDisabled()
		await screen.findByText(/loading/i)
	})

	test('should have sign up link', async () => {
		render(<SignInPage />)
		const signInLink = screen.getByRole('link', { name: /sign up/i })
		expect(signInLink).toBeInTheDocument()
	})
})
