package main

type corporationDraw struct {
	clickCost int
	numCards  int
}

func (c corporationDraw) action(g *game) *game {
	g.corporation.clicks -= c.clickCost
	for i := 0; i < c.numCards; i++ {
		var c card
		c, g.corporation.deck = g.corporation.deck[0], g.corporation.deck[1:]
		g.corporation.hand = append(g.corporation.hand, c)
	}
	return g
}

type corporationClickCredit struct {
	click  int
	credit int
}

func (ccc corporationClickCredit) action(g *game) *game {
	g.corporation.clicks -= ccc.click
	g.dispatch(gainCredits{corporation: ccc.credit})
	return g
}

type corporationInstallNewRemote struct {
	handIndex int
}

func (c corporationInstallNewRemote) action(g *game) *game {
	card := g.corporation.hand[c.handIndex]
	g.corporation.hand[c.handIndex], g.corporation.hand = g.corporation.hand[len(g.corporation.hand)-1], g.corporation.hand[:len(g.corporation.hand)-1]
	g.corporation.servers = append(g.corporation.servers, server{t: "remote", ice: []serverCard{}, root: []serverCard{{card: card}}})
	g.corporation.clicks--
	return g
}

type corporationAdvanceCard struct {
	serverIndex   int
	location      string
	locationIndex int
}

func (c corporationAdvanceCard) action(g *game) *game {
	switch c.location {
	case "root":
		g.corporation.servers[c.serverIndex].root[c.locationIndex].advancementTokens++
	case "ice":
		g.corporation.servers[c.serverIndex].ice[c.locationIndex].advancementTokens++
	}
	g.corporation.clicks--
	g.corporation.credits--
	return g
}

type corporationTurn int

func (c corporationTurn) action(g *game) *game {
	g.turnPhase = "corporation-1.1"
	g.corporation.clicks = int(c)
	return g
}
