import { Task } from '../../api'
import { TaskItem } from './item'

interface TaskListProps {
	tasks: Task[]
}

export function TaskList({ tasks }: TaskListProps) {
	return (
		<div className="flex flex-col gap-4">
			{tasks.map((task) => (
				<TaskItem key={task.id} task={task} />
			))}
		</div>
	)
}
