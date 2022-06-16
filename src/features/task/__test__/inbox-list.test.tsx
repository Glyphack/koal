import { mockTasks } from '../../../test/handlers'
import { render, screen } from '../../../test/test-utils'
import { InboxTaskList } from '../inbox-list'

describe('InboxTaskList', () => {
	test('should render inbox task list', async () => {
		render(<InboxTaskList tasks={mockTasks} />)
		const title1 = screen.getByText(mockTasks[0].title)
		const title2 = screen.getByText(mockTasks[1].title)
		const taskProject1 = screen.getByText(mockTasks[0].project.name)
		const taskProject2 = screen.getByText(mockTasks[1].project.name)
		const completion = screen.getAllByTestId('task-completion')
		const deletion = screen.getAllByTestId('task-deletion')
		expect(title1).toBeInTheDocument()
		expect(title2).toBeInTheDocument()
		expect(taskProject1).toBeInTheDocument()
		expect(taskProject2).toBeInTheDocument()
		expect(completion.length).toBe(2)
		expect(deletion.length).toBe(2)
	})
})
