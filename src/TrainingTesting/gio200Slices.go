package main

import(
	"fmt"
	"os"
	"log"
	"path/filepath"
	"time"
	
	"github.com/sqweek/dialog"
	"NoteblockRobot/Util"
	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/speaker"

	"image/color"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/layout"
	"gioui.org/widget/material"
	"gioui.org/widget"
)


type songSegment  struct{
	segment beep.Streamer
	beginLoc float32
	endLoc float32
}

func newSongSegment(segment beep.Streamer, beginLoc float32, endLoc float32) songSegment{
	nss := new(songSegment)
	nss.segment = segment;
	nss.beginLoc = beginLoc;
	nss.endLoc = endLoc;
	return *nss
}

type segmentButton struct{
	songSeg songSegment
	songButt widget.Clickable
}

func newSegmentButton(segment beep.Streamer, beginLoc float32, endLoc float32) segmentButton{
	nsb := new(segmentButton);
	nsb.songSeg = newSongSegment(segment, beginLoc, endLoc);
	var newButton widget.Clickable;
	nsb.songButt = newButton

	return *nsb
}

func (sb *segmentButton) drawButton (buttonContext layout.Context) layout.Dimensions{
	th := material.NewTheme()
	buttonVisual := material.Button(th, &sb.songButt, fmt.Sprintf("From %d to %d" , sb.songSeg.beginLoc, sb.songSeg.endLoc )  )
	return buttonVisual.Layout(buttonContext)
}

func (sb *segmentButton) handleClicks(clickTracker layout.Context){
	if sb.songButt.Clicked(clickTracker){
		speaker.Play(sb.songSeg.segment)	
		fmt.Println("button clicked!")
	}
}


func main(){
	var labelTheme = material.NewTheme()
	labelTheme.Palette = material.Palette{Bg : color.NRGBA{R: 100, G:0, B: 0, A:255}, Fg:color.NRGBA{R: 0, G:0, B: 0, A:255}  } 

	absFilepath, _ := filepath.Abs("./TrainingTesting/trainingOggs")
	filename, err := dialog.File().SetStartDir(absFilepath).Load()
	if err != nil {
		log.Fatal(err)
	}

	fullSong, format := utilitiesBeep.LoadAudioFileOgg(filename);

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	buttons10SecondSubsects := make([]segmentButton,0)

	for i:= 0 ; i * format.SampleRate.N(time.Second)  < fullSong.Len(); i +=10{
		// Sync Waitgroup? Unescissary!
		endLoc := i + 10;
		additional10Secs := fullSong.Streamer(i, endLoc)
		nss := newSegmentButton(additional10Secs, float32(i), float32(endLoc))
		buttons10SecondSubsects = append(buttons10SecondSubsects, nss);
	}

	go func(){
		w := new(app.Window)
		ops := new(op.Ops)

		for {
			evt := w.Event()

			switch typ := evt.(type){
			case app.FrameEvent:
				flexContext := app.NewContext( ops, typ)

				labelTop := layout.Flex{Axis : layout.Vertical}
				//var buttonSegregator layout.Flex
				//var tenSeconds layout.List
				clickTracker := layout.Context{Ops : ops}
				var flexPosting layout.Flex
				buttonDims:= make([]layout.FlexChild,0)

				for _, button := range(buttons10SecondSubsects){
					button.handleClicks(clickTracker);
					buttonDims = append(buttonDims, layout.Flexed(1, button.drawButton ) )
				} 

				labelTop.Layout(flexContext, layout.Rigid( material.Label(labelTheme, 14, "absFilepath").Layout), 
				layout.Flexed( 5,  func(gtx layout.Context)layout.Dimensions{ return  flexPosting.Layout(gtx, buttonDims... ) } ),
				)

				typ.Frame(ops)

			case app.DestroyEvent:
				fmt.Println(typ);
				os.Exit(0)
			}
		}
	}()
	app.Main()

}