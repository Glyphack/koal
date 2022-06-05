import { useQuery } from 'react-query'
import { useParams } from 'react-router-dom'
import { api, QueryKey } from '../api'
import { ProjectDeletion } from '../features/project'
import { TaskForm, TaskList } from '../features/task'
import { Loader } from '../features/ui'

export function ProjectPage() {
	const { id: projectId = '' } = useParams<{ id: string }>()
	const projectQuery = useQuery([QueryKey.Project, projectId], () => api.fetchProject(projectId))
	const project = projectQuery.data?.data

	if (projectQuery.isLoading || !project) return <Loader />
	return (
		<main>
			<div>
				<span>{project.info.name}</span>
				<ProjectDeletion projectId={projectId} />
			</div>
			<TaskForm projectId={projectId} />
			<TaskList tasks={project.items} />
		</main>
	)
}
