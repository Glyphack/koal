import clsx from 'clsx'
import { ButtonHTMLAttributes } from 'react'
import { Loader } from './loader'

interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
	loading?: boolean
	variant?: 'filled' | 'text' | 'icon'
}

export function Button({ children, variant = 'filled', disabled, loading, ...rest }: ButtonProps) {
	return (
		<button
			{...rest}
			className={clsx(
				'flex justify-center items-center transition rounded-lg hover:text-emerald-800 focus:text-emerald-800 outline-none text-emerald-900 focus:ring-1 ring-emerald-300 ring-offset-2',
				variant === 'filled' &&
					'hover:bg-emerald-300 w-full font-black py-3 px-4 bg-emerald-400 disabled:bg-emerald-400 shadow-emerald-100 shadow-sm focus:bg-emerald-300',
				variant === 'text' && 'font-semibold gap-1',
				variant === 'icon' &&
					'w-10 h-10 text-lg rounded-full hover:bg-emerald-50 focus:bg-emerald-50'
			)}
			disabled={disabled || loading}
		>
			{loading ? <Loader size={variant === 'icon' ? 18 : 36} /> : children}
		</button>
	)
}
