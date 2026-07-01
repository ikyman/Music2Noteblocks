package mCraftnBlockmLearning

import (
	"fmt"
	_"github.com/gopxl/beep/v2"
	"gonum.org/v1/gonum/mat"

	"NoteblockRobot/Util"
)

const (
	ASSUMED_MAX_INSTRUMENTS int = 24
)

var(
	INSTUMENT_COLUMNS [ASSUMED_MAX_INSTRUMENTS]string;
)

func init(){
	INSTUMENT_COLUMNS = [ASSUMED_MAX_INSTRUMENTS]string{
		"harp",
		"bass",
		"snare",
		"hat",
		"basedrum",
		"bell",
		"flute",
		"chime",
		"guitar",
		"xylophone",
		"iron_xylophone",
		"cow_bell",
		"didgeridoo",
		"bit",
		"banjo",
		"pling",
		"trumpet",
		"trumpet_exposed",
		"trumpet_weathered",
		"trumpet_oxidized",
	}
}



func instrumentColumnIndex(instrumentName string) (int, bool) {
	canonicalName := utilitiesBeep.InstrumentNameFor(instrumentName)
	idx := 0
	for ; idx < len(INSTUMENT_COLUMNS); idx++{
		if canonicalName == INSTUMENT_COLUMNS[idx]{
			return idx, true
		}
	}
	return -1, false
}
// Half Written by the AI, who didn't understand me.
// The other half was written by me, who didn't understand the AI.
func SongSegment2Matrix(songSeg utilitiesBeep.SongSegment) mat.Matrix {
	// Target one row per sample in the segment span.
	nSamples := songSeg.EndLoc - songSeg.BeginLoc
	if nSamples <= 0 {
		return mat.NewDense(0, 1, nil)
	}
	retMatrix := mat.NewDense(nSamples, 1, nil)

	iterableStreamer := songSeg.AsBuffer().Streamer(0, nSamples)

	chunk := make([][2]float64, 512)

	iSample := 0;
	for ; iSample < nSamples;{
		num, _ := iterableStreamer.Stream(chunk)
		for i := 0; i < num && iSample < nSamples ; i+=2 {
			sampleVal := chunk[i][0] + 256*(chunk[i][1] + 256*(chunk[i+1][0] + 256*(chunk[i+1][1])))
			retMatrix.Set(iSample , 0, sampleVal)
			iSample +=2;
		}
	}

	return retMatrix
}

func SheetNote2Matrix(sheetNote utilitiesBeep.SheetNote) mat.Matrix {
	length := len(sheetNote.NotesByTime)
	retMatrix := mat.NewDense(length, ASSUMED_MAX_INSTRUMENTS, nil)

	for timeStep, notesAtTime := range sheetNote.NotesByTime {
		if notesAtTime == nil {
			continue
		}
		for instrumentName, pitch := range notesAtTime {
			col, ok := instrumentColumnIndex(instrumentName)
			if !ok {
				// unknown instrument
				continue
			}
			retMatrix.Set(timeStep, col, float64(pitch))
		}
	}

	return retMatrix
}
/*func matrix2SheetNote(mat mat.Matrix)  utilitiesBeep.SheetNote{

}*/

//func Matrix2BeepBuffer()
