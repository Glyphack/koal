import { useForm } from 'react-hook-form'
import { useMutation } from 'react-query'
import { api } from '../api'
import { useAuthorize } from '../features/auth'
import { Button, ErrorMessage, Field, Link, Loader } from '../features/ui'

export function SignInPage() {
	const authorize = useAuthorize()
	const form = useForm({ defaultValues: { email: '', password: '' } })
	const signInMutation = useMutation(api.signIn, {
		onSuccess: (data) => {
			const token = data.data.token
			authorize(token)
		},
	})
	const signIn = form.handleSubmit((values) => signInMutation.mutate(values))

	return (
		<main>
			<form onSubmit={signIn}>
				<Field type="text" placeholder="Email" {...form.register('email')} />
				<Field type="password" placeholder="Password" {...form.register('password')} />
				<Button type="submit" disabled={signInMutation.isLoading}>
					Sign In
				</Button>
				<Loader loading={signInMutation.isLoading} />
			</form>
			<ErrorMessage error={signInMutation.error} />
			<Link to="/sign-up">Sign Up</Link>
		</main>
	)
}
