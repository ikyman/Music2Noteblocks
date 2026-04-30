package utilitiesBeep

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type SheetNote struct {
	// instrument -> duration -> pitch
	NotesByInstrument map[string]map[int]int

	CoveredLength int
}

func emptySheetNote() SheetNote {
	return SheetNote{
		NotesByInstrument: make(map[string]map[int]int),
	}
}

// AI Generated. As the csv.Newreader returns a [][]string, this should work...
// I was a bit sloshed when reviewing the below. That I feel WIN! is natural. 
func LoadSheetNoteFromCSV(csvPath string) (SheetNote, error) {
	file, err := os.Open(csvPath)
	if err != nil {
		return SheetNote{}, fmt.Errorf("open csv %q: %w", csvPath, err)
	}
	defer file.Close()

	rows, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return SheetNote{}, fmt.Errorf("read csv %q: %w", csvPath, err)
	}
	if len(rows) == 0 {
		return SheetNote{}, fmt.Errorf("csv %q is empty", csvPath)
	}
	if len(rows[0]) < 2 {
		return SheetNote{}, fmt.Errorf("csv %q missing instrument header columns", csvPath)
	}

	sheetNote := emptySheetNote()
	instruments := rows[0]
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

			if sheetNote.NotesByInstrument[instrument] == nil {
				sheetNote.NotesByInstrument[instrument] = make(map[int]int)
			}
			sheetNote.NotesByInstrument[instrument][timestamp] = pitch
		}
		maxTime = max(maxTime, timestamp);
	}
	sheetNote.CoveredLength = maxTime

	return sheetNote, nil
}
