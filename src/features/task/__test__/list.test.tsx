import { mockTasks } from '../../../test/handlers'
import { render, screen } from '../../../test/test-utils'
import { TaskList } from '../list'

describe('TaskList', () => {
	test('should render task list', async () => {
		render(<TaskList tasks={mockTasks} />)
		const title1 = screen.getByText(mockTasks[0].title)
		const title2 = screen.getByText(mockTasks[1].title)
		const completion = screen.getAllByTestId('task-completion')
		const deletion = screen.getAllByTestId('task-deletion')
		expect(title1).toBeInTheDocument()
		expect(title2).toBeInTheDocument()
		expect(completion.length).toBe(2)
		expect(deletion.length).toBe(2)
	})
})
