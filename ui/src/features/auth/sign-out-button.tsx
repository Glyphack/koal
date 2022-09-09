import { Button } from '@mantine/core'
import { HiArrowSmLeft } from 'react-icons/hi'
import { useQueryClient } from 'react-query'
import { useAuth } from './use-auth'

export function SignOutButton() {
	const signOut = useAuth((state) => state.signOut)
	const queryClient = useQueryClient()
	const handleSignOut = () => {
		signOut()
		queryClient.clear()
	}

	return (
		<Button variant="subtle" leftIcon={<HiArrowSmLeft />} onClick={handleSignOut}>
			Sign Out
		</Button>
	)
}
