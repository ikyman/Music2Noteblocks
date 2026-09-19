package utilitiesBeep

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func InstrumentNameFor(unaliasedName string) string{
	lowercaseUnaliased := strings.ToLower(unaliasedName)

	aliasedName := instrumentAliases[lowercaseUnaliased];
	
	if (aliasedName != ""){
		return aliasedName
	}
	return lowercaseUnaliased
}


type SheetNote struct {
	// index = time step; each entry is instrument -> pitch at that step
	NotesByTime []map[string]int

	CoveredLength int
}

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
