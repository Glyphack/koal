import { FormHTMLAttributes, ReactNode } from 'react'

interface FormProps extends FormHTMLAttributes<HTMLFormElement> {
	fields: ReactNode
	actions: ReactNode
}

export function Form({ fields, actions, ...rest }: FormProps) {
	return (
		<form {...rest} className="space-y-16">
			<div className="flex flex-col gap-6">{fields}</div>
			<div className="flex flex-col gap-2">{actions}</div>
		</form>
	)
}
