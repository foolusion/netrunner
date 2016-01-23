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
	fmt.Println(g)
	g.Dispatch(runnerTurn(4))
	fmt.Println(g)
}

type actioner interface {
	action(*gameState) *gameState
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

func (g *gameState) Dispatch(a actioner) {
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
