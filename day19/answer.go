package day19

import (
	"advent/util"
	"fmt"
	"strings"
)

type Point struct {
	x, y, z int
}

func (p Point) String() string {
	return fmt.Sprintf("%d,%d,%d", p.x, p.y, p.z)
}

func (p Point) isZero() bool {
	return p.x == 0 && p.y == 0 && p.z == 0
}

func addPoint(p1, p2 Point) Point {
	return Point{p1.x + p2.x, p1.y + p2.y, p1.z + p2.z}
}

type Scanner []Point

func (sc Scanner) String() string {
	builder := strings.Builder{}

	for _, point := range sc {
		builder.WriteString(fmt.Sprintf("%v\n", point))
	}

	return builder.String()
}

type Transform [3][3]int

func (t *Transform) apply(p *Point) Point {
	return Point{
		p.x*t[0][0] + p.y*t[0][1] + p.z*t[0][2],
		p.x*t[1][0] + p.y*t[1][1] + p.z*t[1][2],
		p.x*t[2][0] + p.y*t[2][1] + p.z*t[2][2],
	}
}

func mulMat3(m1, m2 *Transform) Transform {
	return Transform{
		{m1[0][0]*m2[0][0] + m1[0][1]*m2[1][0] + m1[0][2]*m2[2][0], m1[0][0]*m2[0][1] + m1[0][1]*m2[1][1] + m1[0][2]*m2[2][1], m1[0][0]*m2[0][2] + m1[0][1]*m2[1][2] + m1[0][2]*m2[2][2]},
		{m1[1][0]*m2[0][0] + m1[1][1]*m2[1][0] + m1[1][2]*m2[2][0], m1[1][0]*m2[0][1] + m1[1][1]*m2[1][1] + m1[1][2]*m2[2][1], m1[1][0]*m2[0][2] + m1[1][1]*m2[1][2] + m1[1][2]*m2[2][2]},
		{m1[2][0]*m2[0][0] + m1[2][1]*m2[1][0] + m1[2][2]*m2[2][0], m1[2][0]*m2[0][1] + m1[2][1]*m2[1][1] + m1[2][2]*m2[2][1], m1[2][0]*m2[0][2] + m1[2][1]*m2[1][2] + m1[2][2]*m2[2][2]},
	}
}

func genRotation(formula string) Transform {
	result := I
	for _, ch := range formula {
		switch ch {
		case 'X':
			result = mulMat3(&result, &X)
		case 'Y':
			result = mulMat3(&result, &Y)
		default:
			panic("bad formula")
		}
	}
	return result
}

var I Transform = Transform{
	{1, 0, 0},
	{0, 1, 0},
	{0, 0, 1},
}

var X Transform = Transform{
	{1, 0, 0},
	{0, 0, -1},
	{0, 1, 0},
}

var Y Transform = Transform{
	{0, 0, 1},
	{0, 1, 0},
	{-1, 0, 0},
}

// https://stackoverflow.com/questions/16452383/how-to-get-all-24-rotations-of-a-3-dimensional-array
var transforms [24]Transform = [24]Transform{
	I,
	X,
	Y,
	genRotation("XX"),
	genRotation("XY"),
	genRotation("YX"),
	genRotation("YY"),
	genRotation("XXX"),
	genRotation("XXY"),
	genRotation("XYX"),
	genRotation("XYY"),
	genRotation("YXX"),
	genRotation("YYX"),
	genRotation("YYY"),
	genRotation("XXXY"),
	genRotation("XXYX"),
	genRotation("XXYY"),
	genRotation("XYXX"),
	genRotation("XYYY"),
	genRotation("YXXX"),
	genRotation("YYYX"),
	genRotation("XXXYX"),
	genRotation("XYXXX"),
	genRotation("XYYYX"),
}

func parse(path string) []Scanner {
	lines := util.ReadLines(path)
	currentScanner := Scanner{}
	scanners := []Scanner{}

	for _, line := range lines[1:] {
		if strings.HasPrefix(line, "--- scanner") {
			scanners = append(scanners, currentScanner)
			currentScanner = Scanner{}
		} else if line != "" {
			var x, y, z int
			fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
			currentScanner = append(currentScanner, Point{x, y, z})
		}
	}
	scanners = append(scanners, currentScanner)

	return scanners
}

type Edge struct {
	delta Point
	t     Transform
}

// Adjaceny matrix, delta = (0,0,0) indicates no edge
type RelationGraph [][]Edge

func (g RelationGraph) isAllConnectedToScanner0() bool {
	numScanners := len(g)
	for i := 1; i < numScanners; i++ {
		p := g[0][i].delta
		if p.isZero() {
			return false
		}
	}
	return true
}

func absInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Answer() {
	scanners := parse("day19/input")
	numScanners := len(scanners)

	relationGraph := make(RelationGraph, numScanners)
	for i := 0; i < numScanners; i++ {
		relationGraph[i] = make([]Edge, numScanners)
	}

	for i := 0; i < numScanners; i++ {
		for j := 0; j < numScanners; j++ {
			if i == j {
				continue
			}

			for tidx, transform := range transforms {
				deltas := make(map[Point]int)

				for _, pti := range scanners[i] {
					for _, ptj := range scanners[j] {

						Tptj := transform.apply(&ptj)
						delta := Point{pti.x - Tptj.x, pti.y - Tptj.y, pti.z - Tptj.z}
						deltas[delta]++
					}
				}

				for k, v := range deltas {
					if v >= 12 {
						relationGraph[i][j] = Edge{delta: k, t: transforms[tidx]}
						// fmt.Printf("relative to scanner %d, scanner %d is at (%v)\n", i, j, k)
						break
					}
				}
			}
		}
	}

	for !relationGraph.isAllConnectedToScanner0() {
		for i := 1; i < numScanners; i++ {
			p := relationGraph[0][i].delta
			transform := relationGraph[0][i].t
			if !p.isZero() {
				for j := 1; j < numScanners; j++ {
					if i == j {
						continue
					}

					delta := relationGraph[i][j].delta
					if !delta.isZero() {
						relationGraph[0][j] = Edge{delta: addPoint(p, transform.apply(&delta)), t: mulMat3(&transform, &relationGraph[i][j].t)}
					}
				}
			}
		}
	}

	// Taking the coordinate system of scanner 0 as the absolute coordinate system
	absolutePositions := relationGraph[0]

	// for i := 1; i < numScanners; i++ {
	// 	delta := relationGraph[0][i].delta
	// 	fmt.Printf("relative to scanner 0, scanner %d is at (%v)\n", i, delta)
	// }

	// count beacons
	beacons := make(map[Point]bool)

	// add scanner 0's beacons
	for _, beacon := range scanners[0] {
		beacons[beacon] = true
	}

	// add other beacons

	for sc, rel := range absolutePositions {
		if sc == 0 {
			continue
		}

		for _, beacon := range scanners[sc] {
			transformedBeacon := addPoint(rel.delta, rel.t.apply(&beacon))
			beacons[transformedBeacon] = true
		}
	}

	fmt.Printf("Answer (Part 1): %v\n", len(beacons))

	maxDist := 0
	for _, rel1 := range absolutePositions {
		for _, rel2 := range absolutePositions {
			p1 := rel1.delta
			p2 := rel2.delta

			dist := absInt(p1.x-p2.x) + absInt(p1.y-p2.y) + absInt(p1.z-p2.z)
			if dist > maxDist {
				maxDist = dist
			}
		}
	}

	fmt.Printf("Answer (Part 2): %v\n", maxDist)
}
