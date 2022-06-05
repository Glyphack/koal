import { Task } from '../../api'
import { TaskItem } from './item'

interface TaskListProps {
	tasks: Task[]
}

export function TaskList({ tasks }: TaskListProps) {
	return (
		<div>
			{tasks.map((task) => (
				<TaskItem key={task.id} task={task} />
			))}
		</div>
	)
}
