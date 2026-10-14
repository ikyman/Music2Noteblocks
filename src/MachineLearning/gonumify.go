package mCraftnBlockmLearning

import (
	_"fmt"
	_"github.com/gopxl/beep/v2"
	"gonum.org/v1/gonum/mat"

	"NoteblockRobot/Util"
)

func init(){
	INSTUMENT_COLUMNS = [ASSUMED_MAX_INSTRUMENTS]string{
		"trumpet",
		"trumpet_exposed",
		"trumpet_weathered",
		"trumpet_oxidized",
	}
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
	return sheetNote
}

//func Matrix2BeepBuffer()
