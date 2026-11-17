// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

import "github.com/texterLine/exyuline/server/item"

// Snow is a full-sized block of snow.
type Snow struct {
	solid
}

// BreakInfo ...
func (s Snow) BreakInfo() BreakInfo {
	return newBreakInfo(0.2, alwaysHarvestable, shovelEffective, silkTouchDrop(item.NewStack(item.Snowball{}, 4), item.NewStack(s, 1)))
}

// EncodeItem ...
func (Snow) EncodeItem() (name string, meta int16) {
	return "minecraft:snow", 0
}

// EncodeBlock ...
func (Snow) EncodeBlock() (string, map[string]any) {
	return "minecraft:snow", nil
}
