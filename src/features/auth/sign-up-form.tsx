import { useForm } from 'react-hook-form'
import { useMutation } from 'react-query'
import { useNavigate } from 'react-router-dom'
import { api } from '../../api'
import { Button, ErrorMessage, Field, Loader } from '../ui'
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
		<form onSubmit={signUp}>
			<Field type="text" placeholder="Email" {...form.register('email')} />
			<Field type="password" placeholder="Password" {...form.register('password')} />
			<Button type="submit" disabled={signUpMutation.isLoading}>
				Sign Up
			</Button>
			<Loader loading={signUpMutation.isLoading} />
			<ErrorMessage error={signUpMutation.error} />
		</form>
	)
}
