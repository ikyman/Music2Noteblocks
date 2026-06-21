package sheetMusicPlayer

import (
	"github.com/gopxl/beep/v2/speaker"

	"NoteblockRobot/Util"
)

type Instrument struct {
	midBeat        utilitiesBeep.SongReference
	instrumentName string
}

func NewInstrument(instrumentName string) Instrument {
	return Instrument{
		instrumentName: utilitiesBeep.InstrumentNameFor(instrumentName),
		midBeat:        utilitiesBeep.LoadAudioFileOgg("./Util/InstrumentBeats/" + instrumentName + ".ogg"),
	}
}

func (inst *Instrument) playNote(pitch int) {
	_ = pitch // pitch resampling to be added later
	streamer := inst.midBeat.SongBuffer.Streamer(0, inst.midBeat.SongBuffer.Len())
	speaker.Play(streamer)
}
