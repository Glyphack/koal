import { rest } from 'msw'
import { FetchProjectResponse, FetchProjectsResponse } from '../api'

const API_URL = import.meta.env.VITE_API_URL

export const mockProject1 = { id: 'p1', name: 'project 1' }
export const mockProject2 = { id: 'p2', name: 'project 2' }
export const mockTask = {
	id: 't1',
	title: 'task https://task1.com 1',
	isDone: false,
	project: mockProject1,
}
export const mockDoneTask = {
	id: 't2',
	title: 'task 2',
	isDone: true,
	project: mockProject2,
}
export const mockTasks = [mockTask, mockDoneTask]

export const handlers = [
	rest.get(`${API_URL}/mock`, (req, res, ctx) => res(ctx.status(200))),
	rest.post(`${API_URL}/auth/register`, (req, res, ctx) => res(ctx.status(200))),
	rest.post(`${API_URL}/auth/login`, (req, res, ctx) => res(ctx.status(200))),
	rest.post(`${API_URL}/todo/projects`, (req, res, ctx) => res(ctx.status(200))),
	rest.get(`${API_URL}/todo/projects`, (req, res, ctx) =>
		res(ctx.status(200), ctx.json<FetchProjectsResponse>({ projects: [mockProject1] }))
	),
	rest.delete(`${API_URL}/todo/projects/:id`, (req, res, ctx) => res(ctx.status(200))),
	rest.post(`${API_URL}/todo/items`, (req, res, ctx) => res(ctx.status(200))),
	rest.delete(`${API_URL}/todo/items/:id`, (req, res, ctx) => res(ctx.status(200))),
	rest.patch(`${API_URL}/todo/items/:id`, (req, res, ctx) => res(ctx.status(200))),
	rest.get(`${API_URL}/todo/projects/:id`, (req, res, ctx) =>
		res(
			ctx.status(200),
			ctx.json<FetchProjectResponse>({ info: mockProject1, items: mockTasks })
		)
	),
	rest.get(`${API_URL}/todo/lists/undone`, (req, res, ctx) =>
		res(
			ctx.status(200),
			ctx.json<FetchProjectResponse>({ info: mockProject1, items: [mockTask] })
		)
	),
]
