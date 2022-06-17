import clsx from 'clsx'
import { Task } from '../../api'
import { Linkified } from '../ui'
import { TaskCompletion } from './completion'
import { TaskDeletion } from './deletion'

interface InboxTaskItemProps {
	task: Task
}

export function InboxTaskItem({ task }: InboxTaskItemProps) {
	return (
		<div className="flex items-center justify-between gap-10">
			<div>
				<h4
					className={clsx(task.isDone && 'line-through text-gray-500')}
					data-testid="task-title"
				>
					<Linkified>{task.title}</Linkified>
				</h4>
				<h5 className="text-xs text-gray-500">{task.project.name}</h5>
			</div>
			<div className="flex items-center gap-6">
				<TaskCompletion taskId={task.id} isDone={task.isDone} />
				<TaskDeletion taskId={task.id} />
			</div>
		</div>
	)
}
