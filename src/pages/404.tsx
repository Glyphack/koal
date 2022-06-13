import { Link } from '../features/ui'

export function NotFoundPage() {
	return (
		<main className="flex flex-col items-center gap-6">
			<h2 className="font-thin text-9xl">404</h2>
			<h3 className="text-4xl font-thin">Page Not Found</h3>
			<Link className="mt-10 text-xl" to="/project">
				Home
			</Link>
		</main>
	)
}
