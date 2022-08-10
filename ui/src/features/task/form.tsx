import { useForm } from 'react-hook-form'
import { useMutation, useQueryClient } from 'react-query'
import { api, QueryKey } from '../../api'
import { Button, Field, Form } from '../ui'

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
		<Form
			onSubmit={createTask}
			fields={
				<Field
					label="Task"
					type="text"
					placeholder="task name"
					{...taskForm.register('title')}
				/>
			}
			actions={
				<Button type="submit" loading={createTaskMutation.isLoading}>
					Add Task
				</Button>
			}
		></Form>
	)
}
