package utilitiesBeep

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/mat"
)

type SheetNote = mat.Matrix

// I Instructed AI to Re-write this from a Map-based SheetNote to a Matrix-based SheetNote.
// That's why this function is too long.
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

	csvIndexerInstruments := make(IndexerInstruments)
	for colIndex := 1; colIndex < len(rows[0]); colIndex++ {
		instrument := strings.TrimSpace(rows[0][colIndex])
		if instrument == "" {
			continue
		}

		csvIndexerInstruments[colIndex] = InstrumentAliasFor(instrument)
	}

	sheetNoteNumberOfTicks := len(rows) -1

	returnSheetNote := mat.NewDense(sheetNoteNumberOfTicks, ASSUMED_MAX_INSTRUMENTS, nil)

	for rowIndex := 1; rowIndex < sheetNoteNumberOfTicks; rowIndex++ {
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
			return loadSheetNoteFromCSVError(fmt.Errorf("invalid time at row %d: %w", rowIndex+1, err))
		}
		if timestamp < 0 {
			return loadSheetNoteFromCSVError(fmt.Errorf("negative time at row %d", rowIndex+1))
		}

		for colIndex := 1; colIndex < len(rows[0]); colIndex++ {
			if colIndex >= len(row) {
				continue
			}

			pitchText := strings.TrimSpace(row[colIndex])
			if pitchText == "" {
				continue
			}
			pitch, err := strconv.Atoi(pitchText)
			if err != nil {
				return loadSheetNoteFromCSVError(fmt.Errorf(
					"invalid pitch for instrument %q at row %d col %d: %w",
					instrument, rowIndex+1, colIndex+1, err,
				))
			}
			returnSheetNote.Set(rowIndex-1, colIndex-1, pitch)

		}
	}

	return returnSheetNote, indexer, nil
}

// Helper function for automatically creating empty SheetNote and Empty IndexerInstruments
func loadSheetNoteFromCSVError(errorMessage error) (SheetNote, IndexerInstruments, error) {
	return mat.Dense{}, make(IndexerInstruments), errorMessage
}