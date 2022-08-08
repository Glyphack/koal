import { render, screen } from '../../../test/test-utils'
import { Linkified } from '../linkified'

describe('Linkified', () => {
	test('should render linkified text', () => {
		render(<Linkified>Hello www.world.com again. visit https://mock.net</Linkified>)
		const link1 = screen.getByRole('link', { name: /www.world.com/i })
		const link2 = screen.getByRole('link', { name: /https:\/\/mock.net/i })
		expect(link1).toBeInTheDocument()
		expect(link2).toBeInTheDocument()
	})
})
