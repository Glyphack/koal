import { render, screen } from '../../../test/test-utils'
import { SignUpForm } from '../sign-up-form'

describe('SignUpForm', () => {
	test('should be able to sign up', async () => {
		const { user } = render(<SignUpForm />)
		const emailField = screen.getByPlaceholderText(/email/i)
		const passwordField = screen.getByPlaceholderText(/password/i)
		const signUpButton = screen.getByRole('button', { name: /sign up/i })
		await user.type(emailField, 'test@email.com')
		await user.type(passwordField, 'password')
		await user.click(signUpButton)
		expect(signUpButton).toBeDisabled()
		await screen.findByTestId('loader')
	})
})
