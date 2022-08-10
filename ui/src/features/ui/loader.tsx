import { Jelly } from '@uiball/loaders'

interface LoaderProps {
	loading?: boolean
	size?: number
}

export function Loader({ loading = true, size = 36 }: LoaderProps) {
	if (!loading) return null

	return (
		<div data-testid="loader" className="flex items-center justify-center h-6">
			<Jelly size={size} color="#064e3b" />
		</div>
	)
}
