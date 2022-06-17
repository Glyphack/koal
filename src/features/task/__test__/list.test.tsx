import { mockTasks } from '../../../test/handlers'
import { render, screen } from '../../../test/test-utils'
import { TaskList } from '../list'

describe('TaskList', () => {
	test('should render task list', () => {
		render(<TaskList tasks={mockTasks} />)
		const tasks = screen.getAllByTestId('task-title')
		const title1 = tasks[0]
		const title2 = tasks[1]
		const completion = screen.getAllByTestId('task-completion')
		const deletion = screen.getAllByTestId('task-deletion')
		expect(title1).toHaveTextContent(mockTasks[0].title)
		expect(title2).toHaveTextContent(mockTasks[1].title)
		expect(completion.length).toBe(2)
		expect(deletion.length).toBe(2)
	})
})
