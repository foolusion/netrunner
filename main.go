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
			servers: []server{
				{
					t:    "R&D",
					root: []string{},
					ice:  []string{},
				},
				{
					t:    "Archives",
					root: []string{},
					ice:  []string{},
				},
				{
					t:    "HQ",
					root: []string{},
					ice:  []string{},
				},
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
	g.dispatch(gainCredits{corporation: 5, runner: 5})
	g.dispatch(corporationInstallNewRemote{card: "Nisei MK II"})
	fmt.Println(g)
	g.dispatch(runnerTurn(4))
	fmt.Println(g)
}

type actioner interface {
	action(*gameState) *gameState
}

type gainCredits struct {
	corporation int
	runner      int
}

func (gc gainCredits) action(g *gameState) *gameState {
	g.corporationState.credits += gc.corporation
	g.runnerState.credits += gc.runner
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
	credits  int
	deck     []string
	hand     []string
	servers  []server
}

type runnerState struct {
	identity  string
	credits   int
	deck      []string
	hand      []string
	programs  []string
	hardware  []string
	resources []string
}
