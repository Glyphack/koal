import { useForm } from 'react-hook-form'
import { useMutation } from 'react-query'
import { Link } from 'react-router-dom'
import { api } from '../api'
import { useAuthorize } from '../features/auth'
import { ErrorMessage, Field, Loader } from '../features/ui'

export function SignUpPage() {
	const authorize = useAuthorize()
	const form = useForm({ defaultValues: { email: '', password: '' } })
	const signUpMutation = useMutation(api.signUp, {
		onSuccess: (data) => {
			const token = data.data.token
			authorize(token)
		},
	})
	const signUp = form.handleSubmit((values) => signUpMutation.mutate(values))

	return (
		<main>
			<form onSubmit={signUp}>
				<Field type="text" placeholder="Email" {...form.register('email')} />
				<Field type="password" placeholder="Password" {...form.register('password')} />
				<button type="submit" disabled={signUpMutation.isLoading}>
					Sign Up
				</button>
				<Loader loading={signUpMutation.isLoading} />
			</form>
			<ErrorMessage error={signUpMutation.error} />
			<Link to="/sign-in">Sign In</Link>
		</main>
	)
}
