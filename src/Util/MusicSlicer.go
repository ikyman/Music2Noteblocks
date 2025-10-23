
package main

import(
	"fmt"
	"log"
	"os"
	"time"
	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/vorbis"
	"github.com/gopxl/beep/v2/speaker"

	"path/filepath"
)


func bufferSubsection (buff beep.Buffer, f beep.Format, startLoc float32 , endLoc float32 ) (beep.Streamer){
	subsectionBuffer := beep.NewBuffer(f);
	fullStream := buff.Streamer(0, buff.Len())
	subsectionBuffer.Append(fullStream);
	subsecStart := (f.SampleRate.N(time.Millisecond*time.Duration(int(1000*startLoc))));
	subsecEnd := (f.SampleRate.N(time.Millisecond*time.Duration(int(1000*endLoc))));
	return subsectionBuffer.Streamer(subsecStart, subsecEnd);
}

// Note to self: Go find the codingGuru best practice for this sort of thing. Command? Visitor?
func loadAudioFileOgg(oggFile string ) (*beep.Buffer, beep.Format) {
	f, err := os.Open(oggFile)
	if err != nil {
		log.Fatal(err)
	}
	streamSeekClo, format, err := vorbis.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamSeekClo.Close()

	newBuffer := beep.NewBuffer(format);
	newBuffer.Append(streamSeekClo);
	
	return newBuffer, format;
}

func main() {
	dir, _ := os.Getwd()
	fmt.Println("Working dir:", dir)

	/* Filepaths are based off of where powershell that calls the go run command was run.
	 * That'll take a little getting used to!
	 */
	oggAbsPath, _ := filepath.Abs("./TrainingTesting/trainingOggs/shop1DeltaRune.ogg") 
	fmt.Println(oggAbsPath)

	b, f := loadAudioFileOgg(oggAbsPath)

	littleSection := bufferSubsection ( *b, f, 10, 20);

	speaker.Init(f.SampleRate, f.SampleRate.N(time.Second/10))

	speaker.Play(littleSection)

	select {}
}