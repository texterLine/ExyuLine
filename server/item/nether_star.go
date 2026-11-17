// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package item

// NetherStar is a rare item dropped by the wither that is used solely to craft beacons.
type NetherStar struct{}

// EncodeItem ...
func (NetherStar) EncodeItem() (name string, meta int16) {
	return "minecraft:nether_star", 0
}

// BlastProof indicates that the item will withstand an explosion as an item entity.
func (NetherStar) BlastProof() bool {
	return true
}
