import { render, screen, waitFor } from '../../../test/test-utils'
import { ProjectDeletion } from '../deletion'

describe('ProjectDeletion', () => {
	test('should delete project', async () => {
		const { user } = render(<ProjectDeletion projectId="p1" />)
		const deleteButton = screen.getByTestId('project-deletion')
		await user.click(deleteButton)
		expect(deleteButton).toBeDisabled()
		const loader = await screen.findByTestId('loader')
		expect(loader).toBeInTheDocument()
		await waitFor(() => expect(loader).not.toBeInTheDocument())
	})
})
