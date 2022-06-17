import { mockTasks } from '../../../test/handlers'
import { render, screen } from '../../../test/test-utils'
import { InboxTaskList } from '../inbox-list'

describe('InboxTaskList', () => {
	test('should render inbox task list', async () => {
		render(<InboxTaskList tasks={mockTasks} />)
		const titles = screen.getAllByTestId('task-title')
		const title1 = titles[0]
		const title2 = titles[1]
		const taskProject1 = screen.getByText(mockTasks[0].project.name)
		const taskProject2 = screen.getByText(mockTasks[1].project.name)
		const completion = screen.getAllByTestId('task-completion')
		const deletion = screen.getAllByTestId('task-deletion')
		expect(title1).toHaveTextContent(mockTasks[0].title)
		expect(title2).toHaveTextContent(mockTasks[1].title)
		expect(taskProject1).toBeInTheDocument()
		expect(taskProject2).toBeInTheDocument()
		expect(completion.length).toBe(2)
		expect(deletion.length).toBe(2)
	})
})
