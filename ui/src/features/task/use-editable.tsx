import { useEffect, useRef, useState } from 'react'

type HtmlEditableElement = HTMLInputElement | HTMLTextAreaElement

export function useEditable<T extends HtmlEditableElement>() {
	const [isEditing, setIsEditing] = useState(false)
	const inputRef = useRef<T>(null)

	useEffect(() => {
		if (isEditing) {
			inputRef.current?.focus()
			inputRef.current?.select()
		}
	}, [isEditing])

	return {
		isEditing,
		setIsEditing,
		inputRef,
	}
}
