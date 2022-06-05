import { ButtonHTMLAttributes } from 'react'

type FieldProps = ButtonHTMLAttributes<HTMLButtonElement>

export function Button(props: FieldProps) {
	return <button {...props} />
}
