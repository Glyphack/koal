import { useForm } from 'react-hook-form'
import { useMutation, useQueryClient } from 'react-query'
import { api, QueryKey } from '../../api'
import { Button, Field } from '../ui'

interface TaskFormProps {
	projectId: string
}

export function TaskForm({ projectId }: TaskFormProps) {
	const queryClient = useQueryClient()
	const taskForm = useForm({ defaultValues: { title: '' } })
	const createTaskMutation = useMutation(api.createTask, {
		onSuccess: () => {
			queryClient.invalidateQueries(QueryKey.Project)
			taskForm.reset()
		},
	})
	const createTask = taskForm.handleSubmit((values) =>
		createTaskMutation.mutate({ projectId, title: values.title })
	)

	return (
		<form onSubmit={createTask}>
			<Field type="text" placeholder="Task Name" {...taskForm.register('title')} />
			<Button type="submit" disabled={createTaskMutation.isLoading}>
				Add Task
			</Button>
		</form>
	)
}
