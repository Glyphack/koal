import { useMutation, useQueryClient } from 'react-query'
import { api, QueryKey } from '../../api'
import { Button } from '../ui'

interface TaskDeletionProps {
	taskId: string
}

export function TaskDeletion({ taskId }: TaskDeletionProps) {
	const queryClient = useQueryClient()
	const deleteTaskMutation = useMutation(api.deleteTask, {
		onSuccess: () => queryClient.invalidateQueries(QueryKey.Project),
	})
	const deleteTask = () => deleteTaskMutation.mutate(taskId)

	return (
		<Button type="button" onClick={deleteTask} disabled={deleteTaskMutation.isLoading}>
			Delete
		</Button>
	)
}
