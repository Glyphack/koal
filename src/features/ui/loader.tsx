interface LoaderProps {
	loading?: boolean
}

export function Loader({ loading = true }: LoaderProps) {
	if (!loading) return null

	return <p data-testid="loader">Loading...</p>
}
