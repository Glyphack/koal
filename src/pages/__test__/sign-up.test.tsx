import { render, screen } from '../../test/test-utils'
import { SignUpPage } from '../sign-up'

describe('SignUpPage', () => {
	test('should be able to sign up', async () => {
		const { user } = render(<SignUpPage />)
		const emailField = screen.getByPlaceholderText(/email/i)
		const passwordField = screen.getByPlaceholderText(/password/i)
		const signUpButton = screen.getByRole('button', { name: /sign up/i })
		await user.type(emailField, 'test@email.com')
		await user.type(passwordField, 'password')
		await user.click(signUpButton)
		expect(signUpButton).toBeDisabled()
		await screen.findByText(/loading/i)
	})

	test('should have sign in link', async () => {
		render(<SignUpPage />)
		const signInLink = screen.getByRole('link', { name: /sign in/i })
		expect(signInLink).toBeInTheDocument()
	})
})
