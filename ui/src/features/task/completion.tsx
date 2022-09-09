import { ActionIcon } from '@mantine/core'
import { HiCheck } from 'react-icons/hi'
import { useMutation, useQueryClient } from 'react-query'
import { api, QueryKey } from '../../api'

interface TaskCompletionProps {
	taskId: string
	isDone: boolean
}

export function TaskCompletion({ taskId, isDone }: TaskCompletionProps) {
	const queryClient = useQueryClient()
	const updateTaskMutation = useMutation(api.updateTask, {
		onSuccess: () => {
			queryClient.invalidateQueries([QueryKey.Project])
			queryClient.invalidateQueries([QueryKey.Inbox])
		},
	})
	const checkTask = () => updateTaskMutation.mutate({ id: taskId, payload: { isDone: true } })
	const uncheckTask = () => updateTaskMutation.mutate({ id: taskId, payload: { isDone: false } })
	const onClick = isDone ? uncheckTask : checkTask
	const content = isDone ? <HiCheck /> : <span />

	return (
		<ActionIcon onClick={onClick} loading={updateTaskMutation.isLoading}>
			<span className="flex items-center justify-center w-4 h-4 border-2 border-gray-400 rounded-full">
				{content}
			</span>
		</ActionIcon>
	)
}
