import { Task } from '../../api'
import { InboxTaskItem } from './inbox-item'

interface InboxTaskListProps {
	tasks: Task[]
}

export function InboxTaskList({ tasks }: InboxTaskListProps) {
	if (tasks.length === 0)
		return <p className="text-gray-600">Congrats, you have no incomplete task 😊</p>

	return (
		<div className="flex flex-col gap-6">
			{tasks.map((task) => (
				<div key={task.id} className="flex flex-col">
					<InboxTaskItem task={task} />
				</div>
			))}
		</div>
	)
}
