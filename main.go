package main

import "fmt"

func main() {
	g := &gameState{
		turn:   "corporation",
		clicks: 3,
		corporationState: corporationState{
			identity: "jinteki: personal evolution",
			deck: []string{
				"Nisei MK II",
				"Project Junebug",
				"Snare!",
				"Zaibatsu Loyalty",
				"Neural EMP",
				"Precognition",
				"Cell Portal",
				"Chum",
				"Data Mine",
				"Neural Katana",
				"Wall of Thorns",
				"Akiro Watanabe",
				"Priority Requisition",
				"Private Security Force",
				"Melange Mining Corp.",
				"PAD Campaign",
				"Hedge Fund",
				"Enigma",
				"Hunter",
				"Wall of Static",
			},
		},
		runnerState: runnerState{
			identity: "Gabriel Santiago: Consummmate Professional",
			deck: []string{
				"Account Siphon",
				"Easy Mark",
				"Forged Activation Orders",
				"Inside Job",
				"Special Order",
				"Lemuria Codecracker",
				"Desperado",
				"Aurora",
				"Femme Fatale",
				"Ninja",
				"Sneakdoor Beta",
				"Bank Job",
				"Crash Space",
				"Data Dealer",
				"Decoy",
				"Infiltration",
				"Sure Gamble",
				"Crypsis",
				"Access to Globalsec",
				"Armitage Codebusting",
			},
		},
	}
	g.dispatch(corporationDraw{clickCost: 0, numCards: 5})
	g.dispatch(runnerDraw{clickCost: 0, numCards: 5})
	fmt.Println(g)
	g.dispatch(runnerTurn(4))
	fmt.Println(g)
}

type actioner interface {
	action(*gameState) *gameState
}

type corporationDraw struct {
	clickCost int
	numCards  int
}

func (c corporationDraw) action(g *gameState) *gameState {
	g.clicks -= c.clickCost
	for i := 0; i < c.numCards; i++ {
		var c string
		c, g.corporationState.deck = g.corporationState.deck[0], g.corporationState.deck[1:]
		g.corporationState.hand = append(g.corporationState.hand, c)
	}
	return g
}

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

type corporationTurn int

func (c corporationTurn) action(g *gameState) *gameState {
	g.turn = "corporation"
	g.clicks = int(c)
	return g
}

type runnerTurn int

func (r runnerTurn) action(g *gameState) *gameState {
	g.turn = "runner"
	g.clicks = int(r)
	return g
}

type gameState struct {
	corporationState
	runnerState
	turn   string
	clicks int
}

func (g *gameState) dispatch(a actioner) {
	g = a.action(g)
}

type server struct {
	t    string
	ice  []string
	root []string
}

type corporationState struct {
	identity string
	deck     []string
	hand     []string
	servers  []server
}

type runnerState struct {
	identity  string
	deck      []string
	hand      []string
	programs  []string
	hardware  []string
	resources []string
}
