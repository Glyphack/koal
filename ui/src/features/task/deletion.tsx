import { HiOutlineTrash } from 'react-icons/hi'
import { useMutation, useQueryClient } from 'react-query'
import { api, QueryKey } from '../../api'
import { Button } from '../ui'

interface TaskDeletionProps {
	taskId: string
}

export function TaskDeletion({ taskId }: TaskDeletionProps) {
	const queryClient = useQueryClient()
	const deleteTaskMutation = useMutation(api.deleteTask, {
		onSuccess: () => {
			queryClient.invalidateQueries([QueryKey.Project])
			queryClient.invalidateQueries([QueryKey.Inbox])
		},
	})
	const deleteTask = () => deleteTaskMutation.mutate(taskId)

	return (
		<Button
			variant="icon"
			type="button"
			onClick={deleteTask}
			loading={deleteTaskMutation.isLoading}
			data-testid="task-deletion"
		>
			<HiOutlineTrash />
		</Button>
	)
}
