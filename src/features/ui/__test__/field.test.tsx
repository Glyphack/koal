import { render, screen } from '../../../test/test-utils'
import { Field } from '../field'

describe('Field', () => {
	test('should render', () => {
		render(<Field label="age" />)
		const field = screen.getByRole('textbox')
		const label = screen.getByText(/age/i)
		expect(field).toBeInTheDocument()
		expect(label).toBeInTheDocument()
	})
})
