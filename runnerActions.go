package main

type runnerDraw struct {
	clickCost int
	numCards  int
}

func (r runnerDraw) action(g *game) *game {
	g.clicks -= r.clickCost
	for i := 0; i < r.numCards; i++ {
		var c string
		c, g.runner.deck = g.runner.deck[0], g.runner.deck[1:]
		g.runner.hand = append(g.runner.hand, c)
	}
	return g
}

type runnerClickCredit struct {
	click  int
	credit int
}

func (rcc runnerClickCredit) action(g *game) *game {
	g.clicks -= rcc.click
	g.dispatch(gainCredits{runner: rcc.credit})
	return g
}

type runnerTurn int

func (r runnerTurn) action(g *game) *game {
	g.turn = "runner"
	g.clicks = int(r)
	return g
}
