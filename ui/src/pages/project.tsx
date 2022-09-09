import { useQuery } from 'react-query'
import { useParams } from 'react-router-dom'
import { api, QueryKey } from '../api'
import { ProjectDeletion } from '../features/project'
import { TaskForm, TaskList } from '../features/task'
import { Loader } from '../features/ui'
import { checkNotFound } from '../utils'
import { NotFoundPage } from './404'

export function ProjectPage() {
	const { id: projectId = '' } = useParams<{ id: string }>()
	return <ProjectDetails projectId={projectId} />
}

interface ProjectDetailsProps {
	projectId: string
}

export function ProjectDetails({ projectId }: ProjectDetailsProps) {
	const projectQuery = useQuery([QueryKey.Project, projectId], () => api.fetchProject(projectId))
	const project = projectQuery.data?.data

	if (checkNotFound(projectQuery.error)) return <NotFoundPage />
	if (projectQuery.isLoading || !project) return <Loader />

	return (
		<main className="space-y-20">
			<div className="flex items-center justify-between max-w-xs">
				<h3 className="text-3xl">{project.info.name}</h3>
				<ProjectDeletion projectId={projectId} />
			</div>
			<div className="flex flex-col gap-20 lg:gap-40 lg:flex-row">
				<div className="w-[20rem] shrink-0">
					<TaskForm projectId={projectId} />
				</div>
				<div className="grow">
					<TaskList tasks={project.items} />
				</div>
			</div>
		</main>
	)
}
