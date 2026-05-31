package utilitiesBeep

import (
	"github.com/gopxl/beep/v2/speaker"
)

type Orchestra struct {
	orchestraInstruments map[string]Instrument
}

func emptyOrchestra() Orchestra {
	return Orchestra{
		orchestraInstruments: make(map[string]Instrument),
	}
}

func (orch *Orchestra) addInstrument(newInstrument Instrument) {
	orch.orchestraInstruments[newInstrument.instrumentName] = newInstrument
}

func (orch *Orchestra) playMusic(musicToPlay SheetNote) {
	for _, notesAtTime := range musicToPlay.NotesByTime {
		if notesAtTime == nil {
			continue
		}
		for instrumentName, pitch := range notesAtTime {
			instrument, ok := orch.orchestraInstruments[instrumentName]
			if !ok {
				continue
			}
			instrument.PlayNote(pitch)
		}
	}
}

func (orch *Orchestra) missingInstruments(musicToPlay SheetNote) []string {
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

type Instrument struct {
	midBeat        SongReference
	instrumentName string
}

func newInstrument(instrumentName string) Instrument {
	return Instrument{
		instrumentName: instrumentName,
		midBeat:        LoadAudioFileOgg("./InstrumentBeats" + instrumentName + ".ogg"),
	}
}

func (inst *Instrument) PlayNote(pitch int) {
	_ = pitch // pitch resampling to be added later
	streamer := inst.midBeat.SongBuffer.Streamer(0, inst.midBeat.SongBuffer.Len())
	speaker.Play(streamer)
}
