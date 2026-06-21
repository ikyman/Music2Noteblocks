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

func (orch *Orchestra) PlayMusic(musicToPlay utilitiesBeep.SheetNote) {
	for _, notesAtTime := range musicToPlay.NotesByTime {
		time.Sleep(time.Duration(1000/orch.beatsPerSecond)*time.Millisecond)
		if notesAtTime == nil {
			continue
		}
		for instrumentName, pitch := range notesAtTime {
			instrument, ok := orch.orchestraInstruments[instrumentName]

			if !ok {
				continue
			}
			go instrument.playNote(pitch)
		}
		
	}
}

func (orch *Orchestra) MissingInstruments(musicToPlay utilitiesBeep.SheetNote) []string {
	mentioned := make(map[string]struct{})
	for _, notesAtTime := range musicToPlay.NotesByTime {
		if notesAtTime == nil {
			continue
		}
		for instrumentName := range notesAtTime {
			mentioned[instrumentName] = struct{}{}
		}
	}

	var missing []string
	for instrumentName := range mentioned {
		if _, ok := orch.orchestraInstruments[instrumentName]; !ok {
			missing = append(missing, instrumentName)
		}
	}
	return missing
}

