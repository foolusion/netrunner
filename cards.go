package main

type player int
type cardType int
type subtype int
type faction int

const (
	playerCorporation player = iota
	playerRunner

	cardTypeIdentity cardType = iota
	cardTypeAgenda
	cardTypeOperation
	cardTypeIce
	cardTypeUpgrade
	cardTypeAsset

	subtypeAP subtype = iota
	subtypeAmbush
	subtypeBarrier
	subtypeCodeGate
	subtypeDeflector
	subtypeGrayOps
	subtypeInitiative
	subtypeMegacorp
	subtypeObserver
	subtypeResearch
	subtypeSecurity
	subtypeSentry
	subtypeSysop
	subtypeTracer
	subtypeTrap
	subtypeTransaction
	subtypeUnorthodox

	factionJinteki = iota
	factionCriminal
	factionNeutral
)

var (
	cardJintekiPersonalEvolution = card{
		player:          playerCorporation,
		title:           "Jinteki",
		subtitle:        "Personal Evolution",
		cardType:        cardTypeIdentity,
		subtypes:        []subtype{subtypeMegacorp},
		text:            "Whenever an agenda is scored or stolen, do 1 net damage.",
		faction:         factionJinteki,
		minimumDeckSize: 45,
		influenceLimit:  15,
	}
	cardNiseiMKII = card{
		player:                 playerCorporation,
		title:                  "Nisei MK II",
		cardType:               cardTypeAgenda,
		subtypes:               []subtype{subtypeInitiative},
		text:                   "Place 1 agenda counter on Nisei MK II when you score it.\n\nHosted agenda counter: End the run.",
		faction:                factionJinteki,
		advancementRequirement: 4,
		agendaPoints:           2,
	}
	cardProjectJunebug = card{
		player:         playerCorporation,
		title:          "Project Junebug",
		cardType:       cardTypeAsset,
		subtypes:       []subtype{subtypeAmbush, subtypeResearch},
		text:           "Project Junebug can be advanced.\n\nIf you pay 1 credit when the Runner access Project Junebug, do 2 net damage for each advancement token on Project Junebug.",
		faction:        factionJinteki,
		influenceValue: 1,
		rezCost:        0,
		trashCost:      0,
	}
	cardSnare = card{
		player:         playerCorporation,
		title:          "Snare!",
		cardType:       cardTypeAsset,
		subtypes:       []subtype{subtypeAmbush},
		text:           "If Snare! is accessed from R&D, the Runner must reveal it.\n\nIf you pay 4 credits when the Runner accesses Snare!, do 3 net damage and give the runner 1 tag, Ignore this effect if the Runner accesses Snare! from Archives.",
		faction:        factionJinteki,
		influenceValue: 2,
		rezCost:        0,
		trashCost:      0,
	}
	cardZaibatsuLoyalty = card{
		player:         playerCorporation,
		title:          "Zaibatsu Loyalty",
		cardType:       cardTypeAsset,
		text:           "If the Runner is about to expose a card, you may rez Zaibatsu Loyalty.\n\n1 credit or trash: Prevent 1 card from being exposed.",
		faction:        factionJinteki,
		influenceValue: 1,
		rezCost:        0,
		trashCost:      4,
	}
	cardNeuralEMP = card{
		player:         playerCorporation,
		title:          "Neural EMP",
		cardType:       cardTypeOperation,
		subtypes:       []subtype{subtypeGrayOps},
		text:           "Play only if the Runnner made a run during his or her last turn.\n\nDo 1 net damage.",
		faction:        factionJinteki,
		influenceValue: 2,
		playCost:       2,
	}
	cardPrecognition = card{
		player:         playerCorporation,
		title:          "Precognition",
		cardType:       cardTypeOperation,
		text:           "Look at the top 5 cards of R&D and arrange them in any order.",
		faction:        factionJinteki,
		influenceValue: 3,
		playCost:       0,
	}
	cardCellPortal = card{
		player:         playerCorporation,
		title:          "Cell Portal",
		cardType:       cardTypeIce,
		subtypes:       []subtype{subtypeCodeGate, subtypeDeflector},
		text:           "subroutine: The Runner approaches the outermost piece of ice protecting the attacked server. Derez Cell Portal.",
		faction:        factionJinteki,
		influenceValue: 2,
		rezCost:        5,
		strength:       7,
	}
	cardChum = card{
		player:         playerCorporation,
		title:          "Chum",
		cardType:       cardTypeIce,
		subtypes:       []subtype{subtypeCodeGate},
		text:           "subroutine: The next piece of ice the Runner encounters during this run has +2 strength. Do 3 net damage unless the Runner breaks all subroutines on that piece of ice.",
		faction:        factionJinteki,
		influenceValue: 1,
		rezCost:        1,
		strength:       4,
	}
	cardDataMine = card{
		player:         playerCorporation,
		title:          "Data Mine",
		cardType:       cardTypeIce,
		subtypes:       []subtype{subtypeTrap, subtypeAP},
		text:           "subroutine: Do 1 net damage. Trash Data Mine.",
		faction:        factionJinteki,
		influenceValue: 2,
		rezCost:        0,
		strength:       2,
	}
	cardNeuralKatana = card{
		player:         playerCorporation,
		title:          "Neural Katana",
		cardType:       cardTypeIce,
		subtypes:       []subtype{subtypeSentry, subtypeAP},
		text:           "subroutine: Do 3 net damage.",
		faction:        factionJinteki,
		influenceValue: 2,
		rezCost:        4,
		strength:       3,
	}
	cardWallOfThorns = card{
		player:         playerCorporation,
		title:          "Wall of Thorns",
		cardType:       cardTypeIce,
		subtypes:       []subtype{subtypeBarrier, subtypeAP},
		text:           "subroutine: Do 2 net damage.\nsubroutine: End the run.",
		faction:        factionJinteki,
		influenceValue: 1,
		rezCost:        8,
		strength:       5,
	}
	cardAkitaroWatanabe = card{
		player:         playerCorporation,
		title:          "Akitaro Watanabe",
		cardType:       cardTypeUpgrade,
		subtypes:       []subtype{subtypeSysop, subtypeUnorthodox},
		text:           "The rez cost of ice protecting this server is lowered by 2.",
		faction:        factionJinteki,
		influenceValue: 2,
		rezCost:        1,
		trashCost:      3,
	}
	cardPriorityRequisition = card{
		player:                 playerCorporation,
		title:                  "Priority Requisition",
		cardType:               cardTypeAgenda,
		subtypes:               []subtype{subtypeSecurity},
		text:                   "When you score Priority Requisition, you may rez a piece of ice ignoring all costs.",
		faction:                factionNeutral,
		advancementRequirement: 5,
		agendaPoints:           3,
	}
	cardPrivateSecurityForce = card{
		player:                 playerCorporation,
		title:                  "Private Security Force",
		cardType:               cardTypeAgenda,
		subtypes:               []subtype{subtypeSecurity},
		text:                   "If the Runner is tagged, Private Security Force gains: \"Click: Do 1 meat damage.\"",
		faction:                factionNeutral,
		advancementRequirement: 4,
		agendaPoints:           2,
	}
	cardMelangeMiningCorp = card{
		player:    playerCorporation,
		title:     "Melange Mining Corp.",
		cardType:  cardTypeAsset,
		text:      "Click, Click, Click: Gain 7 Credits.",
		faction:   factionNeutral,
		rezCost:   1,
		trashCost: 1,
	}
	cardPADCampaign = card{
		player:    playerCorporation,
		title:     "PAD Campaign",
		cardType:  cardTypeAsset,
		text:      "Gain 1 Credit when your turn begins.",
		faction:   factionNeutral,
		rezCost:   2,
		trashCost: 4,
	}
	cardHedgeFund = card{
		player:   playerCorporation,
		title:    "Hegde Fund",
		cardType: cardTypeOperation,
		subtypes: []subtype{subtypeTransaction},
		text:     "Gain 9 CREDIT.",
		faction:  factionNeutral,
		playCost: 5,
	}
	cardEnigma = card{
		player:   playerCorporation,
		title:    "Enigma",
		cardType: cardTypeIce,
		subtypes: []subtype{subtypeCodeGate},
		text:     "SUBROUTINE: The Runner loses CLICK, if able.\nSUBROUTINE: End the run.",
		faction:  factionNeutral,
		rezCost:  3,
		strength: 2,
	}
	cardHunter = card{
		player:   playerCorporation,
		title:    "Hunter",
		cardType: cardTypeIce,
		subtypes: []subtype{subtypeSentry, subtypeTracer, subtypeObserver},
		text:     "SUBROUTINE: TRACE(3) - If successful, give the Runner 1 tag.",
		faction:  factionNeutral,
		rezCost:  1,
		strength: 4,
	}
	cardWallOfStatic = card{
		player:   playerCorporation,
		title:    "Wall of Static",
		cardType: cardTypeIce,
		subtypes: []subtype{subtypeBarrier},
		text:     "SUBROUTINE: End the run.",
		faction:  factionNeutral,
		rezCost:  3,
		strength: 3,
	}
)

type card struct {
	player
	playCost int
	title    string
	subtitle string
	cardType
	subtypes []subtype
	text     string
	faction
	influenceValue         int
	minimumDeckSize        int
	influenceLimit         int
	advancementRequirement int
	agendaPoints           int
	rezCost                int
	trashCost              int
	strength               int
}
