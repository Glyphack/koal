import { privateRequest, publicRequest } from './config'
import {
	CreateProjectRequest,
	CreateTaskRequest,
	FetchProjectResponse,
	FetchProjectsResponse,
	SignInRequest,
	SignInResponse,
	SignUpRequest,
	SignUpResponse,
	UpdateTaskRequest,
} from './types'

const signUp = (payload: SignUpRequest) =>
	publicRequest.post<SignUpResponse>('/auth/register', payload)

const signIn = (payload: SignInRequest) =>
	publicRequest.post<SignInResponse>('/auth/login', payload)

const fetchProjects = () => privateRequest.get<FetchProjectsResponse>('/todo/projects')

const fetchProject = (id: string) =>
	privateRequest.get<FetchProjectResponse>(`/todo/projects/${id}`)

const createProject = (payload: CreateProjectRequest) =>
	privateRequest.post('/todo/projects', payload)

const deleteProject = (id: string) => privateRequest.delete(`/todo/projects/${id}`)

const createTask = (payload: CreateTaskRequest) => privateRequest.post('/todo/items', payload)

const updateTask = ({ id, payload }: { id: string; payload: UpdateTaskRequest }) =>
	privateRequest.patch(`/todo/items/${id}`, payload)

const deleteTask = (id: string) => privateRequest.delete(`/todo/items/${id}`)

export const api = {
	signUp,
	signIn,
	fetchProjects,
	fetchProject,
	createProject,
	deleteProject,
	createTask,
	updateTask,
	deleteTask,
}

export enum QueryKey {
	Projects = 'projects',
	Project = 'project',
}
