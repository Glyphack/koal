import { createEmotionCache, MantineProvider, MantineThemeOverride } from '@mantine/core'
import { QueryClient, QueryClientProvider } from 'react-query'
import { BrowserRouter } from 'react-router-dom'
import { Layout } from '../features/ui'
import { Router } from './router'

export const queryClient = new QueryClient({
	defaultOptions: {
		queries: { retry: false, refetchOnWindowFocus: false },
		mutations: { retry: false },
	},
})

const emotionCache = createEmotionCache({ key: 'mantine' })
const theme: MantineThemeOverride = {
	colors: {
		emerald: [
			'#ecfdf5',
			'#d1fae5',
			'#a7f3d0',
			'#6ee7b7',
			'#34d399',
			'#10b981',
			'#059669',
			'#047857',
			'#065f46',
			'#064e3b',
		],
	},
	primaryColor: 'emerald',
	fontFamily: "'Nunito', 'sans-serif'",
	defaultRadius: 'md',
}

export function App() {
	return (
		<BrowserRouter>
			<QueryClientProvider client={queryClient}>
				<MantineProvider emotionCache={emotionCache} theme={theme}>
					<Layout>
						<Router />
					</Layout>
				</MantineProvider>
			</QueryClientProvider>
		</BrowserRouter>
	)
}
