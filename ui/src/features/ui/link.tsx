import { Anchor } from '@mantine/core'
import { ReactNode } from 'react'
import { Link as RouterLink } from 'react-router-dom'

interface LinkProps {
	children: ReactNode
	to: string
}
export function Link({ children, to }: LinkProps) {
	return (
		<Anchor component={RouterLink} to={to}>
			{children}
		</Anchor>
	)
}
