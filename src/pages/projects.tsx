import { ProjectForm, ProjectList } from '../features/project'

export function ProjectsPage() {
	return (
		<main className="space-y-20">
			<ProjectForm />
			<ProjectList />
		</main>
	)
}
