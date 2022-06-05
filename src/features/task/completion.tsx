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
	const text = isDone ? 'Undone' : 'Done'

	return (
		<Button type="button" onClick={onClick} disabled={updateTaskMutation.isLoading}>
			{text}
		</Button>
	)
}
