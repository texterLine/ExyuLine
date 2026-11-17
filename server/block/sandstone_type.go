// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

// SandstoneType represents a type of sandstone.
type SandstoneType struct {
	sandstone
}

type sandstone uint8

// NormalSandstone is the normal variant of sandstone.
func NormalSandstone() SandstoneType {
	return SandstoneType{0}
}

// CutSandstone is the cut variant of sandstone.
func CutSandstone() SandstoneType {
	return SandstoneType{1}
}

// ChiseledSandstone is the chiseled variant of sandstone.
func ChiseledSandstone() SandstoneType {
	return SandstoneType{2}
}

// SmoothSandstone is the smooth variant of sandstone.
func SmoothSandstone() SandstoneType {
	return SandstoneType{3}
}

// Uint8 returns the sandstone as a uint8.
func (s sandstone) Uint8() uint8 {
	return uint8(s)
}

// Name ...
func (s sandstone) Name() string {
	switch s {
	case 0:
		return "Sandstone"
	case 1:
		return "Cut Sandstone"
	case 2:
		return "Chiseled Sandstone"
	case 3:
		return "Smooth Sandstone"
	}
	panic("unknown sandstone type")
}

// String ...
func (s sandstone) String() string {
	switch s {
	case 0:
		return "default"
	case 1:
		return "cut"
	case 2:
		return "chiseled"
	case 3:
		return "smooth"
	}
	panic("unknown sandstone type")
}

// SandstoneTypes ...
func SandstoneTypes() []SandstoneType {
	return []SandstoneType{NormalSandstone(), CutSandstone(), ChiseledSandstone(), SmoothSandstone()}
}
