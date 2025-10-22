
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

func sliceBySeconds (s beep.Streamer, f beep.Format, startLoc float32 , endLoc float32 ){
	newBuffer := beep.NewBuffer(f);
	newBuffer.Append(s);
	streamSeekClo := newBuffer.Streamer(int(startLoc), int(endLoc));
	//defer streamSeekClo.Close()
	
	speaker.Init(f.SampleRate, f.SampleRate.N(time.Second/10))

	speaker.Play(streamSeekClo)

	select {}

}

// Note to self: Go find the codingGuru best practice for this sort of thing. Command? Visitor?
func loadAudioFileOgg(oggFile string ) (beep.Streamer, beep.Format) {
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
	fmt.Println("format Width= " , format.Width())
	fmt.Println("How big is my new buffer?", newBuffer.Len())
	snippitStart := format.SampleRate.N(time.Second*10);
	snippitEnd := format.SampleRate.N(time.Second*20);
	fmt.Println("Should GO from ", snippitStart, " to  ", snippitEnd);
	streamSeek := newBuffer.Streamer(snippitStart, snippitEnd);
	fmt.Println("How big is my new streamer?", streamSeek.Len())
	
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	// fmt.Println("This should play the whole thing")
	// speaker.Play(streamSeekClo)

	fmt.Println("This should play a snippit")
	speaker.Play(streamSeek)

	select {}
	//defer streamSeekClo.Close()
	
	return streamSeekClo, format;
}

func main() {
	dir, _ := os.Getwd()
	fmt.Println("Working dir:", dir)

	/* Filepaths are based off of where powershell that calls the go run command was run.
	 * That'll take a little getting used to!
	 */
	oggAbsPath, _ := filepath.Abs("./TrainingTesting/trainingOggs/shop1DeltaRune.ogg") 
	fmt.Println(oggAbsPath)

	s, f := loadAudioFileOgg(oggAbsPath)

	sliceBySeconds ( s, f, 10, 20);

}