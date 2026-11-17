// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

// DoubleFlowerType represents a type of double flower.
type DoubleFlowerType struct {
	doubleFlower
}

type doubleFlower uint8

// Sunflower is a sunflower plant.
func Sunflower() DoubleFlowerType {
	return DoubleFlowerType{0}
}

// Lilac is a lilac plant.
func Lilac() DoubleFlowerType {
	return DoubleFlowerType{1}
}

// RoseBush is a rose bush plant.
func RoseBush() DoubleFlowerType {
	return DoubleFlowerType{4}
}

// Peony is a peony plant.
func Peony() DoubleFlowerType {
	return DoubleFlowerType{5}
}

// Uint8 returns the double plant as a uint8.
func (d doubleFlower) Uint8() uint8 {
	return uint8(d)
}

// Name ...
func (d doubleFlower) Name() string {
	switch d {
	case 0:
		return "Sunflower"
	case 1:
		return "Lilac"
	case 4:
		return "Rose Bush"
	case 5:
		return "Peony"
	}
	panic("unknown double plant type")
}

// String ...
func (d doubleFlower) String() string {
	switch d {
	case 0:
		return "sunflower"
	case 1:
		return "lilac"
	case 4:
		return "rose_bush"
	case 5:
		return "peony"
	}
	panic("unknown double plant type")
}

// DoubleFlowerTypes ...
func DoubleFlowerTypes() []DoubleFlowerType {
	return []DoubleFlowerType{Sunflower(), Lilac(), RoseBush(), Peony()}
}
