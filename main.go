package main

import "fmt"

func main() {
	g := &gameState{
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
}

type gameState struct {
	corporationState
	runnerState
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
