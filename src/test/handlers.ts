import { rest } from 'msw'

const API_URL = import.meta.env.VITE_API_URL

export const handlers = [
	rest.get(`${API_URL}/mock`, (req, res, ctx) => res(ctx.status(200))),
	rest.post(`${API_URL}/auth/register`, (req, res, ctx) => res(ctx.status(200))),
	rest.post(`${API_URL}/auth/login`, (req, res, ctx) => res(ctx.status(200))),
	rest.post(`${API_URL}/todo/projects`, (req, res, ctx) => res(ctx.status(200))),
	rest.get(`${API_URL}/todo/projects`, (req, res, ctx) =>
		res(ctx.status(200), ctx.json({ projects: [{ id: '1', name: 'Test Project' }] }))
	),
]
