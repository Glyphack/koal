import { render, screen } from '../../test/test-utils'
import { NotFoundPage } from '../404'

describe('NotFoundPage', () => {
	it('should render', () => {
		render(<NotFoundPage />)
		const statusText = screen.getByText(/404/i)
		const notFoundText = screen.getByText(/page not found/i)
		const homeLink = screen.getByRole('link', { name: /home/i })
		expect(statusText).toBeInTheDocument()
		expect(notFoundText).toBeInTheDocument()
		expect(homeLink).toBeInTheDocument()
	})
})
