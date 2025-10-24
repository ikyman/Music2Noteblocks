package main

import(
	"fmt"
	"gioui.org/app"
	"os"
	"github.com/sqweek/dialog"
	"log"
	"path/filepath"
	"time"
	"NoteblockRobot/Util"
	"github.com/gopxl/beep/v2"

)

type songSegment  struct{
	segment beep.Streamer
	beginLoc float32
	endLoc float32
}

type segmentButton struct{
	songSeg songSegment
	songButt widget.Clickable
}

func  newSegmentButton(segment beep.Streamer, beginLoc float32, endLoc float32) segmentButton{
	nsb := new(segmentButton);
	nsb.songSeg := newSongSegment(segment beep.Streamer, beginLoc float32, endLoc float32);
	var newButton widget.Clickable;
	nsb.songButt := newButton

	return *nsb
}

func newSongSegment(segment beep.Streamer, beginLoc float32, endLoc float32) songSegment{
	nss := new(songSegment)
	nss.segment = segment;
	nss.beginLoc = beginLoc;
	nss.endLoc = endLoc;
	return *nss
}

func main(){
	fmt.Println("Hello World! (Go is oddly hard to get set up)");
	
	absFilepath, _ := filepath.Abs("./TrainingTesting/trainingOggs")
	filename, err := dialog.File().SetStartDir(absFilepath).Load()
	if err != nil {
		log.Fatal(err)
	}

	fullSong, format := utilitiesBeep.LoadAudioFileOgg(filename);

	totalSongSeconds := fullSong.Len()/format.SampleRate.N(time.Second)

	all10Subsections := make([]segmentButton,1)

	for i:= 0 ; i * format.SampleRate.N(time.Second)  < fullSong.Len(); i +=10{
		go func(){
			endLoc := i + 10;
			additional10Secs := fullSong.Streamer(i, endLoc)
			nss := newSongSegment (additional10Secs, float32(i), float32(endLoc))
			all10Subsections = append(all10Subsections, nss);
			fmt.Println("Appended!");
		}()
	}

	for i, v := range(all10Subsections){
		fmt.Println("Index ", i,  " starts at", v.beginLoc );
	}


	fmt.Println("This Song is ", totalSongSeconds, "Seconds long, ")//it will be sliced into ", , " subsections. Put together, it's", ,  " seconds")

	go func(){
		w := new(app.Window)
		ops := new(op.Ops)

		for {
			evt := w.Event()


			switch typ := evt.(type){
			case app.FrameEvent:
				lyoutCntxt = w.NewContext(& ops, typ)

			case app.DestroyEvent:
				fmt.Println(typ);
				os.Exit(0)
			}
		}
	}()
	app.Main()

}