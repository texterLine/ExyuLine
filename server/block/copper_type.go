// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

// CopperType represents a type of copper.
type CopperType struct {
	copper
}

type copper uint8

// NormalCopper is the normal variant of copper.
func NormalCopper() CopperType {
	return CopperType{0}
}

// CutCopper is the cut variant of copper.
func CutCopper() CopperType {
	return CopperType{1}
}

// ChiseledCopper is the chiseled variant of copper.
func ChiseledCopper() CopperType {
	return CopperType{2}
}

// Uint8 returns the copper as a uint8.
func (s copper) Uint8() uint8 {
	return uint8(s)
}

// Name ...
func (s copper) Name() string {
	switch s {
	case 0:
		return "Copper"
	case 1:
		return "Cut Copper"
	case 2:
		return "Chiseled Copper"
	}
	panic("unknown copper type")
}

// String ...
func (s copper) String() string {
	switch s {
	case 0:
		return "default"
	case 1:
		return "cut"
	case 2:
		return "chiseled"
	}
	panic("unknown copper type")
}

// CopperTypes ...
func CopperTypes() []CopperType {
	return []CopperType{NormalCopper(), CutCopper(), ChiseledCopper()}
}
