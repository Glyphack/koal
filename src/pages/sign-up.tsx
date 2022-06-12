import { SignUpForm } from '../features/auth'
import { Link } from '../features/ui'

export function SignUpPage() {
	return (
		<main className="max-w-xs space-y-4">
			<div className="flex justify-end">
				<Link to="/sign-in">Sign In</Link>
			</div>
			<SignUpForm />
		</main>
	)
}
