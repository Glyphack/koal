export interface ResponseError {
	code: number
	message: string
}

export interface SignUpRequest {
	email: string
	password: string
}

export interface SignUpResponse {
	token: string
}

export interface SignInRequest {
	email: string
	password: string
}

export interface SignInResponse {
	token: string
}

export interface FetchProjectsResponse {
	projects: ProjectInfo[]
}

export interface CreateProjectRequest {
	name: string
}

export interface FetchProjectResponse {
	info: ProjectInfo
	items: Task[]
}

export interface CreateTaskRequest {
	projectId: string
	title: string
}

export interface UpdateTaskRequest {
	title?: string
	isDone?: boolean
}

export interface FetchInboxResponse {
	items: Task[]
}

// ####################################################################################################

export interface ProjectInfo {
	id: string
	name: string
}

export interface Task {
	id: string
	title: string
	isDone: boolean
	project: ProjectInfo
}
