import { Jelly } from '@uiball/loaders'

interface LoaderProps {
	loading?: boolean
}

export function Loader({ loading = true }: LoaderProps) {
	if (!loading) return null

	return (
		<div data-testid="loader" className="flex items-center justify-center h-6">
			<Jelly size={40} color="#064e3b" />
		</div>
	)
}
