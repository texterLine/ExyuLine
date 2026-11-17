// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package item

// CarrotOnAStick is an item that can be used to control saddled pigs.
type CarrotOnAStick struct{}

// MaxCount ...
func (CarrotOnAStick) MaxCount() int {
	return 1
}

// EncodeItem ...
func (CarrotOnAStick) EncodeItem() (name string, meta int16) {
	return "minecraft:carrot_on_a_stick", 0
}
