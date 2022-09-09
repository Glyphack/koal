import { Button, Textarea, TextInput } from '@mantine/core'
import { useForm } from '@mantine/form'
import { useMutation, useQueryClient } from 'react-query'
import { api, QueryKey } from '../../api'
import { Form } from '../ui'

interface TaskFormProps {
	projectId: string
}

export function TaskForm({ projectId }: TaskFormProps) {
	const queryClient = useQueryClient()
	const taskForm = useForm({ initialValues: { title: '', description: '' } })
	const createTaskMutation = useMutation(api.createTask, {
		onSuccess: () => {
			queryClient.invalidateQueries([QueryKey.Project])
			taskForm.reset()
		},
	})
	const createTask = taskForm.onSubmit((values) =>
		createTaskMutation.mutate({ projectId, ...values })
	)

	return (
		<Form
			onSubmit={createTask}
			fields={
				<>
					<TextInput
						label="Task"
						placeholder="task name"
						{...taskForm.getInputProps('title')}
					/>
					<Textarea
						label="Description"
						placeholder="task description"
						{...taskForm.getInputProps('description')}
					/>
				</>
			}
			actions={
				<Button type="submit" loading={createTaskMutation.isLoading}>
					Add Task
				</Button>
			}
		></Form>
	)
}
