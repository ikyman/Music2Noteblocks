package main

import(
	"fmt"
	"os"
	"log"
	"path/filepath"
	
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

	singletonSongReference utilitiesBeep.SongReference
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
	play10SecTheme.Palette.ContrastBg  = color.NRGBA{R: 200, G: 255, B: 200, A: 255}
	play10SecTheme.Palette.ContrastFg  = color.NRGBA{R: 0, G: 0, B: 0, A: 255}

	playTickTheme = material.NewTheme()
	playTickTheme.Palette.ContrastBg  = color.NRGBA{R: 100, G: 100, B: 100, A: 255}

}

type segment10Seconds struct{
	subSec utilitiesBeep.SongSegment 
	select10Sec widget.Clickable
	play10Sec widget.Clickable

	tickButtons [][]widget.Clickable
	secondsList layout.List
	tickLists []layout.List
}

func newSegment10Seconds(segment utilitiesBeep.SongSegment) segment10Seconds{
	nsb := new(segment10Seconds);
	nsb.subSec = segment

	nsb.tickButtons = make([][]widget.Clickable, 10)
	nsb.tickLists = make([]layout.List, 10)
	for sec:=0; sec < 10; sec = sec+1{
		nsb.tickButtons[sec] = make([]widget.Clickable, 20)
		nsb.tickLists[sec] = layout.List{Axis : layout.Vertical}
	}

	return *nsb
}

func (s10s *segment10Seconds) draw10Buttons(buttonContext layout.Context) layout.Dimensions{
	selectButtonVisual := material.Button(select10SecTheme, &s10s.select10Sec, fmt.Sprintf("Select: From %d to %d" , s10s.subSec.BeginLoc, s10s.subSec.EndLoc )  )
	playButtonVisual := material.Button(play10SecTheme, &s10s.play10Sec, "Play" )

	return layout.Flex{}.Layout(buttonContext, layout.Flexed(1, playButtonVisual.Layout ), layout.Flexed(2, selectButtonVisual.Layout))
}

func (s10s *segment10Seconds) handle10Clicks(clickTracker layout.Context, selected10SecSeg *segment10Seconds)  *segment10Seconds{
	if s10s.play10Sec.Clicked(clickTracker){
		streamerPlayable := singletonAudioBuffer.Streamer(s10s.subSec.BeginLoc, s10s.subSec.EndLoc);
		speaker.Play(streamerPlayable);
	}

	if s10s.select10Sec.Clicked(clickTracker){	
		return s10s
	}
	return selected10SecSeg
}


func (s10s *segment10Seconds)  drawTickButtons(buttonContext layout.Context) layout.Dimensions{
	initialSec := s10s.subSec.BeginLoc/(20 * s10s.subSec.GetSamplesInSeconds(0.05))
	return s10s.secondsList.Layout(buttonContext, 10, 
		func(gtx layout.Context, secIndex int)layout.Dimensions{
			return s10s.tickLists[secIndex].Layout(gtx, 20,
				func(gtx2 layout.Context, tickIndex int)layout.Dimensions{
					return material.Button(playTickTheme, &s10s.tickButtons[secIndex][tickIndex], fmt.Sprintf("Second %d, tick %d", initialSec+secIndex, tickIndex) ).Layout(gtx2)
				},
			)
		},
	)
}

func (s10s *segment10Seconds)  handleClickTicks(clickTracker layout.Context){
	for sec := 0; sec < 10; sec = sec+1 {
		for tick :=0; tick < 20; tick = tick + 1{
			if s10s.tickButtons[sec][tick].Clicked(clickTracker){
				tickStart := s10s.subSec.BeginLoc + s10s.subSec.GetSamplesInSeconds(float32(sec)) + (tick*s10s.subSec.GetSamplesInSeconds(0.05))
				tickStreamer := singletonAudioBuffer.Streamer(tickStart, tickStart + s10s.subSec.GetSamplesInSeconds(0.05));
				speaker.Play(tickStreamer);
			}
		} 
	}
}

func RunSliceGUI(){
	absFilepath, _ := filepath.Abs("./TrainingTesting/trainingOggs")
	filename, err := dialog.File().SetStartDir(absFilepath).Load()
	if err != nil {
		log.Fatal(err)
	}

	singletonSongReference = utilitiesBeep.LoadAudioFileOgg(filename) 
	singletonAudioBuffer = singletonSongReference.SongBuffer

	buttons10SecondSubsects := make([]segment10Seconds,0)
	
	var selected10SecSeg *segment10Seconds
	tenSecondSegments := utilitiesBeep.SliceSongIntoSegments(&singletonSongReference, 10)
	for _, segment := range tenSecondSegments {
		nss := newSegment10Seconds(segment)
		buttons10SecondSubsects = append(buttons10SecondSubsects, nss);
	}
	if len(buttons10SecondSubsects) == 0 {
		log.Fatal("No 10-second segments were created for this audio file")
	}
	selected10SecSeg = &buttons10SecondSubsects[0]

	tenSeconds := layout.List{Axis : layout.Vertical}
	listed10SecButtons := func(listContext layout.Context)layout.Dimensions{ return tenSeconds.Layout(listContext, len(buttons10SecondSubsects), 
		func(gtx layout.Context, index int) layout.Dimensions{ return buttons10SecondSubsects[index].draw10Buttons(gtx) }) }


	go func(){
		w := new(app.Window)
		ops := new(op.Ops)

		for {
			evt := w.Event()

			switch typ := evt.(type){
			case app.FrameEvent:
				flexContext := app.NewContext( ops, typ)
				for i, _ := range(buttons10SecondSubsects){
					selected10SecSeg = buttons10SecondSubsects[i].handle10Clicks(flexContext, selected10SecSeg)
					selected10SecSeg.handleClickTicks(flexContext)
				}

				labelTop := layout.Flex{Axis : layout.Vertical}
				buttonSegregator := func(gtx layout.Context) layout.Dimensions{
					return layout.Flex{}.Layout(gtx,
						layout.Flexed(1, listed10SecButtons),
						layout.Flexed(3, selected10SecSeg.drawTickButtons))
				}
			
				labelTop.Layout(flexContext, layout.Rigid( material.Label(labelTheme, 14, filename).Layout), layout.Flexed(1, buttonSegregator),
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