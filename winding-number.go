package adventofcode2023

// This is basically a 1:1 copy of Dan Sunday's original from 2012.

import "image"

// isLeft tests if a point P2 is Left|On|Right of an infinite line through P0 and P1.
//
//	Input:  P0, P1, P2 (three points as image.Point)
//	Return: >0 for P2 left of the line through P0 and P1
//	        =0 for P2 on the line
//	        <0 for P2 right of the line
func isLeft(P0, P1, P2 image.Point) int {
	return (P1.X-P0.X)*(P2.Y-P0.Y) - (P2.X-P0.X)*(P1.Y-P0.Y)
}

// wnPnPoly returns the winding number test for a point in a polygon.
//
//	Input:   P = a point,
//	         V[] = vertex points of a polygon V[n+1] with V[n]=V[0]
//	Return:  wn = the winding number (=0 only when P is outside)
func wnPnPoly(P image.Point, V []image.Point) int {
	wn := 0 // the  winding number counter

	// loop through all edges of the polygon
	for i := 0; i < len(V)-1; i++ { // edge from V[i] to V[i+1]
		if V[i].Y <= P.Y { // start y <= P.Y
			if V[i+1].Y > P.Y { // an upward crossing
				if isLeft(V[i], V[i+1], P) > 0 { // P left of  edge
					wn++ // have a valid up intersect
				}
			}
		} else { // start y > P.Y (no test needed)
			if V[i+1].Y <= P.Y { // a downward crossing
				if isLeft(V[i], V[i+1], P) < 0 { // P right of  edge
					wn-- // have a valid down intersect
				}
			}
		}
	}
	return wn
}
