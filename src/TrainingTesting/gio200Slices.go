package main

import(
	"fmt"
	"gioui.org/app"
	"os"
	"github.com/sqweek/dialog"
	"log"
	"path/filepath"
)


func main (){
	fmt.Println("Hello World! (Go is oddly hard to get set up)");
	
	absFilepath, _ := filepath.Abs("./TrainingTesting/trainingOggs")
	filename, err := dialog.File().SetStartDir(absFilepath).Load()
	if err != nil {
		log.Fatal(err)
	}

	//for 

	fmt.Println("File chosen: ", filename)

	go func(){
		w := new(app.Window)
		for {
			evt := w.Event()

			switch typ := evt.(type){
			case app.DestroyEvent:
				fmt.Println(typ);
				os.Exit(0)
			}
		}
	}()
	app.Main()

}