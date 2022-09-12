import urlRegex from 'url-regex'

interface LinkifiedProps {
	children: string
}

export function Linkified({ children }: LinkifiedProps) {
	return (
		<>
			{children.split(/( |\n)/g).map((word, index) =>
				word.match(urlRegex()) ? (
					<a
						key={index}
						className="underline transition hover:text-gray-600 underline-offset-1"
						href={word}
						target="_blank"
						rel="noreferrer"
					>
						{word}
					</a>
				) : (
					<span key={index}> {word} </span>
				)
			)}
		</>
	)
}
