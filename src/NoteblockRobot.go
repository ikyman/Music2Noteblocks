
package main

import(
	"fmt"
	"path/filepath"
	"os"
	"NoteblockRobot/Util"
	_"NoteblockRobot/MachineLearning"

)

func main(){
	absFilepath, _ := filepath.Abs("./TrainingTesting/trainingOggs")

	dirEntries, _ := os.ReadDir(absFilepath)

	for _, dirEntry := range dirEntries{
		if (dirEntry.IsDir()){
			utilitiesBeep.LoadTrainingOGGFolder(absFilepath + "/" + dirEntry.Name())
		}
	}
}