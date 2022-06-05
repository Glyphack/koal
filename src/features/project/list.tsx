import { useQuery } from 'react-query'
import { Link } from 'react-router-dom'
import { api, QueryKey } from '../../api'
import { Loader } from '../ui'

export function ProjectList() {
	const projectsQuery = useQuery(QueryKey.Projects, api.fetchProjects)
	if (projectsQuery.isLoading) return <Loader />
	const projects = projectsQuery.data?.data.projects ?? []

	return (
		<div>
			{projects.map((project) => (
				<div key={project.id}>
					<Link to={`/project/${project.id}`}>{project.name}</Link>
				</div>
			))}
		</div>
	)
}
