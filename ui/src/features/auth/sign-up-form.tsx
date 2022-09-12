import { Button, TextInput } from '@mantine/core'
import { useForm } from '@mantine/form'
import { useMutation } from 'react-query'
import { useNavigate } from 'react-router-dom'
import { api } from '../../api'
import { ErrorMessage, Form } from '../ui'
import { useAuth } from './use-auth'

export function SignUpForm() {
	const navigate = useNavigate()
	const authenticate = useAuth((state) => state.authenticate)
	const form = useForm({ initialValues: { email: '', password: '' } })
	const signUpMutation = useMutation(api.signUp, {
		onSuccess: (data) => {
			authenticate(data.data.token)
			navigate('/project', { replace: true })
		},
	})
	const signUp = form.onSubmit((values) => signUpMutation.mutate(values))

	return (
		<Form
			onSubmit={signUp}
			fields={
				<>
					<TextInput
						label="Email"
						type="text"
						placeholder="email@address.com"
						{...form.getInputProps('email')}
					/>
					<TextInput
						label="Password"
						type="password"
						placeholder="a strong password"
						{...form.getInputProps('password')}
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
