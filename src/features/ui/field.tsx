import { forwardRef, InputHTMLAttributes } from 'react'

type FieldProps = InputHTMLAttributes<HTMLInputElement>

export const Field = forwardRef<HTMLInputElement, FieldProps>((props, ref) => {
	return <input className="border" {...props} ref={ref} />
})

Field.displayName = 'Field'
