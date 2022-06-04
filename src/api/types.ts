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
	projects: { id: string; name: string }[]
}

export interface CreateProjectRequest {
	name: string
}
