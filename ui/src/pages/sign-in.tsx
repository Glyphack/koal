import { SignInForm } from '../features/auth'
import { Link } from '../features/ui'

export function SignInPage() {
	return (
		<main className="space-y-4 lg:max-w-xs">
			<div className="flex justify-end">
				<Link to="/sign-up">Sign Up</Link>
			</div>
			<SignInForm />
		</main>
	)
}
