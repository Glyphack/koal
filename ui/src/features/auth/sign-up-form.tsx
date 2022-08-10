import { useForm } from 'react-hook-form'
import { useMutation } from 'react-query'
import { useNavigate } from 'react-router-dom'
import { api } from '../../api'
import { Button, ErrorMessage, Field, Form } from '../ui'
import { useAuth } from './use-auth'

export function SignUpForm() {
	const navigate = useNavigate()
	const authenticate = useAuth((state) => state.authenticate)
	const form = useForm({ defaultValues: { email: '', password: '' } })
	const signUpMutation = useMutation(api.signUp, {
		onSuccess: (data) => {
			authenticate(data.data.token)
			navigate('/project', { replace: true })
		},
	})
	const signUp = form.handleSubmit((values) => signUpMutation.mutate(values))

	return (
		<Form
			onSubmit={signUp}
			fields={
				<>
					<Field
						label="Email"
						type="text"
						placeholder="email@address.com"
						{...form.register('email')}
					/>
					<Field
						label="Password"
						type="password"
						placeholder="a strong password"
						{...form.register('password')}
					/>
				</>
			}
			actions={
				<>
					<Button type="submit" loading={signUpMutation.isLoading}>
						Sign Up
					</Button>
					<ErrorMessage error={signUpMutation.error} />
				</>
			}
		/>
	)
}
