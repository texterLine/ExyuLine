// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

// BambooLeafSize represents the size of bamboo leaves.
type BambooLeafSize struct {
	bamboo
}

type bamboo uint8

// BambooSizeNoLeaves ...
func BambooSizeNoLeaves() BambooLeafSize {
	return BambooLeafSize{0}
}

// BambooSizeSmallLeaves ...
func BambooSizeSmallLeaves() BambooLeafSize {
	return BambooLeafSize{1}
}

// BambooSizeLargeLeaves ...
func BambooSizeLargeLeaves() BambooLeafSize {
	return BambooLeafSize{2}
}

// Uint8 ...
func (b bamboo) Uint8() uint8 {
	return uint8(b)
}

// String ...
func (b bamboo) String() string {
	switch b {
	case 0:
		return "no_leaves"
	case 1:
		return "small_leaves"
	case 2:
		return "large_leaves"
	}
	panic("unknown bamboo leaf size")
}

// Name ...
func (b bamboo) Name() string {
	switch b {
	case 0:
		return "No Leaves"
	case 1:
		return "Small Leaves"
	case 2:
		return "Large Leaves"
	}
	panic("unknown bamboo leaf size")
}

// BambooLeafSizes returns all possible bamboo leaf sizes.
func BambooLeafSizes() []BambooLeafSize {
	return []BambooLeafSize{BambooSizeNoLeaves(), BambooSizeSmallLeaves(), BambooSizeLargeLeaves()}
}
