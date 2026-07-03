
package utilitiesBeep

import(
	"fmt"
	"os"

	"time"
	"github.com/gopxl/beep/v2"

	"path/filepath"
)


func bufferSubsection(buff beep.Buffer, f beep.Format, startLoc float32 , endLoc float32 ) (beep.Streamer){
	subsectionBuffer := beep.NewBuffer(f);
	fullStream := buff.Streamer(0, buff.Len())
	subsectionBuffer.Append(fullStream);
	subsecStart := (f.SampleRate.N(time.Millisecond*time.Duration(int(1000*startLoc))));
	subsecStart = max(0, subsecStart)
	subsecEnd := (f.SampleRate.N(time.Millisecond*time.Duration(int(1000*endLoc))));
	subsecEnd = min(buff.Len(), subsecEnd)
	return subsectionBuffer.Streamer(subsecStart, subsecEnd);
}

func SliceSongIntoSegments(songRef *SongReference, segmentLengthSeconds float32) []SongSegment {
	if songRef == nil || songRef.SongBuffer == nil || segmentLengthSeconds <= 0 {
		return []SongSegment{}
	}

	songLengthSamples := songRef.SongBuffer.Len()
	if songLengthSamples <= 0 {
		return []SongSegment{}
	}

	segmentLengthSamples := songRef.GetSamplesInSeconds(segmentLengthSeconds)

	segments := make([]SongSegment, 0)
	for i := 0; i < songLengthSamples; i += segmentLengthSamples {
		endLoc := i + segmentLengthSamples
		if endLoc > songLengthSamples {
			endLoc = songLengthSamples
		}

		segments = append(segments, SongSegment{
			BeginLoc: i,
			EndLoc: endLoc,
			refersToSong: songRef,
		})
	}

	return segments
}

func main() {
	dir, _ := os.Getwd()
	fmt.Println("Working dir:", dir)

	/* Filepaths are based off of where powershell that calls the go run command was run.
	 * That'll take a little getting used to!
	 */
	oggAbsPath, _ := filepath.Abs("./TrainingTesting/trainingOggs/shop1DeltaRune.ogg") 
	fmt.Println(oggAbsPath)

}