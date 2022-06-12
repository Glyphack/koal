import { HiArrowSmRight } from 'react-icons/hi'
import { Link as RouterLink, LinkProps } from 'react-router-dom'

export function Link({ children, ...rest }: LinkProps) {
	return (
		<RouterLink
			{...rest}
			className="flex items-center gap-1 font-semibold transition rounded-lg outline-none focus:text-emerald-800 hover:text-emerald-800 ring-offset-2 text-emerald-900 focus:ring-1 ring-emerald-300 max-w-min whitespace-nowrap"
		>
			{children} <HiArrowSmRight />
		</RouterLink>
	)
}
