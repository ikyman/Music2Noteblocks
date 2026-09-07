
package main

import(
	"fmt"
	"path/filepath"
	"os"
	"NoteblockRobot/Util"
	"NoteblockRobot/Util/SheetMusicPlayer"
	_"NoteblockRobot/MachineLearning"

)

var (

)

func init() {

}

func main(){
	var user_input string;
	fmt.Println("Golang Noteblock to Robot application commad line interface")
	fmt.Println("l : Listen to Training music")
	fmt.Println("")

	fmt.Scanln(&user_input)
	if (user_input == "l"){
		absFilepath, _ := filepath.Abs("./TrainingTesting/trainingOggs")

		dirEntries, _ := os.ReadDir(absFilepath)
		var sheetMusics []utilitiesBeep.MusicPage;
	
		for _, dirEntry := range dirEntries{
			if (dirEntry.IsDir()){
				sheetMusics = utilitiesBeep.LoadTrainingOGGFolder(absFilepath + "/" + dirEntry.Name())
			}
		}
		orchestra := sheetMusicPlayer.EmptyOrchestra();
		orchestra.AddInstrument(sheetMusicPlayer.NewInstrument("harp"));
	
		orchestra.PlayMusic(sheetMusics[0].SheetNotesCSV);
	}

}