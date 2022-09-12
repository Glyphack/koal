import { Textarea } from '@mantine/core'
import { useInputState } from '@mantine/hooks'
import { useMutation } from 'react-query'
import { api, Task } from '../../api'
import { Linkified } from '../ui'
import { useEditable } from './use-editable'

interface TaskDescriptionProps {
	task: Task
}

export function TaskDescription({ task }: TaskDescriptionProps) {
	const updateTaskMutation = useMutation(api.updateTask)
	const descriptionEditable = useEditable<HTMLTextAreaElement>()
	const [newDescription, setNewDescription] = useInputState(task.description)

	const form = (
		<form
			onSubmit={(event) => {
				event.preventDefault()
				descriptionEditable.inputRef.current?.blur()
			}}
		>
			<Textarea
				autosize
				unstyled
				ref={descriptionEditable.inputRef}
				value={newDescription}
				onChange={setNewDescription}
				classNames={{ input: 'w-full outline-none resize-none !p-0' }}
				onBlur={() => {
					descriptionEditable.setIsEditing(false)
					updateTaskMutation.mutate({
						id: task.id,
						payload: { description: newDescription },
					})
				}}
			/>
		</form>
	)

	const text = (
		<p
			className="whitespace-pre-line min-h-[4rem]"
			onFocus={() => descriptionEditable.setIsEditing(true)}
			tabIndex={0}
		>
			<Linkified>{newDescription}</Linkified>
		</p>
	)

	return (
		<div className="mt-2 text-xs">
			{descriptionEditable.isEditing && form}
			{!descriptionEditable.isEditing && text}
		</div>
	)
}
