import { render, screen, waitFor } from '../../../test/test-utils'
import { TaskForm } from '../form'

describe('TaskForm', () => {
	test('task form should work', async () => {
		const { user } = render(<TaskForm projectId="1" />)
		const taskNameField = screen.getByPlaceholderText(/task name/i)
		const addTaskButton = screen.getByRole('button', { name: /add task/i })
		await user.type(taskNameField, 'task 1')
		await user.click(addTaskButton)
		expect(addTaskButton).toBeDisabled()
		await waitFor(() => expect(addTaskButton).toBeEnabled())
		expect(taskNameField).not.toHaveValue()
	})
})
