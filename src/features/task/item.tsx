import { Task } from '../../api'
import { TaskCompletion } from './completion'
import { TaskDeletion } from './deletion'

interface TaskItemProps {
	task: Task
}

export function TaskItem({ task }: TaskItemProps) {
	return (
		<div className="flex items-center justify-between gap-10">
			<h4>{task.title}</h4>
			<div className="flex items-center gap-6">
				<TaskCompletion taskId={task.id} isDone={task.isDone} />
				<TaskDeletion taskId={task.id} />
			</div>
		</div>
	)
}
