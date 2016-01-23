package main

type runnerDraw struct {
	clickCost int
	numCards  int
}

func (r runnerDraw) action(g *gameState) *gameState {
	g.clicks -= r.clickCost
	for i := 0; i < r.numCards; i++ {
		var c string
		c, g.runnerState.deck = g.runnerState.deck[0], g.runnerState.deck[1:]
		g.runnerState.hand = append(g.runnerState.hand, c)
	}
	return g
}

type runnerClickCredit struct {
	click  int
	credit int
}

func (rcc runnerClickCredit) action(g *gameState) *gameState {
	g.clicks -= rcc.click
	g.dispatch(gainCredits{runner: rcc.credit})
	return g
}

type runnerTurn int

func (r runnerTurn) action(g *gameState) *gameState {
	g.turn = "runner"
	g.clicks = int(r)
	return g
}
