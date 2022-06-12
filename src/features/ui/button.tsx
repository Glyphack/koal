import { ButtonHTMLAttributes } from 'react'
import { Loader } from './loader'

interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
	loading?: boolean
}

export function Button({ children, disabled, loading, ...rest }: ButtonProps) {
	return (
		<button
			{...rest}
			className="w-full px-4 py-3 font-black transition rounded-lg shadow-sm outline-none hover:bg-emerald-300 bg-emerald-400 text-emerald-900 focus:ring-1 ring-emerald-300 disabled:bg-emerald-400 shadow-emerald-100 ring-offset-2 focus:bg-emerald-300"
			disabled={disabled || loading}
		>
			{loading ? <Loader /> : children}
		</button>
	)
}
