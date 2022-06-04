import { render, screen, waitFor } from '../../../test/test-utils'
import { ProjectForm } from '../form'

describe('ProjectForm', () => {
	test('should add project', async () => {
		const { user } = render(<ProjectForm />)
		const projectNameField = screen.getByPlaceholderText(/project name/i)
		const submitButton = screen.getByRole('button', { name: /add project/i })
		await user.type(projectNameField, 'Test Project')
		await user.click(submitButton)
		expect(submitButton).toBeDisabled()
		await waitFor(() => expect(submitButton).toBeEnabled())
		expect(projectNameField).not.toHaveValue()
	})
})
