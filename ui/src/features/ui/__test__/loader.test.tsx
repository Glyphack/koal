import { render, screen } from '../../../test/test-utils'
import { Loader } from '../loader'

describe('Loader', () => {
	test('renders null if no loading', () => {
		render(<Loader loading={false} />)
		const loader = screen.queryByTestId('loader')
		expect(loader).not.toBeInTheDocument()
	})

	test('renders loader', () => {
		render(<Loader />)
		const loader = screen.getByTestId('loader')
		expect(loader).toBeInTheDocument()
	})
})
