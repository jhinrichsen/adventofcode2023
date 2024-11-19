package adventofcode2023

import (
	"image"
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

// GenerateGrid generates a grid of points within a bounding box
func GenerateGrid(xMin, xMax, yMin, yMax int) []image.Point {
	var grid []image.Point
	for x := xMin; x <= xMax; x++ {
		for y := yMin; y <= yMax; y++ {
			grid = append(grid, image.Point{X: x, Y: y})
		}
	}
	return grid
}

func TestIsPointInPolygonStrict(t *testing.T) {
	// Define the polygon
	polygon := []image.Point{
		{12, 4}, {13, 4}, {13, 5}, {13, 6}, {13, 7}, {13, 8}, {13, 9},
		{14, 9}, {14, 8}, {14, 7}, {15, 7}, {15, 8}, {15, 9}, {16, 9},
		{16, 8}, {16, 7}, {16, 6}, {15, 6}, {15, 5}, {14, 5}, {14, 4},
		{15, 4}, {16, 4}, {16, 5}, {17, 5}, {17, 6}, {18, 6}, {18, 7},
		{19, 7}, {19, 6}, {19, 5}, {18, 5}, {18, 4}, {17, 4}, {17, 3},
		{16, 3}, {15, 3}, {15, 2}, {14, 2}, {14, 1}, {15, 1}, {15, 0},
		{14, 0}, {13, 0}, {13, 1}, {13, 2}, {13, 3}, {12, 3}, {12, 2},
		{12, 1}, {12, 0}, {11, 0}, {11, 1}, {11, 2}, {11, 3}, {11, 4},
		{10, 4}, {10, 3}, {10, 2}, {10, 1}, {10, 0}, {9, 0}, {9, 1},
		{9, 2}, {9, 3}, {8, 3}, {8, 2}, {8, 1}, {8, 0}, {7, 0}, {7, 1},
		{7, 2}, {7, 3}, {6, 3}, {6, 2}, {6, 1}, {6, 0}, {5, 0}, {4, 0},
		{3, 0}, {2, 0}, {1, 0}, {1, 1}, {1, 2}, {1, 3}, {0, 3}, {0, 4},
		{1, 4}, {2, 4}, {3, 4}, {3, 3}, {2, 3}, {2, 2}, {2, 1}, {3, 1},
		{4, 1}, {5, 1}, {5, 2}, {4, 2}, {4, 3}, {5, 3}, {5, 4}, {6, 4},
		{6, 5}, {5, 5}, {4, 5}, {4, 6}, {5, 6}, {5, 7}, {5, 8}, {4, 8},
		{4, 9}, {5, 9}, {6, 9}, {7, 9}, {8, 9}, {8, 8}, {7, 8}, {6, 8},
		{6, 7}, {7, 7}, {7, 6}, {8, 6}, {8, 7}, {9, 7}, {9, 6}, {9, 5},
		{10, 5}, {10, 6}, {10, 7}, {10, 8}, {10, 9}, {11, 9}, {11, 8},
		{11, 7}, {12, 7}, {12, 6}, {11, 6}, {11, 5}, {12, 5}, {12, 4},
	}

	// Generate grid points
	grid := GenerateGrid(0, 20, 0, 10)

	// Check points strictly inside
	var strictlyInside []image.Point
	for _, point := range grid {
		if IsPointInPolygonStrict(point, polygon) {
			strictlyInside = append(strictlyInside, point)
		}
	}

	// Validate results
	if len(strictlyInside) == 0 {
		t.Errorf("Expected strictly inside points, but found none")
	}
	t.Logf("Points strictly inside: %v", strictlyInside)
}
