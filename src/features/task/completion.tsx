import { HiCheck, HiX } from 'react-icons/hi'
import { useMutation, useQueryClient } from 'react-query'
import { api, QueryKey } from '../../api'
import { Button } from '../ui'

interface TaskCompletionProps {
	taskId: string
	isDone: boolean
}

export function TaskCompletion({ taskId, isDone }: TaskCompletionProps) {
	const queryClient = useQueryClient()
	const updateTaskMutation = useMutation(api.updateTask, {
		onSuccess: () => queryClient.invalidateQueries(QueryKey.Project),
	})
	const checkTask = () => updateTaskMutation.mutate({ id: taskId, payload: { isDone: true } })
	const uncheckTask = () => updateTaskMutation.mutate({ id: taskId, payload: { isDone: false } })
	const onClick = isDone ? uncheckTask : checkTask
	const content = isDone ? <HiX /> : <HiCheck />

	return (
		<Button
			variant="icon"
			type="button"
			onClick={onClick}
			loading={updateTaskMutation.isLoading}
			data-testid="task-completion"
		>
			{content}
		</Button>
	)
}
