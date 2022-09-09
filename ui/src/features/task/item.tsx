import clsx from 'clsx'
import { Task } from '../../api'
import { Linkified } from '../ui'
import { TaskCompletion } from './completion'
import { TaskDeletion } from './deletion'

interface TaskItemProps {
	task: Task
}

export function TaskItem({ task }: TaskItemProps) {
	return (
		<div>
			<div className="flex items-center justify-between gap-10">
				<h4
					className={clsx(task.isDone && 'line-through text-gray-500')}
					data-testid="task-title"
				>
					<Linkified>{task.title}</Linkified>
				</h4>
				<div className="flex items-center gap-2">
					<TaskCompletion taskId={task.id} isDone={task.isDone} />
					<TaskDeletion taskId={task.id} />
				</div>
			</div>
			<p className="mt-2 text-xs">{task.description}</p>
		</div>
	)
}
