import { Task } from '../../api'
import { TaskCompletion } from './completion'
import { TaskDeletion } from './deletion'
import { TaskDescription } from './description'
import { TaskTitle } from './title'

interface TaskItemProps {
	task: Task
}

export function TaskItem({ task }: TaskItemProps) {
	return (
		<div className="flex justify-between gap-10">
			<div className="w-full">
				<TaskTitle task={task} />
				<TaskDescription task={task} />
			</div>
			<div className="flex gap-2">
				<TaskCompletion taskId={task.id} isDone={task.isDone} />
				<TaskDeletion taskId={task.id} />
			</div>
		</div>
	)
}
