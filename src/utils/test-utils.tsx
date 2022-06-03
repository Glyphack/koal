/* eslint-disable import/export */
import { cleanup, render } from '@testing-library/react'
import { ReactElement } from 'react'
import { QueryClient, QueryClientProvider } from 'react-query'
import { afterEach } from 'vitest'

afterEach(() => cleanup())

const queryClient = new QueryClient()

const customRender = (ui: ReactElement, options = {}) =>
	render(ui, {
		// wrap provider(s) here if needed
		wrapper: ({ children }) => (
			<QueryClientProvider client={queryClient}>{children}</QueryClientProvider>
		),
		...options,
	})

export * from '@testing-library/react'
export { default as userEvent } from '@testing-library/user-event'
// override render export
export { customRender as render }
