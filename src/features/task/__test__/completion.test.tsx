import { render, screen, waitFor } from '../../../test/test-utils'
import { TaskCompletion } from '../completion'

describe('TaskCompletion', () => {
	test('should done task', async () => {
		const { user } = render(<TaskCompletion isDone={false} taskId="1" />)
		const doneTaskButton = screen.getByTestId('task-completion')
		await user.click(doneTaskButton)
		expect(doneTaskButton).toBeDisabled()
		await waitFor(() => expect(doneTaskButton).toBeEnabled())
	})

	test('should undone task', async () => {
		const { user } = render(<TaskCompletion isDone={true} taskId="1" />)
		const undoneTaskButton = screen.getByTestId('task-completion')
		await user.click(undoneTaskButton)
		expect(undoneTaskButton).toBeDisabled()
		await waitFor(() => expect(undoneTaskButton).toBeEnabled())
	})
})
