import { render, screen, waitFor } from '../../../test/test-utils'
import { ProjectList } from '../list'

describe('ProjectList', () => {
	test('should render projects', async () => {
		render(<ProjectList />)
		const loader = screen.getByTestId('loader')
		expect(loader).toBeInTheDocument()
		await waitFor(() => expect(loader).not.toBeInTheDocument())
		const project = screen.getByRole('link', { name: 'Test Project' })
		expect(project).toBeInTheDocument()
	})
})
