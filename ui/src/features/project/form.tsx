import { useForm } from 'react-hook-form'
import { useMutation, useQueryClient } from 'react-query'
import { api, QueryKey } from '../../api'
import { Button, Field, Form } from '../ui'

export function ProjectForm() {
	const queryClient = useQueryClient()
	const form = useForm({ defaultValues: { name: '' } })
	const mutation = useMutation(api.createProject, {
		onSuccess: () => {
			queryClient.invalidateQueries([QueryKey.Projects])
			form.reset()
		},
	})
	const createProject = form.handleSubmit((values) => mutation.mutate(values))

	return (
		<Form
			onSubmit={createProject}
			fields={
				<Field
					label="Project"
					type="text"
					placeholder="project name"
					{...form.register('name')}
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
