import { Button, TextInput } from '@mantine/core'
import { useForm } from '@mantine/form'
import { useMutation } from 'react-query'
import { useNavigate } from 'react-router-dom'
import { api } from '../../api'
import { ErrorMessage, Form } from '../ui'
import { useAuth } from './use-auth'

export function SignInForm() {
	const navigate = useNavigate()
	const { authenticate, triedToVisitPage } = useAuth((state) => ({
		authenticate: state.authenticate,
		triedToVisitPage: state.triedToVisitPage,
	}))
	const form = useForm({ initialValues: { email: '', password: '' } })
	const signInMutation = useMutation(api.signIn, {
		onSuccess: (data) => {
			authenticate(data.data.token)
			navigate(triedToVisitPage || '/project', { replace: true })
		},
	})
	const signIn = form.onSubmit((values) => signInMutation.mutate(values))

	return (
		<Form
			onSubmit={signIn}
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
						placeholder="your password"
						{...form.getInputProps('password')}
					/>
				</>
			}
			actions={
				<>
					<Button type="submit" loading={signInMutation.isLoading}>
						Sign In
					</Button>
					<ErrorMessage error={signInMutation.error} />
				</>
			}
		/>
	)
}
