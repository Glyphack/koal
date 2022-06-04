import { SignInForm } from '../features/auth'
import { Link } from '../features/ui'

export function SignInPage() {
	return (
		<main>
			<SignInForm />
			<Link to="/sign-up">Sign Up</Link>
		</main>
	)
}
