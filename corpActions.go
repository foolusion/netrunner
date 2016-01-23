package main

type corporationDraw struct {
	clickCost int
	numCards  int
}

func (c corporationDraw) action(g *gameState) *gameState {
	g.clicks -= c.clickCost
	for i := 0; i < c.numCards; i++ {
		var c card
		c, g.corporationState.deck = g.corporationState.deck[0], g.corporationState.deck[1:]
		g.corporationState.hand = append(g.corporationState.hand, c)
	}
	return g
}

type corporationClickCredit struct {
	click  int
	credit int
}

func (ccc corporationClickCredit) action(g *gameState) *gameState {
	g.clicks -= ccc.click
	g.dispatch(gainCredits{corporation: ccc.credit})
	return g
}

type corporationInstallNewRemote struct {
	card
}

func (c corporationInstallNewRemote) action(g *gameState) *gameState {
	for i, v := range g.corporationState.hand {
		if v == c.card {
			g.corporationState.hand[i], g.corporationState.hand = g.corporationState.hand[len(g.corporationState.hand)-1], g.corporationState.hand[:len(g.corporationState.hand)-1]
			break
		}
	}
	g.corporationState.servers = append(g.corporationState.servers, server{t: "remote", ice: []serverCard{}, root: []serverCard{{card: c.card}}})
	g.clicks--
	return g
}

type corporationAdvanceCard struct {
	serverIndex   int
	location      string
	locationIndex int
}

func (c corporationAdvanceCard) action(g *gameState) *gameState {
	switch c.location {
	case "root":
		g.corporationState.servers[c.serverIndex].root[c.locationIndex].advancementTokens++
	case "ice":
		g.corporationState.servers[c.serverIndex].ice[c.locationIndex].advancementTokens++
	}
	g.clicks--
	g.corporationState.credits--
	return g
}

type corporationTurn int

func (c corporationTurn) action(g *gameState) *gameState {
	g.turn = "corporation"
	g.clicks = int(c)
	return g
}
