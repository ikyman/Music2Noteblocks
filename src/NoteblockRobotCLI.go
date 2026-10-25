
package main

import(
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/sqweek/dialog"
	"github.com/gopxl/beep/v2/speaker"

	"NoteblockRobot/Util"
	"NoteblockRobot/Util/SheetMusicPlayer"
	_"NoteblockRobot/MachineLearning"
)

var (

)

func init() {

}

func main(){
	orchest := sheetMusicPlayer.EmptyOrchestra();
	orchest.AddInstrument(sheetMusicPlayer.NewInstrument("harp")) 

	/* Filepaths are based off of where powershell that calls the go run command was run.
	 * That'll take a little getting used to!
	 */
	testTrainFilePath, _ := filepath.Abs("./TrainingTesting/trainingOggs")

	dir, _ := os.Getwd()
	fmt.Println("Working dir:", dir)
	fmt.Println("\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\")

	var user_input string;

	fmt.Println("Golang Noteblock to Robot application commad line interface")
	fmt.Println("l : Listen to either a csv or an ogg, depending on which file is chosen.") 
	fmt.Println("lp : Paired Listening: a .ogg playing at the same time as the .csv")
	/*fmt.Println("lra : Listen to raw audiofile")
	fmt.Println("lrc : Listen to raw csv")*/
	fmt.Println("s : Slice music with the music Gio UI-ed music slicer")
	fmt.Println("t : Attempt training using the default training folder.")
	fmt.Println("b : Create Mineflayer bot from a CSV or collection of CSVs")

	fmt.Scanln(&user_input)
	if (user_input == "l"){
		filename, err := dialog.File().SetStartDir(testTrainFilePath).Load()
		if err != nil {
			log.Fatal(err)
		}

		fileExtension := filename[len(filename)-4:]

		if (fileExtension) == ".csv"{
			sn, ii, _ := utilitiesBeep.LoadSheetNoteFromCSV(filename)
			orchest.PlayMusic(sn,ii)
		}else if (fileExtension) == ".ogg"{
			songRef := utilitiesBeep.LoadAudioFileOgg(filename)

			songRefStreamer := songRef.SongBuffer.Streamer(0, songRef.SongBuffer.Len())
			speaker.Init(songRef.SongFormat.SampleRate, 90)
			speaker.PlayAndWait(songRefStreamer);
		}else{
			log.Fatal("Unknown file extension!")
		}
	}

	if user_input == "lp"{
		selectedFolder, err := dialog.Directory().SetStartDir(testTrainFilePath).Browse()
		if err != nil {
			log.Fatal(err)
		}
	
		musicPair := utilitiesBeep.LoadTrainingOGGFolder(selectedFolder)
		if len(musicPair) == 0 {
			log.Fatal("LoadTrainingOGGFolder returned no AudioSheetIndexer slices")
		}

		firstPair := musicPair[0]
		audioBuffer := firstPair.Audio.AsBuffer()
		speaker.Play(audioBuffer.Streamer(0, audioBuffer.Len()))
		orchest.PlayMusic(firstPair.Sheet, firstPair.Indexer)

		log.Fatal("TODO: loop over the full []AudioSheetIndexer from LoadTrainingOGGFolder, not only index 0")
	}

	if (user_input == "s"){
		RunSliceGUI()
	}
	if (user_input == "t"){
		dirEntries, _ := os.ReadDir(testTrainFilePath)
		var musicPairs []utilitiesBeep.AudioSheetIndexer;
	
		for _, dirEntry := range dirEntries{
			if (dirEntry.IsDir()){
				musicPairs = utilitiesBeep.LoadTrainingOGGFolder(testTrainFilePath + "/" + dirEntry.Name())
			}
		}
		fmt.Println(musicPairs)
		//orchest.PlayMusic(musicPairs[0].SheetNotesCSV);

		log.Fatal("Jumble of Matrixies not yet ready for backpropagation!")
	}

	if (user_input == "b"){
		log.Fatal("Custom Mineflayer Bots not yet created!")

	}

}