package utilitiesBeep

import(
	"fmt"
	"path/filepath"
	"log"
)


// Application themes - initialized once, accessible throughout the package
var (
	musicSliceSecondLength  int
)

func init() {
	// Initialize themes once at package load time
	musicSliceSecondLength = 10;
}

type MusicPage struct {
	Segment SongSegment
	SheetNotesCSV SheetNote 
}

// How It works: Ticks are the fundamental timespeed of minecraft. Thus, our representations shall just be "What note gets played in each 1/20th second interval"
func LoadTrainingOGGFolder(trainingFolderFilePath string ) []MusicPage{
	// Warning! No methodology for checking that the Segment size is the same as the SheetNotes Size
	musicPages := make([]MusicPage, 0);

	oggFilename := filepath.Join(trainingFolderFilePath, filepath.Base(trainingFolderFilePath)+".ogg")

	loadedSongReference := LoadAudioFileOgg(oggFilename)

	musicAudioSegments := SliceSongIntoSegments(&loadedSongReference, float32(musicSliceSecondLength));

	for i, _ := range musicAudioSegments {
		csvPath := filepath.Join(trainingFolderFilePath, fmt.Sprintf("%dTo%d.csv", i*musicSliceSecondLength, (i+1)*musicSliceSecondLength))
		sheetNotesCSV, err := LoadSheetNoteFromCSV(csvPath)
		if err != nil {
			log.Fatal(err)
		}
		musicPages = append(musicPages, MusicPage{
			Segment: musicAudioSegments[i],
			SheetNotesCSV: sheetNotesCSV,
		})
	}

	return musicPages
}

/*func main(){
	absFilepath, _ := filepath.Abs("./TrainingTesting/trainingOggs")

	selectedFolder, err := dialog.Directory().SetStartDir(absFilepath).Browse()
	if err != nil {
		log.Fatal(err)
	}

	musicPages := LoadTrainingOGGFolder(selectedFolder);

	orchest := sheetMusicPlayer.EmptyOrchestra();
	orchest.AddInstrument(sheetMusicPlayer.NewInstrument("harp")) 

	//orchest.PlayMusic(musicPages[0].SheetNotesCSV);

	fmt.Println(mCraftnBlockmLearning.SheetNote2Matrix(musicPages[0].SheetNotesCSV))

	
}*/