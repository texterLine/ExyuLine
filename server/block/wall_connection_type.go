// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

// WallConnectionType represents the connection type of a wall.
type WallConnectionType struct {
	wallConnectionType
}

// NoWallConnection returns the no connection type of a wall.
func NoWallConnection() WallConnectionType {
	return WallConnectionType{0}
}

// ShortWallConnection returns the short connection type of a wall.
func ShortWallConnection() WallConnectionType {
	return WallConnectionType{1}
}

// TallWallConnection returns the tall connection type of a wall.
func TallWallConnection() WallConnectionType {
	return WallConnectionType{2}
}

// WallConnectionTypes returns a list of all wall connection types.
func WallConnectionTypes() []WallConnectionType {
	return []WallConnectionType{NoWallConnection(), ShortWallConnection(), TallWallConnection()}
}

type wallConnectionType uint8

// Uint8 returns the wall connection as a uint8.
func (w wallConnectionType) Uint8() uint8 {
	return uint8(w)
}

// String ...
func (w wallConnectionType) String() string {
	switch w {
	case 0:
		return "none"
	case 1:
		return "short"
	case 2:
		return "tall"
	}
	panic("unknown wall connection type")
}

// Height returns the height of the connection for the block model.
func (w wallConnectionType) Height() float64 {
	switch w {
	case 0:
		return 0
	case 1:
		return 0.75
	case 2:
		return 1
	}
	panic("unknown wall connection type")
}
