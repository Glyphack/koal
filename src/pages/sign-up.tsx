import { AxiosError } from 'axios'
import { useForm } from 'react-hook-form'
import { useMutation } from 'react-query'
import { api, ResponseError } from '../api'
import { Field } from '../features/ui'

interface SignUpFormValues {
	email: string
	password: string
}

export function SignUpPage() {
	const form = useForm<SignUpFormValues>({ defaultValues: { email: '', password: '' } })
	const signUpMutation = useMutation(api.signUp)
	const signUp = form.handleSubmit((values: SignUpFormValues) => signUpMutation.mutate(values))

	return (
		<main>
			<form onSubmit={signUp}>
				<Field type="text" placeholder="Email" {...form.register('email')} />
				<Field type="password" placeholder="Password" {...form.register('password')} />
				<button type="submit">Sign Up</button>
				{signUpMutation.isLoading && <p>Submitting...</p>}
			</form>
			{signUpMutation.error instanceof AxiosError && (
				<p>{(signUpMutation.error.response?.data as ResponseError).message}</p>
			)}
		</main>
	)
}
