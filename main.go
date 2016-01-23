package main

import "fmt"

func main() {
	g := &game{
		turn:   "corporation",
		clicks: 3,
		corporation: corporation{
			identity: "jinteki: personal evolution",
			deck: []card{
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
					root: []serverCard{},
					ice:  []serverCard{},
				},
				{
					t:    "Archives",
					root: []serverCard{},
					ice:  []serverCard{},
				},
				{
					t:    "HQ",
					root: []serverCard{},
					ice:  []serverCard{},
				},
			},
		},
		runner: runner{
			identity: "Gabriel Santiago: Consummmate Professional",
			deck: []card{
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
	g.dispatch(corporationAdvanceCard{serverIndex: 3, location: "root", locationIndex: 0})
	g.dispatch(corporationAdvanceCard{serverIndex: 3, location: "root", locationIndex: 0})
	fmt.Println(g)
	g.dispatch(runnerTurn(4))
	fmt.Println(g)
}

type actioner interface {
	action(*game) *game
}

type gainCredits struct {
	corporation int
	runner      int
}

func (gc gainCredits) action(g *game) *game {
	g.corporation.credits += gc.corporation
	g.runner.credits += gc.runner
	return g
}

type game struct {
	corporation
	runner
	turn   string
	clicks int
}

type card string

func (g *game) dispatch(a actioner) {
	g = a.action(g)
}

type serverCard struct {
	rezzed            bool
	advancementTokens int
	virusCounters     int
	card
}

type server struct {
	t    string
	ice  []serverCard
	root []serverCard
}

type corporation struct {
	identity string
	credits  int
	deck     []card
	hand     []card
	servers  []server
}

type runner struct {
	identity  string
	credits   int
	deck      []card
	hand      []card
	programs  []card
	hardware  []card
	resources []card
}
