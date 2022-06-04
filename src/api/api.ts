import { privateRequest, publicRequest } from './config'
import {
	CreateProjectRequest,
	FetchProjectsResponse,
	SignInRequest,
	SignInResponse,
	SignUpRequest,
	SignUpResponse,
} from './types'

const signUp = (payload: SignUpRequest) =>
	publicRequest.post<SignUpResponse>('/auth/register', payload)

const signIn = (payload: SignInRequest) =>
	publicRequest.post<SignInResponse>('/auth/login', payload)

const fetchProjects = () => privateRequest.get<FetchProjectsResponse>('/todo/projects')

const createProject = (payload: CreateProjectRequest) =>
	privateRequest.post('/todo/projects', payload)

export const api = {
	signUp,
	signIn,
	fetchProjects,
	createProject,
}

export enum QueryKey {
	Projects = 'projects',
}
