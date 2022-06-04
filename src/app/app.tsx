import { QueryClient, QueryClientProvider } from 'react-query'
import { BrowserRouter } from 'react-router-dom'
import { Layout } from './layout'
import { Router } from './router'

const client = new QueryClient()

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
