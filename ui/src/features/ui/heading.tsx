import clsx from 'clsx'
import { ReactNode } from 'react'

type HeadingElement = 'h1' | 'h2' | 'h3' | 'h4' | 'h5' | 'h6'

interface HeadingProps {
	children: ReactNode
	as?: HeadingElement
	className?: string
}

export function Heading({ children, as = 'h1', className }: HeadingProps) {
	const Element = as
	return <Element className={clsx('font-thin', className)}>{children}</Element>
}
