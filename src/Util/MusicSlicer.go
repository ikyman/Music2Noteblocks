
package utilitiesBeep

func SliceSongIntoSegments(songRef *SongReference, segmentLengthSeconds float32) []SongSegment {
	if songRef == nil || songRef.SongBuffer == nil || segmentLengthSeconds <= 0 {
		return []SongSegment{}
	}

	songLengthSamples := songRef.SongBuffer.Len()
	if songLengthSamples <= 0 {
		return []SongSegment{}
	}

	segmentLengthSamples := songRef.GetSamplesInSeconds(segmentLengthSeconds)

	segments := make([]SongSegment, 0)
	for i := 0; i < songLengthSamples; i += segmentLengthSamples {
		endLoc := i + segmentLengthSamples
		if endLoc > songLengthSamples {
			endLoc = songLengthSamples
		}

		segments = append(segments, SongSegment{
			BeginLoc: i,
			EndLoc: endLoc,
			refersToSong: songRef,
		})
	}

	return segments
}