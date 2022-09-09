import { ReactNode } from 'react'

interface LayoutProps {
	children: ReactNode
}

export function Layout({ children }: LayoutProps) {
	return (
		<div className="container min-h-screen p-4 mx-auto text-gray-700 selection:bg-emerald-500 selection:text-emerald-900 font-body lg:px-20 lg:py-16">
			<div>{children}</div>
		</div>
	)
}
