// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

import "github.com/texterLine/exyuline/server/world"

// Terracotta is a block formed from clay, with a hardness and blast resistance comparable to stone. For colouring it,
// take a look at StainedTerracotta.
type Terracotta struct {
	solid
	bassDrum
}

// SoilFor ...
func (Terracotta) SoilFor(block world.Block) bool {
	_, ok := block.(DeadBush)
	return ok
}

// BreakInfo ...
func (t Terracotta) BreakInfo() BreakInfo {
	return newBreakInfo(1.25, pickaxeHarvestable, pickaxeEffective, oneOf(t)).withBlastResistance(4.2)
}

// EncodeItem ...
func (Terracotta) EncodeItem() (name string, meta int16) {
	return "minecraft:hardened_clay", meta
}

// EncodeBlock ...
func (Terracotta) EncodeBlock() (string, map[string]any) {
	return "minecraft:hardened_clay", nil
}
