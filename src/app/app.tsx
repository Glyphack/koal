import { QueryClient, QueryClientProvider } from 'react-query'
import { BrowserRouter } from 'react-router-dom'
import { Layout } from '../features/ui'
import { Router } from './router'

const client = new QueryClient({
	defaultOptions: {
		queries: { retry: false, refetchOnWindowFocus: false },
		mutations: { retry: false },
	},
})

export function App() {
	return (
		<BrowserRouter>
			<QueryClientProvider client={client}>
				<Layout>
					<Router />
				</Layout>
			</QueryClientProvider>
		</BrowserRouter>
	)
}
