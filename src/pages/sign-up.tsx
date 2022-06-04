import { useForm } from 'react-hook-form'
import { useMutation } from 'react-query'
import { useNavigate } from 'react-router-dom'
import { api } from '../api'
import { useAuth } from '../features/auth'
import { Button, ErrorMessage, Field, Link, Loader } from '../features/ui'

export function SignUpPage() {
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
		<main>
			<form onSubmit={signUp}>
				<Field type="text" placeholder="Email" {...form.register('email')} />
				<Field type="password" placeholder="Password" {...form.register('password')} />
				<Button type="submit" disabled={signUpMutation.isLoading}>
					Sign Up
				</Button>
				<Loader loading={signUpMutation.isLoading} />
			</form>
			<ErrorMessage error={signUpMutation.error} />
			<Link to="/sign-in">Sign In</Link>
		</main>
	)
}
