import { mockProject } from '../../test/handlers'
import { render, screen, waitFor } from '../../test/test-utils'
import { ProjectsPage } from '../projects'

describe('ProjectsPage', () => {
	it('should render', async () => {
		render(<ProjectsPage />)
		const loader = screen.getByTestId('loader')
		expect(loader).toBeInTheDocument()
		await waitFor(() => expect(loader).not.toBeInTheDocument())
		const addProjectButton = screen.getByRole('button', { name: /add project/i })
		expect(addProjectButton).toBeInTheDocument()
		const projectTitle = screen.getByText(mockProject.name)
		expect(projectTitle).toBeInTheDocument()
	})
})
