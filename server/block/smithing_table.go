// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

import (
	"github.com/texterLine/exyuline/server/block/cube"
	"github.com/texterLine/exyuline/server/item"
	"github.com/texterLine/exyuline/server/world"
)

// SmithingTable is a toolsmith's job site block that generates in villages. It can be used to upgrade diamond gear into
// netherite gear.
type SmithingTable struct {
	bass
	solid
}

// EncodeItem ...
func (SmithingTable) EncodeItem() (name string, meta int16) {
	return "minecraft:smithing_table", 0
}

// EncodeBlock ...
func (SmithingTable) EncodeBlock() (name string, properties map[string]interface{}) {
	return "minecraft:smithing_table", nil
}

// BreakInfo ...
func (s SmithingTable) BreakInfo() BreakInfo {
	return newBreakInfo(2.5, alwaysHarvestable, axeEffective, oneOf(s))
}

// Activate ...
func (SmithingTable) Activate(pos cube.Pos, _ cube.Face, tx *world.Tx, u item.User, _ *item.UseContext) bool {
	if opener, ok := u.(ContainerOpener); ok {
		opener.OpenBlockContainer(pos, tx)
		return true
	}
	return false
}
