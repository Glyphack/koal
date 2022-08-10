import { render, screen } from '../../test/test-utils'
import { SignUpPage } from '../sign-up'

describe('SignUpPage', () => {
	test('should have sign in link', async () => {
		render(<SignUpPage />)
		const signInLink = screen.getByRole('link', { name: /sign in/i })
		expect(signInLink).toBeInTheDocument()
	})
})
