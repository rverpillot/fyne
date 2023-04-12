// SPDX-License-Identifier: Unlicense OR BSD-3-Clause

package tables

import (
	"encoding/binary"
	"fmt"
)

// Code generated by binarygen from head_src.go. DO NOT EDIT

func (item *Head) mustParse(src []byte) {
	_ = src[53] // early bound checking
	item.majorVersion = binary.BigEndian.Uint16(src[0:])
	item.minorVersion = binary.BigEndian.Uint16(src[2:])
	item.fontRevision = binary.BigEndian.Uint32(src[4:])
	item.checksumAdjustment = binary.BigEndian.Uint32(src[8:])
	item.magicNumber = binary.BigEndian.Uint32(src[12:])
	item.flags = binary.BigEndian.Uint16(src[16:])
	item.UnitsPerEm = binary.BigEndian.Uint16(src[18:])
	item.created = binary.BigEndian.Uint64(src[20:])
	item.modified = binary.BigEndian.Uint64(src[28:])
	item.XMin = int16(binary.BigEndian.Uint16(src[36:]))
	item.YMin = int16(binary.BigEndian.Uint16(src[38:]))
	item.XMax = int16(binary.BigEndian.Uint16(src[40:]))
	item.YMax = int16(binary.BigEndian.Uint16(src[42:]))
	item.macStyle = binary.BigEndian.Uint16(src[44:])
	item.lowestRecPPEM = binary.BigEndian.Uint16(src[46:])
	item.fontDirectionHint = int16(binary.BigEndian.Uint16(src[48:]))
	item.IndexToLocFormat = int16(binary.BigEndian.Uint16(src[50:]))
	item.glyphDataFormat = int16(binary.BigEndian.Uint16(src[52:]))
}

func ParseHead(src []byte) (Head, int, error) {
	var item Head
	n := 0
	if L := len(src); L < 54 {
		return item, 0, fmt.Errorf("reading Head: "+"EOF: expected length: 54, got %d", L)
	}
	item.mustParse(src)
	n += 54
	return item, n, nil
}