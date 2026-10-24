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

		csvIndexerInstruments[colIndex-1] = InstrumentAliasFor(instrument)
	}

	sheetNoteNumberOfTicks := len(rows) -1

	returnSheetNote := mat.NewDense(sheetNoteNumberOfTicks, ASSUMED_MAX_INSTRUMENTS, nil)

	for tickRow := 1; tickRow < sheetNoteNumberOfTicks; tickRow++ {
		row := rows[tickRow]
		if len(row) == 0 {
			continue
		}

		timeText := strings.TrimSpace(row[0])
		if timeText == "" {
			continue
		}
		timestamp, err := strconv.Atoi(timeText)
		if err != nil {
			return loadSheetNoteFromCSVError(fmt.Errorf("invalid time at row %d: %w", tickRow, err))
		}
		if timestamp < 0 {
			return loadSheetNoteFromCSVError(fmt.Errorf("negative time at row %d", tickRow))
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
			// Obvious bug: I can use numbers outside of Minecraft's acceptable notblock range.
			// Leaving it in because it could be funny.
			if err != nil {
				return loadSheetNoteFromCSVError(fmt.Errorf(
					"invalid pitch at row %d col %d: %w",
					tickRow, colIndex, err,
				))
			}
			returnSheetNote.Set(timestamp, colIndex-1, float64(pitch))

		}
	}

	return returnSheetNote, csvIndexerInstruments, nil
}

// Helper function for automatically creating empty SheetNote and Empty IndexerInstruments
func loadSheetNoteFromCSVError(errorMessage error) (SheetNote, IndexerInstruments, error) {
	return mat.NewDense(0, 0, nil), make(IndexerInstruments), errorMessage
}