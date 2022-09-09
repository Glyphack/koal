import { useQuery } from 'react-query'
import { api, QueryKey } from '../../api'
import { Link, Loader } from '../ui'

export function ProjectList() {
	const projectsQuery = useQuery(QueryKey.Projects, api.fetchProjects)
	if (projectsQuery.isLoading) return <Loader />
	const projects = projectsQuery.data?.data.projects ?? []

	if (projects.length === 0) return <p>Your projects will show up here 📜</p>

	return (
		<div className="space-y-8">
			{projects.map((project) => (
				<div key={project.id}>
					<Link className="text-3xl font-normal" to={`/project/${project.id}`}>
						{project.name}
					</Link>
				</div>
			))}
		</div>
	)
}
