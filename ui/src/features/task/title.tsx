import { useInputState } from '@mantine/hooks'
import clsx from 'clsx'
import { useMutation } from 'react-query'
import { api, Task } from '../../api'
import { Linkified } from '../ui'
import { useEditable } from './use-editable'

interface TaskTitleProps {
	task: Task
}

export function TaskTitle({ task }: TaskTitleProps) {
	const titleEditable = useEditable<HTMLInputElement>()
	const updateTaskMutation = useMutation(api.updateTask)
	const [newTitle, setNewTitle] = useInputState(task.title)

	const form = (
		<form
			onSubmit={(event) => {
				event.preventDefault()
				titleEditable.inputRef.current?.blur()
			}}
			className="w-full"
		>
			<input
				ref={titleEditable.inputRef}
				value={newTitle}
				onChange={setNewTitle}
				className="w-full outline-none"
				onBlur={() => {
					titleEditable.setIsEditing(false)
					updateTaskMutation.mutate({
						id: task.id,
						payload: { title: newTitle },
					})
				}}
			/>
		</form>
	)

	const text = (
		<h4
			className={clsx(task.isDone && 'line-through text-gray-500')}
			data-testid="task-title"
			onFocus={() => titleEditable.setIsEditing(true)}
			tabIndex={0}
		>
			<Linkified>{newTitle}</Linkified>
		</h4>
	)

	return (
		<div>
			{titleEditable.isEditing && form}
			{!titleEditable.isEditing && text}
		</div>
	)
}
