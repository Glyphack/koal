import { mockTask } from '../../../test/handlers'
import { render, screen } from '../../../test/test-utils'
import { TaskItem } from '../item'

describe('TaskItem', () => {
	test('should render task', async () => {
		render(<TaskItem task={mockTask} />)
		const title = screen.getByTestId('task-title')
		const completion = screen.getByTestId('task-completion')
		const deletion = screen.getByTestId('task-deletion')
		expect(title).toHaveTextContent(mockTask.title)
		expect(completion).toBeInTheDocument()
		expect(deletion).toBeInTheDocument()
	})
})
