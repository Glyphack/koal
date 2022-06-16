import { HiOutlineTrash } from 'react-icons/hi'
import { useMutation } from 'react-query'
import { useNavigate } from 'react-router-dom'
import { api } from '../../api'
import { Button } from '../ui'

interface ProjectDeletionProps {
	projectId: string
}

export function ProjectDeletion({ projectId }: ProjectDeletionProps) {
	const navigate = useNavigate()
	const deleteProjectMutation = useMutation(api.deleteProject, {
		onSuccess: () => navigate('/project', { replace: true }),
	})
	const deleteProject = () => deleteProjectMutation.mutate(projectId)

	return (
		<Button
			variant="icon"
			type="button"
			onClick={deleteProject}
			loading={deleteProjectMutation.isLoading}
			data-testid="project-deletion"
		>
			<HiOutlineTrash />
		</Button>
	)
}
