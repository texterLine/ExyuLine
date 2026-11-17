// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

import (
	"github.com/texterLine/exyuline/server/item"
	"github.com/texterLine/exyuline/server/world/sound"
)

// Clay is a block that can be found underwater.
type Clay struct {
	solid
}

// Instrument ...
func (c Clay) Instrument() sound.Instrument {
	return sound.Flute()
}

// BreakInfo ...
func (c Clay) BreakInfo() BreakInfo {
	return newBreakInfo(0.6, alwaysHarvestable, shovelEffective, silkTouchDrop(item.NewStack(item.ClayBall{}, 4), item.NewStack(c, 1)))
}

// SmeltInfo ...
func (Clay) SmeltInfo() item.SmeltInfo {
	return newSmeltInfo(item.NewStack(Terracotta{}, 1), 0.35)
}

// EncodeItem ...
func (c Clay) EncodeItem() (name string, meta int16) {
	return "minecraft:clay", 0
}

// EncodeBlock ...
func (c Clay) EncodeBlock() (name string, properties map[string]any) {
	return "minecraft:clay", nil
}
