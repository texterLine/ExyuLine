// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package item

// WrittenBookGeneration represents a WrittenBook generation.
type WrittenBookGeneration struct {
	generation
}

type generation uint8

// OriginalGeneration is the original WrittenBook.
func OriginalGeneration() WrittenBookGeneration {
	return WrittenBookGeneration{0}
}

// CopyGeneration is a copy of the original WrittenBook.
func CopyGeneration() WrittenBookGeneration {
	return WrittenBookGeneration{1}
}

// CopyOfCopyGeneration is a copy of a copy of the original WrittenBook.
func CopyOfCopyGeneration() WrittenBookGeneration {
	return WrittenBookGeneration{2}
}

// Uint8 returns the generation as a uint8.
func (g generation) Uint8() uint8 {
	return uint8(g)
}

// String ...
func (g generation) String() string {
	switch g {
	case 0:
		return "original"
	case 1:
		return "copy of original"
	case 2:
		return "copy of copy"
	}
	panic("unknown written book generation")
}
