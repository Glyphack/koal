import { ProjectForm, ProjectList } from '../features/project'

export function ProjectsPage() {
	return (
		<main className="flex flex-col gap-28 lg:flex-row lg:gap-40">
			<div className="max-w-xs grow">
				<ProjectForm />
			</div>
			<ProjectList />
		</main>
	)
}
