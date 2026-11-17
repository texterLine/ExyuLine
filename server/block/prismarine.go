// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

import (
	"github.com/texterLine/exyuline/server/world"
)

// Prismarine is a type of stone that appears underwater in ruins and ocean monuments.
type Prismarine struct {
	solid
	bassDrum

	// Type is the type of prismarine of the block.
	Type PrismarineType
}

// BreakInfo ...
func (p Prismarine) BreakInfo() BreakInfo {
	return newBreakInfo(1.5, pickaxeHarvestable, pickaxeEffective, oneOf(p)).withBlastResistance(6)
}

// EncodeItem ...
func (p Prismarine) EncodeItem() (id string, meta int16) {
	return "minecraft:" + p.Type.String(), 0
}

// EncodeBlock ...
func (p Prismarine) EncodeBlock() (name string, properties map[string]any) {
	return "minecraft:" + p.Type.String(), nil
}

// allPrismarine returns a list of all prismarine block variants.
func allPrismarine() (c []world.Block) {
	for _, t := range PrismarineTypes() {
		c = append(c, Prismarine{Type: t})
	}
	return
}
