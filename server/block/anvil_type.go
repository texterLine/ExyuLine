// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

// AnvilType represents a type of anvil, such as undamaged, slightly damaged, or very damaged.
type AnvilType struct {
	anvil
}

// UndamagedAnvil returns the undamaged anvil type.
func UndamagedAnvil() AnvilType {
	return AnvilType{0}
}

// SlightlyDamagedAnvil returns the slightly damaged anvil type.
func SlightlyDamagedAnvil() AnvilType {
	return AnvilType{1}
}

// VeryDamagedAnvil returns the very damaged anvil type.
func VeryDamagedAnvil() AnvilType {
	return AnvilType{2}
}

// AnvilTypes returns all anvil types.
func AnvilTypes() []AnvilType {
	return []AnvilType{UndamagedAnvil(), SlightlyDamagedAnvil(), VeryDamagedAnvil()}
}

type anvil uint8

// Uint8 returns the anvil type as a uint8.
func (a anvil) Uint8() uint8 {
	return uint8(a)
}

// String returns the anvil type as a string.
func (a anvil) String() string {
	switch a {
	case 0:
		return "anvil"
	case 1:
		return "chipped_anvil"
	case 2:
		return "damaged_anvil"
	}
	panic("should never happen")
}
