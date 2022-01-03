package day16

import (
	"encoding/hex"
	"fmt"
	"math"
	"os"
)

type Packet struct {
	version  uint8
	typeID   uint8
	literal  uint64
	children []Packet
}

func (packet Packet) CountVersionNumbers() (sum int) {
	sum += int(packet.version)
	if packet.children != nil {
		for _, child := range packet.children {
			sum += child.CountVersionNumbers()
		}
	}
	return
}

func (packet Packet) Evaluate() uint64 {
	switch packet.typeID {
	case 0:
		var sum uint64
		for _, child := range packet.children {
			sum += child.Evaluate()
		}
		return sum
	case 1:
		var product uint64 = 1
		for _, child := range packet.children {
			product *= child.Evaluate()
		}
		return product
	case 2:
		var min uint64 = math.MaxUint64
		for _, child := range packet.children {
			v := child.Evaluate()
			if v < min {
				min = v
			}
		}
		return min
	case 3:
		var max uint64
		for _, child := range packet.children {
			v := child.Evaluate()
			if v > max {
				max = v
			}
		}
		return max
	case 4:
		return packet.literal
	case 5:
		if packet.children[0].Evaluate() > packet.children[1].Evaluate() {
			return 1
		}
		return 0
	case 6:
		if packet.children[0].Evaluate() < packet.children[1].Evaluate() {
			return 1
		}
		return 0
	case 7:
		if packet.children[0].Evaluate() == packet.children[1].Evaluate() {
			return 1
		}
		return 0
	default:
		panic("Invalid packet typeID")
	}
}

type BitStream []byte

func (bs BitStream) at(offset int) uint8 {
	index := offset / 8
	shift := offset % 8
	data := bs[index]

	return (data << byte(shift)) >> 7
}

func (bs BitStream) readBits(offset, length int) (value uint64) {
	for i := 0; i < length; i++ {
		value = value | (uint64(bs.at(offset+i)) << (length - i - 1))
	}
	return
}

func parseHexString(s string) BitStream {
	if len(s)%2 == 1 {
		s = s + "0"
	}

	bytes, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return bytes
}

func parsePacket(stream BitStream, offset int) (Packet, int) {
	var version, typeID uint8

	version = uint8(stream.readBits(offset, 3))
	offset += 3
	typeID = uint8(stream.readBits(offset, 3))
	offset += 3

	if typeID == 4 /* literal */ {
		var literal uint64

		for stream.at(offset) == 1 {
			offset += 1
			nibble := stream.readBits(offset, 4)
			offset += 4
			literal = literal<<4 | nibble
		}
		offset += 1
		nibble := stream.readBits(offset, 4)
		offset += 4
		literal = literal<<4 | nibble

		return Packet{version: version, typeID: typeID, literal: literal}, offset
	}

	/* operator */
	packet := Packet{version: version, typeID: typeID}

	lengthTypeID := stream.at(offset)
	offset++
	if lengthTypeID == 0 {
		totalLength := stream.readBits(offset, 15)
		offset += 15

		for oldOffset := offset; offset-oldOffset < int(totalLength); {
			var subPacket Packet

			subPacket, offset = parsePacket(stream, offset)
			packet.children = append(packet.children, subPacket)
		}

	} else {
		subPacketCount := stream.readBits(offset, 11)
		offset += 11

		for i := 0; i < int(subPacketCount); i++ {
			var subPacket Packet

			subPacket, offset = parsePacket(stream, offset)
			packet.children = append(packet.children, subPacket)
		}
	}

	return packet, offset
}

func Answer() {
	hexString, err := os.ReadFile("day16/input")
	if err != nil {
		panic(err)
	}

	stream := parseHexString(string(hexString))
	packet, _ := parsePacket(stream, 0)

	fmt.Printf("Answer (Part 1): %v\n", packet.CountVersionNumbers())
	fmt.Printf("Answer (Part 2): %v\n", packet.Evaluate())
}
