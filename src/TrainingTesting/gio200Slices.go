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

// Application themes - initialized once, accessible throughout the package
var (
	labelTheme  *material.Theme
	select10SecTheme *material.Theme
	play10SecTheme *material.Theme
	playTickTheme *material.Theme

	singletonAudioBuffer *beep.Buffer
	singletonAudioFormat *beep.Format
)

func init() {
	// Initialize themes once at package load time
	labelTheme = material.NewTheme()
	labelTheme.Palette = material.Palette{
		Bg: color.NRGBA{R: 100, G: 0, B: 0, A: 255},
		Fg: color.NRGBA{R: 0, G: 0, B: 0, A: 255},
	}
	
	select10SecTheme = material.NewTheme()

	play10SecTheme = material.NewTheme()
	play10SecTheme.Palette.Bg = color.NRGBA{R: 200, G: 255, B: 200, A: 255}

	playTickTheme = material.NewTheme()
}

func getSamplesInTick() int{
	if singletonAudioFormat == nil{
		log.Fatal("Cannot Calculate Sample Without a format")
	}

	return singletonAudioFormat.SampleRate.N(time.Millisecond*time.Duration(50));
}
func getSamplesIn10Seconds() int{
	return 200 * getSamplesInTick()
}


type segment10Seconds struct{
	beginLoc int
	endLoc int
	select10Sec widget.Clickable;
	play10Sec widget.Clickable
}

func newSegment10Seconds(beginLoc int, endLoc int) segment10Seconds{
	nsb := new(segment10Seconds);
	nsb.beginLoc = beginLoc
	nsb.endLoc = endLoc

	return *nsb
}

func (s10s *segment10Seconds) drawButtons(buttonContext layout.Context) layout.Dimensions{
	buttonVisual := material.Button(select10SecTheme, &s10s.select10Sec, fmt.Sprintf("From %d to %d" , s10s.beginLoc, s10s.endLoc )  )
	// material.Button(play10SecTheme, &sb.songButt, "Listen" )

	return buttonVisual.Layout(buttonContext)
}

func (s10s *segment10Seconds) handleClicks(clickTracker layout.Context){
	if s10s.select10Sec.Clicked(clickTracker){	
		fmt.Println("No longer my Job!")
		//speaker.Play(sb.songSeg.segment)
	}
}


func main(){
	absFilepath, _ := filepath.Abs("./TrainingTesting/trainingOggs")
	filename, err := dialog.File().SetStartDir(absFilepath).Load()
	if err != nil {
		log.Fatal(err)
	}

	var format beep.Format
	singletonAudioBuffer, format = utilitiesBeep.LoadAudioFileOgg(filename)
	singletonAudioFormat = &format

	speaker.Init(singletonAudioFormat.SampleRate, singletonAudioFormat.SampleRate.N(time.Second/10))

	buttons10SecondSubsects := make([]segment10Seconds,0)

	for i:= 0 ; i  < singletonAudioBuffer.Len(); i += getSamplesIn10Seconds(){
		// Sync Waitgroup? Unescissary!
		endLoc := i + getSamplesIn10Seconds()
		nss := newSegment10Seconds( i , endLoc)
		buttons10SecondSubsects = append(buttons10SecondSubsects, nss);
	}

	tenSeconds := layout.List{Axis : layout.Vertical}
	listed10SecButtons := func(listContext layout.Context)layout.Dimensions{ return tenSeconds.Layout(listContext, len(buttons10SecondSubsects), 
		func(gtx layout.Context, index int) layout.Dimensions{ return buttons10SecondSubsects[index].drawButtons(gtx) }) }


	go func(){
		w := new(app.Window)
		ops := new(op.Ops)

		for {
			evt := w.Event()

			switch typ := evt.(type){
			case app.FrameEvent:
				flexContext := app.NewContext( ops, typ)
				for i, _ := range(buttons10SecondSubsects){
					buttons10SecondSubsects[i].handleClicks(flexContext)
				}

				labelTop := layout.Flex{Axis : layout.Vertical}
				buttonSegregator := func(gtx layout.Context) layout.Dimensions{ return layout.Flex{}.Layout(gtx,  layout.Flexed(1, listed10SecButtons)) }
				//clickTracker := layout.Context{Ops : ops}
				//var flexPosting layout.Flex

								
				labelTop.Layout(flexContext, layout.Rigid( material.Label(labelTheme, 14, "absFilepath").Layout), layout.Flexed(1, buttonSegregator),
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