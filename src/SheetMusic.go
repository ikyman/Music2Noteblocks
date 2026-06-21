package main

import(
	"fmt"
	"path/filepath"
	"github.com/sqweek/dialog"
	"log"

	"NoteblockRobot/Util"
	"NoteblockRobot/Util/SheetMusicPlayer"
)

type MusicPage struct {
	// Warning! No methodology for checking that the Segment size is the same as the SheetNotes Size
	Segment utilitiesBeep.SongSegment
	SheetNotesCSV utilitiesBeep.SheetNote 
}

// Application themes - initialized once, accessible throughout the package
var (
	musicSliceSecondLength  int
)

func init() {
	// Initialize themes once at package load time
	musicSliceSecondLength = 10;
}

// How It works: Ticks are the fundamental timespeed of minecraft. Thus, our representations shall just be "What note gets played in each 1/20th second interval"

func main(){
	absFilepath, _ := filepath.Abs("./TrainingTesting/trainingOggs")

	//selectedFolder := pickSongFolder(absFilepath)
	selectedFolder, err := dialog.Directory().SetStartDir(absFilepath).Browse()
	if err != nil {
		log.Fatal(err)
	}

	oggFilename := filepath.Join(selectedFolder, filepath.Base(selectedFolder)+".ogg")

	loadedSongReference := utilitiesBeep.LoadAudioFileOgg(oggFilename)

	musicAudioSegments := utilitiesBeep.SliceSongIntoSegments(&loadedSongReference, float32(musicSliceSecondLength));
	musicPages := make([]MusicPage, 0);

	for i, _ := range musicAudioSegments {
		csvPath := filepath.Join(selectedFolder, fmt.Sprintf("%dTo%d.csv", i*musicSliceSecondLength, (i+1)*musicSliceSecondLength))
		sheetNotesCSV, err := utilitiesBeep.LoadSheetNoteFromCSV(csvPath)
		if err != nil {
			log.Fatal(err)
		}
		musicPages = append(musicPages, MusicPage{
			Segment: musicAudioSegments[i],
			SheetNotesCSV: sheetNotesCSV,
		})
	}

	orchest := sheetMusicPlayer.EmptyOrchestra();
	orchest.AddInstrument(sheetMusicPlayer.NewInstrument("harp")) 

	orchest.PlayMusic(musicPages[0].SheetNotesCSV);
}