import { mockTask } from '../../test/handlers'
import { render, screen, waitFor } from '../../test/test-utils'
import { InboxPage } from '../inbox'

describe('InboxPage', () => {
	it('should render inbox page', async () => {
		render(<InboxPage />)
		const loader = screen.getByTestId('loader')
		expect(loader).toBeInTheDocument()
		await waitFor(() => expect(loader).not.toBeInTheDocument())
		const pageTitle = screen.getByText(/inbox/i)
		expect(pageTitle).toBeInTheDocument()
		const taskName = screen.getByText(mockTask.title)
		expect(taskName).toBeInTheDocument()
	})
})
