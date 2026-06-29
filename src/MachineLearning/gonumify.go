package mCraftnBlockmLearning

import (
	_"github.com/gopxl/beep/v2"
	"gonum.org/v1/gonum/mat"

	"NoteblockRobot/Util"
)

const sheetNoteInstrumentCount = 24

// instrumentColumns maps column index -> canonical instrument name (harp -> 0).
// Columns 16–23 are reserved until the remaining instrument slots are confirmed.
var instrumentColumns = [sheetNoteInstrumentCount]string{
	"harp",
	"bass",
	"bd",
	"snare",
	"hat",
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
	"", "", "", "", "", "", "", "",
}

var instrumentNameToIndex map[string]int

func init() {
	instrumentNameToIndex = make(map[string]int, sheetNoteInstrumentCount+8)
	for i, name := range instrumentColumns {
		if name == "" {
			continue
		}
		instrumentNameToIndex[name] = i
	}

	// Common aliases and alternate spellings seen in sheet CSV headers.
	instrumentNameToIndex["basedrum"] = instrumentNameToIndex["bd"]
	instrumentNameToIndex["stone"] = instrumentNameToIndex["bd"]
	instrumentNameToIndex["sticks"] = instrumentNameToIndex["hat"]
	instrumentNameToIndex["hihat"] = instrumentNameToIndex["hat"]
}

func instrumentColumnIndex(instrumentName string) (int, bool) {
	canonical := utilitiesBeep.InstrumentNameFor(instrumentName)
	idx, ok := instrumentNameToIndex[canonical]
	return idx, ok
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
	retMatrix := mat.NewDense(length, sheetNoteInstrumentCount, nil)

	for timeStep, notesAtTime := range sheetNote.NotesByTime {
		if notesAtTime == nil {
			continue
		}
		for instrumentName, pitch := range notesAtTime {
			col, ok := instrumentColumnIndex(instrumentName)
			if !ok {
				// TODO: unknown instrument — define column mapping or extend instrumentColumns.
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
