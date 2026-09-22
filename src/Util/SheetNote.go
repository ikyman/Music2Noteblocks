package utilitiesBeep

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"gonum.org/v1/gonum/mat"
)

type SheetNote struct {
	// index = time step; each entry is instrument -> pitch at that step
	NotesByTime []map[string]int

	CoveredLength int
}

type SheetNote = mat.Matrix;

func emptySheetNote() SheetNote {
	return SheetNote{
		NotesByTime: make([]map[string]int, 0),
	}
}

// This must be AI generated, for I don't have the foggiest clue what the first half means.
func ensureTimeSlot(sheet *SheetNote, timestamp int) {
	for len(sheet.NotesByTime) <= timestamp {
		sheet.NotesByTime = append(sheet.NotesByTime, nil)
	}
	if sheet.NotesByTime[timestamp] == nil {
		sheet.NotesByTime[timestamp] = make(map[string]int)
	}
}

// AI Generated. As the csv.Newreader returns a [][]string, this should work...
// I was a bit sloshed when reviewing the below. That I feel WIN! is natural.

// At this present time, I absolutly cannot summon any effort to re-write this.
// Wait untill I get more AI juice. Inform the AI that the SheetNote is no longer a Dictionary, but instead a matrix.
// Very simple. We have row and column index from the CSV! A CSV is a table, after all.
// Call the top column of the CSV for any differences between the CSV's Insturment-column key and the default.
// Throw an error if the IndexerInstruments isn't consistant (Or translate for the second and subsequant CSVs)
func LoadSheetNoteFromCSV(csvPath string) (SheetNote, IndexerInstruments, error) {
	file, err := os.Open(csvPath)
	if err != nil {
		return loadSheetNoteFromCSVError(fmt.Errorf("open csv %q: %w", csvPath, err))
	}
	defer file.Close()

	rows, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return loadSheetNoteFromCSVError( fmt.Errorf("read csv %q: %w", csvPath, err))
	}
	if len(rows) == 0 {
		return loadSheetNoteFromCSVError(fmt.Errorf("csv %q is empty", csvPath) )
	}
	if len(rows[0]) < 2 {
		return loadSheetNoteFromCSVError( fmt.Errorf("csv %q missing instrument header columns", csvPath) )
	}

	sheetNote := emptySheetNote()
	instruments := rows[0]
	for instIndex := 0; instIndex < len(instruments); instIndex++{
		instruments[instIndex] = InstrumentNameFor(instruments[instIndex])
	}
	maxTime := 0

	for rowIndex := 1; rowIndex < len(rows); rowIndex++ {
		row := rows[rowIndex]
		if len(row) == 0 {
			continue
		}

		timeText := strings.TrimSpace(row[0])
		if timeText == "" {
			continue
		}
		timestamp, err := strconv.Atoi(timeText)
		if err != nil {
			return SheetNote{}, fmt.Errorf("invalid time at row %d: %w", rowIndex+1, err)
		}

		for colIndex := 1; colIndex < len(instruments); colIndex++ {
			if colIndex >= len(row) {
				continue
			}

			instrument := strings.TrimSpace(instruments[colIndex])
			if instrument == "" {
				continue
			}

			pitchText := strings.TrimSpace(row[colIndex])
			if pitchText == "" {
				continue
			}

			pitch, err := strconv.Atoi(pitchText)
			if err != nil {
				return SheetNote{}, fmt.Errorf(
					"invalid pitch for instrument %q at row %d col %d: %w",
					instrument,
					rowIndex+1,
					colIndex+1,
					err,
				)
			}

			ensureTimeSlot(&sheetNote, timestamp)
			sheetNote.NotesByTime[timestamp][instrument] = pitch
		}
		maxTime = max(maxTime, timestamp)
	}
	sheetNote.CoveredLength = maxTime

	return sheetNote, nil
}

// Helper function for automatically creating empty SheetNote and Empty IndexerInstruments
func loadSheetNoteFromCSVError(errorMessage error) (SheetNote, IndexerInstruments, error){
	return mat.Dense{}, make(IndexerInstruments),	errorMessage
}