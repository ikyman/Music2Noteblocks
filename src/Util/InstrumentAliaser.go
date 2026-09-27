package utilitiesBeep

import(
	"strings"
)

type IndexerInstruments = map[int]string

var (
	ASSUMED_MAX_INSTRUMENTS int

	DEFAULT_INDEXER_INSTRUMETS IndexerInstruments

	instrumentNameAliaser map[string]string
)

func init() {
	ASSUMED_MAX_INSTRUMENTS = 32
	
	DEFAULT_INDEXER_INSTRUMETS = IndexerInstruments{
		0:  "harp",
		1:  "bass",
		2:  "bd",
		3:  "hat",
		4:  "snare",
		5:  "pling",
		6:  "8bit",
		7:  "banjo",
		8:  "bell",
		9:  "chimes",
		10: "cowbell",
		11: "didgeridoo",
		12: "guitar",
		13: "flute",
		14: "iron_xylophone",
		15: "xylophone",
	}

	instrumentNameAliaser = map[string]string{
		"stone":    "bd",
		"basedrum": "bd",
	}
}

func InstrumentAliasFor(unaliasedName string) string {
	lowercaseUnaliased := strings.ToLower(unaliasedName)

	aliasedName := instrumentNameAliaser[lowercaseUnaliased]
	if aliasedName != "" {
		return aliasedName
	}
	return lowercaseUnaliased
}
