import { render, screen } from '../../../test/test-utils'
import { Button } from '../button'

describe('Button', () => {
	test('should render', () => {
		render(<Button>boom</Button>)
		const button = screen.getByRole('button', { name: /boom/i })
		expect(button).toBeInTheDocument()
	})

	test('should render loader when loading', () => {
		render(<Button loading>boom</Button>)
		const button = screen.getByRole('button')
		const loader = screen.getByTestId('loader')
		expect(button).toBeDisabled()
		expect(loader).toBeInTheDocument()
	})
})
