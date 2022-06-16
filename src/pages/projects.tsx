import { ProjectForm, ProjectList } from '../features/project'
import { Link } from '../features/ui'

export function ProjectsPage() {
	return (
		<main className="flex flex-col gap-28 lg:flex-row lg:gap-40">
			<div className="max-w-xs grow">
				<ProjectForm />
			</div>
			<div className="flex flex-col gap-10 grow">
				<Link className="text-3xl font-normal" to="/inbox">
					Inbox
				</Link>
				<hr />
				<ProjectList />
			</div>
		</main>
	)
}
