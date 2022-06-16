import { mockTask } from '../../../test/handlers'
import { render, screen } from '../../../test/test-utils'
import { InboxTaskItem } from '../inbox-item'

describe('InboxTaskItem', () => {
	test('should render inbox task', async () => {
		render(<InboxTaskItem task={mockTask} />)
		const title = screen.getByText(mockTask.title)
		const project = screen.getByText(mockTask.project.name)
		const completion = screen.getByTestId('task-completion')
		const deletion = screen.getByTestId('task-deletion')
		expect(title).toBeInTheDocument()
		expect(project).toBeInTheDocument()
		expect(completion).toBeInTheDocument()
		expect(deletion).toBeInTheDocument()
	})
})
