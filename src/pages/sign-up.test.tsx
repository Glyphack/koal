import { test } from 'vitest'
import { render, screen } from '../test/test-utils'
import { SignUpPage } from './sign-up'

describe('SignUpPage', () => {
	test('should be able to sign up', async () => {
		const { user } = render(<SignUpPage />)
		const emailField = screen.getByPlaceholderText(/email/i)
		const passwordField = screen.getByPlaceholderText(/password/i)
		const signUpButton = screen.getByRole('button', { name: /sign up/i })
		user.type(emailField, 'test@email.com')
		user.type(passwordField, 'password')
		await user.click(signUpButton)
		await screen.findByText(/submitting/i)
	})
})
