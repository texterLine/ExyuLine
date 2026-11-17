// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package item

// WarpedFungusOnAStick is an item that can be used to control saddled striders.
type WarpedFungusOnAStick struct{}

// MaxCount ...
func (WarpedFungusOnAStick) MaxCount() int {
	return 1
}

// EncodeItem ...
func (WarpedFungusOnAStick) EncodeItem() (name string, meta int16) {
	return "minecraft:warped_fungus_on_a_stick", 0
}
