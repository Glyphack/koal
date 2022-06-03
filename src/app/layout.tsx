import { ReactNode } from 'react'

interface LayoutProps {
	children: ReactNode
}

export function Layout({ children }: LayoutProps) {
	return <div>{children}</div>
}
