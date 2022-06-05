import { Task } from '../../api'
import { TaskCompletion } from './completion'
import { TaskDeletion } from './deletion'

interface TaskItemProps {
	task: Task
}

export function TaskItem({ task }: TaskItemProps) {
	return (
		<div>
			<span>{task.title}</span>
			<TaskCompletion taskId={task.id} isDone={task.isDone} />
			<TaskDeletion taskId={task.id} />
		</div>
	)
}
