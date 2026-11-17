// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package item

// DiscFragment is a music disc fragment obtained from ancient city loot chests. They are extremely rare to find and
// nine of them in a crafting table makes a music disc named, "5".
type DiscFragment struct{}

// EncodeItem ...
func (DiscFragment) EncodeItem() (name string, meta int16) {
	return "minecraft:disc_fragment_5", 0
}
