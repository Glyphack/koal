import { forwardRef, InputHTMLAttributes } from 'react'

interface FieldProps extends InputHTMLAttributes<HTMLInputElement> {
	label?: string
}

export const Field = forwardRef<HTMLInputElement, FieldProps>(({ label, ...rest }, ref) => {
	return (
		<div className="flex flex-col gap-1">
			{label && (
				<label className="text-sm font-semibold" htmlFor={rest.name}>
					{label}
				</label>
			)}
			<input
				{...rest}
				className="p-1 transition rounded-lg outline-none bg-gray-50 focus:ring-1 ring-offset-2 ring-emerald-300 focus:bg-emerald-50"
				id={rest.name}
				ref={ref}
			/>
		</div>
	)
})

Field.displayName = 'Field'
