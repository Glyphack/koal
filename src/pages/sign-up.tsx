import { SignUpForm } from '../features/auth'
import { Link } from '../features/ui'

export function SignUpPage() {
	return (
		<main>
			<SignUpForm />
			<Link to="/sign-in">Sign In</Link>
		</main>
	)
}
