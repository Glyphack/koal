import { cleanup, render } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { ReactElement } from 'react'
import { QueryClient, QueryClientProvider } from 'react-query'
import { afterEach } from 'vitest'

afterEach(() => cleanup())

const queryClient = new QueryClient({ defaultOptions: { queries: { retry: false } } })

const customRender = (ui: ReactElement, options = {}) => ({
	user: userEvent.setup(),
	...render(ui, {
		// wrap provider(s) here if needed
		wrapper: ({ children }) => (
			<QueryClientProvider client={queryClient}>{children}</QueryClientProvider>
		),
		...options,
	}),
})

export * from '@testing-library/react'
// override render export
export { customRender as render }
export default userEvent
