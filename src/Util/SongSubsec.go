package utilitiesBeep

import(
	"log"
	"time"
	"os"

	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/vorbis"
	"github.com/gopxl/beep/v2/speaker"
)

type SongReference struct{
	SongBuffer *beep.Buffer
	songFormat beep.Format
}

// Note to self: Go find the codingGuru best practice for this sort of thing. Command? Visitor?
func LoadAudioFileOgg(oggFile string ) SongReference {
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

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	return SongReference{newBuffer, format};
}

type SongSegment struct{
	BeginLoc int
	EndLoc int

	refersToSong *SongReference
}

func (ss *SongSegment) GetSamplesInSeconds( seconds float32) int{
	return ss.refersToSong.GetSamplesInSeconds(seconds)
}

func (sr *SongReference) GetSamplesInSeconds( seconds float32) int{
	if sr.songFormat == (beep.Format{}){
		log.Fatal("Cannot Calculate Sample Without a format")
	}
	return sr.songFormat.SampleRate.N(time.Millisecond*time.Duration(1000 * seconds));
}

/*func (sr *SongReference) PlaySubsection( subsection SongSegment) int{

	tickStreamer := singletonAudioBuffer.Streamer(tickStart, tickStart + s10s.subSec.GetSamplesInSeconds(0.05));
	speaker.Play(tickStreamer);

	if sr.songFormat == (beep.Format{}){
		log.Fatal("Cannot Calculate Sample Without a format")
	}
	return sr.songFormat.SampleRate.N(time.Millisecond*time.Duration(1000 * seconds));
}*/
