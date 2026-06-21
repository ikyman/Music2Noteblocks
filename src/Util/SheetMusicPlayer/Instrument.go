package sheetMusicPlayer

import (
	"math"

	"github.com/gopxl/beep/v2"
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
	oldSampleRate := inst.midBeat.SongFormat.SampleRate;
	sampleRateRatio := math.Pow(2, float64(pitch-12)/float64(12)) // pitch resampling to be added later
	midBeat := inst.midBeat.SongBuffer.Streamer(0, inst.midBeat.SongBuffer.Len())
	pitchedNote := beep.Resample(3, oldSampleRate , beep.SampleRate(float64(oldSampleRate)/sampleRateRatio), midBeat)

	speaker.Play(pitchedNote)
}
