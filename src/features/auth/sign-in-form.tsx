import { useForm } from 'react-hook-form'
import { useMutation } from 'react-query'
import { useNavigate } from 'react-router-dom'
import { api } from '../../api'
import { Button, ErrorMessage, Field, Form } from '../ui'
import { useAuth } from './use-auth'

export function SignInForm() {
	const navigate = useNavigate()
	const { authenticate, triedToVisitPage } = useAuth((state) => ({
		authenticate: state.authenticate,
		triedToVisitPage: state.triedToVisitPage,
	}))
	const form = useForm({ defaultValues: { email: '', password: '' } })
	const signInMutation = useMutation(api.signIn, {
		onSuccess: (data) => {
			authenticate(data.data.token)
			navigate(triedToVisitPage || '/project', { replace: true })
		},
	})
	const signIn = form.handleSubmit((values) => signInMutation.mutate(values))

	return (
		<Form
			onSubmit={signIn}
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
						placeholder="your password"
						{...form.register('password')}
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
