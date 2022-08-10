import { useQuery } from 'react-query'
import { api, QueryKey } from '../api'
import { InboxTaskList } from '../features/task'
import { Loader } from '../features/ui'

export function InboxPage() {
	const inboxQuery = useQuery(QueryKey.Inbox, api.fetchInbox)
	const inbox = inboxQuery.data?.data

	if (inboxQuery.isLoading || !inbox) return <Loader />

	return (
		<main className="space-y-20">
			<h3 className="text-3xl">Inbox</h3>
			<InboxTaskList tasks={inbox.items} />
		</main>
	)
}
