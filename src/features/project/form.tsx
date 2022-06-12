import { useForm } from 'react-hook-form'
import { useMutation, useQueryClient } from 'react-query'
import { api, QueryKey } from '../../api'
import { Button, Field } from '../ui'

export function ProjectForm() {
	const queryClient = useQueryClient()
	const form = useForm({ defaultValues: { name: '' } })
	const mutation = useMutation(api.createProject, {
		onSuccess: () => {
			queryClient.invalidateQueries(QueryKey.Projects)
			form.reset()
		},
	})
	const createProject = form.handleSubmit((values) => mutation.mutate(values))

	return (
		<form className="space-y-4" onSubmit={createProject}>
			<Field type="text" placeholder="Project Name" {...form.register('name')} />
			<Button type="submit" disabled={mutation.isLoading}>
				Add Project
			</Button>
		</form>
	)
}
