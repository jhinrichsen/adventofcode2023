package adventofcode2023

import (
	"testing"
)

func TestFindFarthestCell(t *testing.T) {
	// Lade das Grid aus der Datei im testdata-Verzeichnis
	grid, err := LoadGridFromFile("testdata/day10_example.txt")
	if err != nil {
		t.Fatalf("Error loading grid: %v", err)
	}

	// Erwartete Koordinaten und maximale Distanz
	expectedX, expectedY, expectedDistance := 4, 3, 4

	// Finde die am weitesten entfernte Zelle
	x, y, distance := grid.FindFarthestCell()

	// Überprüfen, ob die Ergebnisse stimmen
	if x != expectedX || y != expectedY || distance != expectedDistance {
		t.Errorf("Expected (%d, %d, %d), got (%d, %d, %d)", expectedX, expectedY, expectedDistance, x, y, distance)
	}
}
