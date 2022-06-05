import { rest } from 'msw'
import { FetchProjectResponse, FetchProjectsResponse } from '../api'

const API_URL = import.meta.env.VITE_API_URL

export const mockProject = { id: 'p1', name: 'project 1' }
export const mockTask = { id: 't1', title: 'task 1', isDone: false, project: mockProject }

export const handlers = [
	rest.get(`${API_URL}/mock`, (req, res, ctx) => res(ctx.status(200))),
	rest.post(`${API_URL}/auth/register`, (req, res, ctx) => res(ctx.status(200))),
	rest.post(`${API_URL}/auth/login`, (req, res, ctx) => res(ctx.status(200))),
	rest.post(`${API_URL}/todo/projects`, (req, res, ctx) => res(ctx.status(200))),
	rest.get(`${API_URL}/todo/projects`, (req, res, ctx) =>
		res(ctx.status(200), ctx.json<FetchProjectsResponse>({ projects: [mockProject] }))
	),
	rest.post(`${API_URL}/todo/items`, (req, res, ctx) => res(ctx.status(200))),
	rest.delete(`${API_URL}/todo/items/:id`, (req, res, ctx) => res(ctx.status(200))),
	rest.patch(`${API_URL}/todo/items/:id`, (req, res, ctx) => res(ctx.status(200))),
	rest.get(`${API_URL}/todo/projects/:id`, (req, res, ctx) =>
		res(
			ctx.status(200),
			ctx.json<FetchProjectResponse>({ info: mockProject, items: [mockTask] })
		)
	),
]
