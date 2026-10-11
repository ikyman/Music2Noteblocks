package sheetMusicPlayer

import (
	"time"

	"NoteblockRobot/Util"
)

var (
	DEFAULT_ORCHESTRA_BPS float32;
)

func init() {
	DEFAULT_ORCHESTRA_BPS = 20;
}

type Orchestra struct {
	orchestraInstruments map[string]Instrument

	beatsPerSecond float32;
}

func EmptyOrchestra() Orchestra {
	return Orchestra{
		orchestraInstruments: make(map[string]Instrument),
		beatsPerSecond: DEFAULT_ORCHESTRA_BPS,
	}
}

func (orch *Orchestra) AddInstrument(newInstrument Instrument) {
	orch.orchestraInstruments[newInstrument.instrumentName] = newInstrument
}

func (orch *Orchestra) PlayMusic(musicToPlay utilitiesBeep.SheetNote, indexer utilitiesBeep.IndexerInstruments) {
	total_ticks, cols := musicToPlay.Dims()

	for tick := range total_ticks {
		time.Sleep(time.Duration(1000/orch.beatsPerSecond)*time.Millisecond)
		for col := range cols{
			pitch := int(musicToPlay.At(tick, col))
			if pitch == 0 {
				continue
			}
			instrumentName := indexer[col]
			instrument, ok := orch.orchestraInstruments[instrumentName]
			if !ok {
				continue
			}
			go instrument.playNote(pitch)
		}
	}
}

func (orch *Orchestra) MissingInstruments(indexer utilitiesBeep.IndexerInstruments) []string {
	mentioned := make([]string, 0)
	for _, instrumentName := range indexer {
		mentioned = append(mentioned, instrumentName)
	}

	var missing []string
	for _, instrumentName := range mentioned {
		if _, ok := orch.orchestraInstruments[instrumentName]; !ok {
			missing = append(missing, instrumentName)
		}
	}
	return missing
}

