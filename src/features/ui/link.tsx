import clsx from 'clsx'
import { HiArrowSmLeft, HiArrowSmRight } from 'react-icons/hi'
import { Link as RouterLink, LinkProps as RouterLinkProps } from 'react-router-dom'

interface LinkProps extends RouterLinkProps {
	withLeftIcon?: boolean
}

export function Link({ children, className, withLeftIcon, ...rest }: LinkProps) {
	return (
		<RouterLink
			{...rest}
			className={clsx(
				'flex items-center gap-1 font-semibold transition rounded-lg outline-none focus:text-emerald-800 hover:text-emerald-800 ring-offset-2 text-emerald-900 focus:ring-1 ring-emerald-300 max-w-min whitespace-nowrap',
				className
			)}
		>
			{withLeftIcon && <HiArrowSmLeft />} {children} {!withLeftIcon && <HiArrowSmRight />}
		</RouterLink>
	)
}
