package adventofcode2023

import (
	"bufio"
	"os"
	"testing"
)

// ... (Rest der Funktion Day10Gemini bleibt unverändert)

func TestDay10Gemini(t *testing.T) {
	t.Skip("Gemini-generated code - intentionally failing example documented in README.adoc")
	grid, err := loadGridFromText("testdata/day10_example.txt")
	if err != nil {
		t.Fatalf("Fehler beim Laden der Datei: %v", err)
	}

	const expected = 4 // Erwartetes Ergebnis für den Beispieltest
	actual := Day10Gemini(grid)

	if actual != expected {
		t.Errorf("Falsches Ergebnis: erwartet %d, aber erhalten %d", expected, actual)
	}
}

func loadGridFromText(filename string) ([][]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() { _ = file.Close() }()

	scanner := bufio.NewScanner(file)
	var grid [][]byte

	for scanner.Scan() {
		line := scanner.Bytes()
		grid = append(grid, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return grid, nil
}
