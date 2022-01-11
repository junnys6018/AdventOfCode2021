package day20

import (
	"advent/util"
	"fmt"
	"strings"
)

const CHUNK_SIZE = 16

type Pixel int

const (
	PIXEL_DARK Pixel = 0
	PIXEL_LIT  Pixel = 1
)

type Chunk [CHUNK_SIZE]uint16

type Point struct {
	x, y int
}

type ImageData map[Point]Chunk

func split(x, y int) (cx, fx, cy, fy int) {
	cx = x / CHUNK_SIZE
	fx = x % CHUNK_SIZE
	cy = y / CHUNK_SIZE
	fy = y % CHUNK_SIZE

	if fx < 0 {
		fx += CHUNK_SIZE
	}

	if x < 0 {
		cx = (x - 15) / CHUNK_SIZE
	}

	if fy < 0 {
		fy += CHUNK_SIZE
	}

	if y < 0 {
		cy = (y - 15) / CHUNK_SIZE
	}
	return
}

func (id ImageData) at(x, y int) Pixel {
	cx, fx, cy, fy := split(x, y)

	chunk, ok := id[Point{cx, cy}]

	if !ok {
		panic("bug!")
	}

	if chunk[fy]&(1<<fx) != 0 {
		return 1
	}
	return 0
}

func (id ImageData) set(x, y int, value Pixel) {
	cx, fx, cy, fy := split(x, y)

	chunkCoord := Point{cx, cy}
	chunk := id[chunkCoord]
	if value == PIXEL_LIT {
		chunk[fy] = chunk[fy] | (1 << fx)
	} else {
		chunk[fy] = chunk[fy] & ^(1 << fx)
	}
	id[chunkCoord] = chunk
}

type Image struct {
	data      ImageData
	expanse   Pixel // Whether the infinite expanse of pixels are lit or not
	algorithm [512]bool

	// Bounding box
	x1, x2, y1, y2 int
}

func (image *Image) count() (count int) {
	for y := image.y1; y <= image.y2; y++ {
		for x := image.x1; x <= image.x2; x++ {
			if image.at(x, y) == PIXEL_LIT {
				count++
			}
		}
	}
	return
}

func (image *Image) at(x, y int) Pixel {
	if x >= image.x1 && x <= image.x2 && y >= image.y1 && y <= image.y2 {
		return image.data.at(x, y)
	}
	return image.expanse
}

func (image *Image) getNumber(x, y int) uint16 {
	return uint16(image.at(x-1, y-1))<<8 |
		uint16(image.at(x, y-1))<<7 |
		uint16(image.at(x+1, y-1))<<6 |
		uint16(image.at(x-1, y))<<5 |
		uint16(image.at(x, y))<<4 |
		uint16(image.at(x+1, y))<<3 |
		uint16(image.at(x-1, y+1))<<2 |
		uint16(image.at(x, y+1))<<1 |
		uint16(image.at(x+1, y+1))

}

func (image *Image) enhance() {
	newData := make(ImageData)
	for y := image.y1 - 1; y <= image.y2+1; y++ {
		for x := image.x1 - 1; x <= image.x2+1; x++ {
			index := image.getNumber(x, y)
			if image.algorithm[index] {
				newData.set(x, y, PIXEL_LIT)
			} else {
				newData.set(x, y, PIXEL_DARK)
			}
		}
	}

	image.x1--
	image.x2++
	image.y1--
	image.y2++

	if image.expanse == PIXEL_DARK && image.algorithm[0] {
		image.expanse = PIXEL_LIT
	} else if image.expanse == PIXEL_LIT && !image.algorithm[511] {
		image.expanse = PIXEL_DARK
	}

	image.data = newData
}

func (image *Image) String() string {
	builder := strings.Builder{}

	for y := image.y1; y <= image.y2; y++ {
		for x := image.x1; x <= image.x2; x++ {
			if image.at(x, y) == PIXEL_LIT {
				builder.WriteRune('#')
			} else {
				builder.WriteRune('.')

			}
		}
		if y != image.y2 {
			builder.WriteRune('\n')
		}
	}

	return builder.String()
}

func parse(path string) *Image {
	image := Image{}
	image.data = make(ImageData)

	lines := util.ReadLines(path)

	algorithm := lines[0]

	for i, ch := range algorithm {
		if ch == '#' {
			image.algorithm[i] = true
		} else {
			image.algorithm[i] = false
		}
	}

	data := lines[2:]

	for y, row := range data {
		for x, ch := range row {
			if ch == '#' {
				image.data.set(x, y, PIXEL_LIT)
			}
		}
	}

	image.x2 = len(data[0]) - 1
	image.y2 = len(data) - 1

	return &image
}

func Answer() {
	input := "day20/input"

	image := parse(input)

	image.enhance()
	image.enhance()

	fmt.Printf("Answer (Part 1): %v\n", image.count())

	image = parse(input)

	for i := 0; i < 50; i++ {
		image.enhance()
	}

	fmt.Printf("Answer (Part 2): %v\n", image.count())
}
