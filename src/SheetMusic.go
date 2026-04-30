
package main

import(
	"fmt"
	"path/filepath"
	"NoteblockRobot/Util"

)

type MusicPage struct {
	// Warning! No methodology for checking that the Segment size is the same as the SheetNotes Size
	Segment utilitiesBeep.SongSegment
	SheetNotesCSV string // TODO: Temporary a String as I get SheetNote Working
}



// How It works: Ticks are the fundamental timespeed of minecraft. Thus, our representations shall just be "What note gets played in each 1/20th second interval"

func main (){
	absFilepath, _ := filepath.Abs("./TrainingTesting/trainingOggs")
	selectedFolder := pickSongFolder(absFilepath)
	filename := pickOggFileInFolder(selectedFolder)

	audioBuffer, format := utilitiesBeep.LoadAudioFileOgg(filename)

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	var selected10SecSeg *segment10Seconds
	for i:= 0 ; i  < singletonAudioBuffer.Len(); i += getSamplesInSeconds(10){
		// Sync Waitgroup? Unescissary!
		endLoc := i + getSamplesInSeconds(10)
		startSecond := i / getSamplesInSeconds(1)
		endSecond := endLoc / getSamplesInSeconds(1)
		csvPath := filepath.Join(selectedFolder, fmt.Sprintf("%dTo%d.csv", startSecond, endSecond))
		nss := newSegment10Seconds( i , endLoc, csvPath)
		buttons10SecondSubsects = append(buttons10SecondSubsects, nss);
	}
	selected10SecSeg = &buttons10SecondSubsects[0]

	tenSeconds := layout.List{Axis : layout.Vertical}
	listed10SecButtons := func(listContext layout.Context)layout.Dimensions{ return tenSeconds.Layout(listContext, len(buttons10SecondSubsects), 
		func(gtx layout.Context, index int) layout.Dimensions{ return buttons10SecondSubsects[index].draw10Buttons(gtx) }) }


}