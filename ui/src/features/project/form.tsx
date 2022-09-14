import { Button, TextInput } from '@mantine/core'
import { useForm } from '@mantine/form'
import { useMutation, useQueryClient } from 'react-query'
import { api, QueryKey } from '../../api'
import { Form } from '../ui'

export function ProjectForm() {
	const queryClient = useQueryClient()
	const form = useForm({ initialValues: { name: '' } })
	const mutation = useMutation(api.createProject, {
		onSuccess: () => {
			queryClient.invalidateQueries([QueryKey.Projects])
			form.reset()
		},
	})
	const createProject = form.onSubmit((values) => mutation.mutate(values))

	return (
		<Form
			onSubmit={createProject}
			fields={
				<TextInput
					label="Project"
					type="text"
					placeholder="project name"
					{...form.getInputProps('name')}
				/>
			}
			actions={
				<Button type="submit" loading={mutation.isLoading}>
					Add Project
				</Button>
			}
		></Form>
	)
}
