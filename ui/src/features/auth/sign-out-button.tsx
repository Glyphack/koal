import { HiArrowSmLeft } from 'react-icons/hi'
import { useQueryClient } from 'react-query'
import { Button } from '../ui'
import { useAuth } from './use-auth'

export function SignOutButton() {
	const signOut = useAuth((state) => state.signOut)
	const queryClient = useQueryClient()
	const handleSignOut = () => {
		signOut()
		queryClient.clear()
	}

	return (
		<Button variant="text" onClick={handleSignOut}>
			<HiArrowSmLeft /> Sign Out
		</Button>
	)
}
