import { useForm } from 'react-hook-form'
import { useMutation } from 'react-query'
import { useNavigate } from 'react-router-dom'
import { api } from '../../api'
import { Button, ErrorMessage, Field, Loader } from '../ui'
import { useAuth } from './use-auth'

export function SignInForm() {
	const navigate = useNavigate()
	const authenticate = useAuth((state) => state.authenticate)
	const form = useForm({ defaultValues: { email: '', password: '' } })
	const signInMutation = useMutation(api.signIn, {
		onSuccess: (data) => {
			authenticate(data.data.token)
			navigate('/project', { replace: true })
		},
	})
	const signIn = form.handleSubmit((values) => signInMutation.mutate(values))

	return (
		<form onSubmit={signIn}>
			<Field type="text" placeholder="Email" {...form.register('email')} />
			<Field type="password" placeholder="Password" {...form.register('password')} />
			<Button type="submit" disabled={signInMutation.isLoading}>
				Sign In
			</Button>
			<Loader loading={signInMutation.isLoading} />
			<ErrorMessage error={signInMutation.error} />
		</form>
	)
}
