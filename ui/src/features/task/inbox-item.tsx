import { Task } from '../../api'
import { TaskCompletion } from './completion'
import { TaskDeletion } from './deletion'
import { TaskDescription } from './description'
import { TaskTitle } from './title'

interface InboxTaskItemProps {
	task: Task
}

export function InboxTaskItem({ task }: InboxTaskItemProps) {
	return (
		<div className="flex justify-between gap-10">
			<div className="w-full">
				<div className="flex items-center gap-2">
					<TaskTitle task={task} />
					<p className="px-2 py-0.5 text-xs text-gray-500 rounded-full bg-gray-50">
						{task.project.name}
					</p>
				</div>
				<TaskDescription task={task} />
			</div>
			<div className="flex gap-2">
				<TaskCompletion taskId={task.id} isDone={task.isDone} />
				<TaskDeletion taskId={task.id} />
			</div>
		</div>
	)
}
