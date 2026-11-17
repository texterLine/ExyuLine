// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package item

// EnchantedBook is an item that lets players add enchantments to certain items using an anvil.
type EnchantedBook struct{}

// MaxCount ...
func (b EnchantedBook) MaxCount() int {
	return 1
}

// EncodeItem ...
func (EnchantedBook) EncodeItem() (name string, meta int16) {
	return "minecraft:enchanted_book", 0
}
