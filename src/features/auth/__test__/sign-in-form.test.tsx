import { render, screen } from '../../../test/test-utils'
import { SignInForm } from '../sign-in-form'

describe('SignInForm', () => {
	test('should be able to sign in', async () => {
		const { user } = render(<SignInForm />)
		const emailField = screen.getByPlaceholderText(/email/i)
		const passwordField = screen.getByPlaceholderText(/password/i)
		const signInButton = screen.getByRole('button', { name: /sign in/i })
		await user.type(emailField, 'test@email.com')
		await user.type(passwordField, 'password')
		await user.click(signInButton)
		expect(signInButton).toBeDisabled()
		await screen.findByText(/loading/i)
	})
})
